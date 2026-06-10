package service

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/L1mus/Tickitz-backend/internal/cache"
	"github.com/L1mus/Tickitz-backend/internal/dto"
	"github.com/L1mus/Tickitz-backend/internal/repository"
	"github.com/L1mus/Tickitz-backend/pkg"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type UserService struct {
	userRepository *repository.UserRepository
	rdb            *redis.Client
	db             *pgxpool.Pool
}

func NewUserService(userRepository *repository.UserRepository, rdb *redis.Client, db *pgxpool.Pool) *UserService {
	return &UserService{
		userRepository: userRepository,
		rdb:            rdb,
		db:             db,
	}
}

const userProfileTTL = 10 * time.Minute

func (s *UserService) GetProfile(ctx context.Context, userID int) (dto.UserProfileResponse, error) {
	/*
		 Cek cache Redis dulu dengan key: "user:profile:{userID}"
		 Jika ada di cache → parse dan return langsung (tidak query DB)
		 Jika tidak ada di cache:
		     Ambil dari DB (first_name, last_name, email, phone)
		     Jika error → return error
		     Simpan ke Redis dengan TTL (misalnya 10 menit)
		     return data profile, nil
		strategi cache-aside:
		1. Cek Redis dulu (key: "user:profile:{userID}")
		2. Jika HIT → parse dan return langsung (tanpa query DB)
		3. Jika MISS → query DB → simpan ke Redis (TTL 10 menit) → return
	*/
	cacheKey := fmt.Sprintf("user:profile:%d", userID)
	var resp dto.UserProfileResponse
	hit, err := cache.GetFromCache(ctx, s.rdb, cacheKey, &resp)
	if err != nil {
		_ = err
	}
	if hit {
		return resp, nil
	}

	profile, err := s.userRepository.GetUserProfile(ctx, userID)
	if err != nil {
		return dto.UserProfileResponse{}, fmt.Errorf("get user profile from db: %w", err)
	}

	fullName := strings.TrimSpace(profile.FirstName + " " + profile.LastName)

	resp = dto.UserProfileResponse{
		FirstName: &profile.FirstName,
		LastName:  &profile.LastName,
		FullName:  &fullName,
		Email:     profile.Email,
		Phone:     &profile.Phone,
		Photo:     &profile.Photo,
	}
	_ = cache.SaveToCache(ctx, s.rdb, cacheKey, resp, userProfileTTL)

	return resp, nil
}

func (s *UserService) InvalidateProfileCache(ctx context.Context, userID int) error {
	cacheKey := fmt.Sprintf("user:profile:%d", userID)
	return cache.DelFromCache(ctx, s.rdb, cacheKey)
}

func (s *UserService) UpdateProfile(ctx context.Context, userID int, req dto.UserUpdateProfileReq, photoURL *string) (dto.UserUpdateProfileRes, error) {
	var hashPassword *string

	if req.NewPassword != nil && *req.NewPassword != "" {
		if req.ConfirmPassword == nil || *req.NewPassword != *req.ConfirmPassword {
			return dto.UserUpdateProfileRes{}, fmt.Errorf("Confirm Password Does Not Match New Password")
		}

		var hash pkg.HashConfig
		hash.UseRecommended()
		hashed := hash.GenHash(*req.NewPassword)
		hashPassword = &hashed
	}

	user, err := s.userRepository.UpdateProfileById(ctx, userID, req.FirstName, req.LastName, req.Phone, photoURL, hashPassword)
	if err != nil {
		return dto.UserUpdateProfileRes{}, err
	}

	cacheKey := fmt.Sprintf("user:profile:%d", userID)
	errCache := cache.DelFromCache(ctx, s.rdb, cacheKey)
	if errCache != nil {
		log.Printf("failed to delete cache for key %s: %v", cacheKey, errCache)
	}

	return dto.UserUpdateProfileRes{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Phone:     user.Phone,
		Photo:     user.Photo,
	}, nil
}

func (s *UserService) GetOrderHistory(ctx context.Context, userID int) ([]dto.OrderHistoryRes, error) {
	dbHistory, err := s.userRepository.GetOrderHistoryById(ctx, userID)
	if err != nil {
		return nil, err
	}

	var responseList []dto.OrderHistoryRes

	for _, item := range dbHistory {
		formatDate := item.Showtime.Date.Format("Tuesday, 07 July 2020")

		formatTime := item.Showtime.Time
		parsedTime, err := time.Parse("16:30:00", item.Showtime.Time)
		if err == nil {
			formatTime = strings.ToLower(parsedTime.Format("04:30pm"))
		}

		fullShowtime := fmt.Sprintf("%s - %s", formatDate, formatTime)

		ticketStatus := "Ticket used"
		if item.Booking.StatusTicket == "active" {
			ticketStatus = "Ticket in active"
		}

		paymentStatus := "Not Paid"
		if item.Booking.StatusPaid == "paid" {
			paymentStatus = "Paid"
		}

		responseList = append(responseList, dto.OrderHistoryRes{
			BookingId:    item.Booking.Id,
			MovieTitle:   item.Movie.Title,
			CinemaName:   item.Cinema.Name,
			CinemaLogo:   item.Cinema.Logo,
			Showtime:     fullShowtime,
			StatusTicket: ticketStatus,
			StatusPaid:   paymentStatus,
		})

	}
	return responseList, nil
}

func (s *UserService) GetInformationDetail(ctx context.Context, bookingID, userID int) (*dto.DetailInformationRes, error) {
	raw, err := s.userRepository.GetDetailById(ctx, bookingID, userID)
	if err != nil {
		return nil, err
	}

	ticketStatus := "Ticket used"
	if raw.StatusTicket == "active" {
		ticketStatus = "Ticket in active"
	}

	paymentStatus := "Ticked used"
	if raw.StatusPaid == "paid" {
		paymentStatus = "Paid"
	}

	res := dto.DetailInformationRes{
		BookingId:    raw.BookingId,
		StatusTicket: ticketStatus,
		StatusPaid:   paymentStatus,
	}

	if raw.TotalPrice != nil {
		res.TotalPrice = *raw.TotalPrice
	}

	if raw.StatusPaid != "paid" {
		if raw.VirtualRek != nil {
			res.VirtualRek = *raw.VirtualRek
		}

		dueDate := raw.CreatedAt.Add(24 * time.Hour)
		res.DueDateMessage = fmt.Sprintf("Pay this payment bill before it is due, on %s. If the bill has not been paid by the specified time, it will be forfeited", dueDate.Format("January 02, 2006"))
	} else {
		if raw.QrCode != nil {
			res.QrCode = *raw.QrCode
		}
		res.Category = raw.Category
		res.MovieTitle = raw.MovieTitle
		res.Quantity = raw.Quantity
		res.ShowtimeDate = raw.ShowtimeDate.Format("02 Jan")

		t, err := time.Parse("15:04:00", raw.ShowtimeTime)
		if err == nil {
			res.ShowtimeTime = strings.ToLower(t.Format("03:04pm"))
		}

		if raw.SeatList != nil && *raw.SeatList != "" {
			res.Seats = strings.Split(*raw.SeatList, ",")
		} else {
			res.Seats = []string{}
		}
	}

	return &res, nil
}
