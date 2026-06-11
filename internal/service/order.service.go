package service

import (
	"context"
	"log"

	apperror "github.com/L1mus/Tickitz-backend/internal/appError"
	"github.com/L1mus/Tickitz-backend/internal/dto"
	"github.com/L1mus/Tickitz-backend/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderService struct {
	orderRepo *repository.OrderRepository
	db        *pgxpool.Pool
}

func NewOrderService(orderRepo *repository.OrderRepository, db *pgxpool.Pool) *OrderService {
	return &OrderService{
		orderRepo: orderRepo,
		db:        db,
	}
}

func (s *OrderService) GetSeats(ctx context.Context, showtimeID int) (dto.SeatPageResponse, error) {
	/*
	   Mulai DB transaction (BEGIN)
	   defer: rollback jika ada panic/error
	   Ambil summary showtime (movie, cinema, tanggal, harga)
	   Jika error → rollback, return error
	   Ambil cinema_id dari summary (dibutuhkan untuk query kursi)
	   Ambil semua kursi + status (Available/Sold) untuk showtime & cinema ini
	   Jika error → rollback, return error
	   Commit DB transaction
	   return response, nil
	*/
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return dto.SeatPageResponse{}, err
	}
	defer tx.Rollback(ctx)

	detailMovie, err := s.orderRepo.GetShowtimeSummary(ctx, tx, showtimeID)
	if err != nil {
		return dto.SeatPageResponse{}, err
	}

	data, err := s.orderRepo.GetSeatsByShowtime(ctx, tx, showtimeID, detailMovie.CinemaID)
	if err != nil {
		return dto.SeatPageResponse{}, err
	}

	var seats []dto.SeatDTO
	for _, seat := range data {
		seat := dto.SeatDTO{
			SeatID:     seat.SeatID,
			Row:        seat.Row,
			SeatNumber: seat.SeatNumber,
			SeatType:   seat.SeatType,
			SeatStatus: seat.SeatStatus,
		}
		seats = append(seats, seat)
	}

	if err := tx.Commit(ctx); err != nil {
		return dto.SeatPageResponse{}, err
	}

	return dto.SeatPageResponse{
		Summary: dto.SummaryMovieDTO{
			MovieTitle:  detailMovie.MovieTitle,
			MoviePoster: detailMovie.MoviePoster,
			Category:    detailMovie.Category,
			CinemaName:  detailMovie.CinemaName,
			CinemaLogo:  detailMovie.CinemaLogo,
			ShowDate:    detailMovie.ShowDate,
			ShowTime:    detailMovie.ShowTime,
			TicketPrice: detailMovie.TicketPrice,
		},
		Seats: seats,
	}, nil
}

func (s *OrderService) CreateBooking(ctx context.Context, req dto.CreateBookingRequest, userID int) (dto.CreateBookingResponse, error) {
	/*
	   Validasi: pastikan req.SeatIDs tidak kosong
	   Validasi: pastikan jumlah seat == req.Quantity
	   Mulai DB transaction (BEGIN)
	   defer: rollback jika ada panic/error
	   Re-cek kursi yang diminta apakah masih Available
	   (query ulang status kursi di dalam transaction)
	   Jika ada yang sudah Sold → rollback, return error "Seat already taken"
	   INSERT ke tabel bookings
	       → set status_ticket = 'active', status_paid = 'not_paid'
	       → return booking_id
	   Loop setiap seat_id di req.SeatIDs:
	       INSERT ke booking_seats (booking_id, seat_id, showtime_id)
	       Jika error (kemungkinan race condition / duplicate) → rollback, return error
	   Commit DB transaction
	   return booking_id, nil
	*/
	if len(req.SeatIDs) == 0 || len(req.SeatIDs) < 0 {
		return dto.CreateBookingResponse{}, apperror.InvalidSeatsInput
	}
	if len(req.SeatIDs) != req.Quantity {
		return dto.CreateBookingResponse{}, apperror.InvalidQuantity
	}
	tx, err := s.db.Begin(ctx)
	if err != nil {
		log.Println("failed to begin transaction", err)
	}
	defer tx.Rollback(ctx)

	takenCount, err := s.orderRepo.CheckSeatsAvailable(ctx, tx, req.SeatIDs, req.ShowtimeID)
	if err != nil {
		return dto.CreateBookingResponse{}, err
	}
	if takenCount > 0 {
		return dto.CreateBookingResponse{}, apperror.SeatsUnavailable
	}

	bookingId, err := s.orderRepo.CreateBooking(ctx, tx, req, userID)
	if err != nil {
		return dto.CreateBookingResponse{}, err
	}

	for _, v := range req.SeatIDs {
		if err := s.orderRepo.CreateBookingSeat(ctx, tx, bookingId, v, req.ShowtimeID); err != nil {
			return dto.CreateBookingResponse{}, err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return dto.CreateBookingResponse{}, err
	}

	return dto.CreateBookingResponse{
		BookingID: bookingId,
	}, nil
}
