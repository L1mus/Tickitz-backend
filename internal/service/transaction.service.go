package service

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	apperror "github.com/L1mus/Tickitz-backend/internal/appError"
	"github.com/L1mus/Tickitz-backend/internal/dto"
	"github.com/L1mus/Tickitz-backend/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionService struct {
	transactionRepo *repository.TransactionRepository
	orderRepo       *repository.OrderRepository
	db              *pgxpool.Pool
}

func NewTransactionService(transactionRepo *repository.TransactionRepository, OrderRepo *repository.OrderRepository, db *pgxpool.Pool) *TransactionService {
	return &TransactionService{
		transactionRepo: transactionRepo,
		orderRepo:       OrderRepo,
		db:              db,
	}
}

func (s *TransactionService) GetPaymentPage(ctx context.Context, bookingID int, userID int) (dto.PaymentPageResponse, error) {
	/*
	   Ambil summary booking (movie, cinema, qty, total) by bookingID
	   Jika error → return error
	   Validasi: pastikan booking.user_id == userID
	   (user tidak boleh akses booking milik orang lain)
	   Jika tidak match → return error "Forbidden"
	   Ambil label kursi yang dipilih ("C4, C5, C6") by bookingID
	   Ambil data user yang login (nama, email, phone) untuk pre-fill form
	   Ambil semua metode pembayaran untuk ditampilkan di grid
	   Rakit semua jadi satu response
	   return response, nil
	*/
	tx, err := s.db.Begin(ctx)
	if err != nil {
		log.Println(err)
		return dto.PaymentPageResponse{}, err
	}
	defer tx.Rollback(ctx)
	booking, err := s.transactionRepo.GetBookingSummary(ctx, tx, bookingID)
	if err != nil {
		return dto.PaymentPageResponse{}, apperror.BookingNotFound
	}

	if userID != booking.UserID {
		return dto.PaymentPageResponse{}, apperror.UserIstMatch
	}

	dataSeats, err := s.transactionRepo.GetBookedSeats(ctx, tx, booking.BookingID)
	if err != nil {
		return dto.PaymentPageResponse{}, err
	}

	seats := make([]string, len(dataSeats))
	for i, u := range dataSeats {
		seats[i] = u.Label
	}

	userinfo, err := s.transactionRepo.GetUser(ctx, tx, userID)
	if err != nil {
		return dto.PaymentPageResponse{}, apperror.ErrUserNotFound
	}

	dataPaymentMethods, err := s.transactionRepo.GetAllPaymentMethods(ctx, tx)
	if err != nil {
		return dto.PaymentPageResponse{}, err
	}

	var paymentMethods []dto.PaymentMethodDTO
	for _, payment := range dataPaymentMethods {
		m := dto.PaymentMethodDTO{
			ID:   payment.ID,
			Name: payment.Name,
			Logo: payment.Logo,
		}
		paymentMethods = append(paymentMethods, m)
	}
	if err := tx.Commit(ctx); err != nil {
		return dto.PaymentPageResponse{}, err
	}
	return dto.PaymentPageResponse{
		BookingID:      booking.BookingID,
		MovieTitle:     booking.MovieTitle,
		CinemaName:     booking.CinemaName,
		ShowDate:       booking.ShowDate,
		ShowTime:       booking.ShowTime,
		Quantity:       booking.Quantity,
		SeatLabels:     strings.Join(seats, ", "),
		TotalPrice:     booking.TotalPayment,
		FullName:       fmt.Sprintf("%s %s", userinfo.FirstName, userinfo.LastName),
		Email:          userinfo.Email,
		Phone:          userinfo.Phone,
		PaymentMethods: paymentMethods,
	}, nil
}

func (s *TransactionService) SubmitPayment(ctx context.Context, userID int, req dto.SubmitPaymentRequest) (dto.TransactionModalResponse, error) {
	/*
	   Hitung total harga:
	       total = booking.quantity × showtime.price
	   Generate nomor virtual rekening (angka acak unik, simpan sebagai BIGINT)
	   Mulai DB transaction (BEGIN)
	   defer: rollback jika ada panic/error
	   INSERT ke transactions:
	       → booking_id, payment_method_id, virtual_rek, total_price, status = 'pending'
	       → return transaction_id
	   Commit DB transaction
	   Ambil data modal (virtual_rek, total, deadline) by bookingID
	   return data modal, nil
	*/
	summary, err := s.transactionRepo.GetBookingSummary(ctx, s.db, req.BookingID)
	if err != nil {
		return dto.TransactionModalResponse{}, err
	}
	if summary.UserID != userID {
		return dto.TransactionModalResponse{}, apperror.ForbiddenBooking
	}

	totalPrice := summary.TotalPayment

	virtualRek := generateVirtualRek()

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return dto.TransactionModalResponse{}, err
	}
	defer tx.Rollback(ctx)

	_, err = s.transactionRepo.CreateTransaction(ctx, tx, req.BookingID, req.PaymentMethodID, virtualRek, totalPrice)
	if err != nil {
		return dto.TransactionModalResponse{}, err
	}

	if err := tx.Commit(ctx); err != nil {
		return dto.TransactionModalResponse{}, err
	}

	modal, err := s.transactionRepo.GetTransactionModal(ctx, s.db, req.BookingID)
	if err != nil {
		return dto.TransactionModalResponse{}, err
	}

	return dto.TransactionModalResponse{
		TransactionID:   modal.TransactionID,
		VirtualRek:      modal.VirtualRek,
		TotalPrice:      modal.TotalPrice,
		Status:          modal.Status,
		PaymentDeadline: modal.PaymentDeadline,
	}, nil
}

func (s *TransactionService) CheckPayment(ctx context.Context, transactionID, bookingID int) (dto.TicketResultResponse, error) {
	/*
	      [Di production: cek ke payment gateway apakah sudah dibayar]
	      [Di sini kita simulasikan selalu berhasil]
	      Generate URL QR code untuk tiket
	      Mulai DB transaction (BEGIN)
	      defer: rollback jika ada panic/error
	      UPDATE transactions:
	   	   → set status = 'completed', qr_code = url
	   	   by transaction_id
	      UPDATE bookings:
	   	   → set status_paid = 'paid', status_ticket = 'active', updated_at = NOW()
	   	   by booking_id
	      Commit DB transaction
	      Ambil data tiket lengkap (qr_code, movie, kursi, total) by transactionID
	      return data tiket, nil
	*/
	isAlreadyPay, err := s.transactionRepo.CheckAlreadyTransaction(ctx, s.db, transactionID)
	if err != nil {
		return dto.TicketResultResponse{}, err
	}
	if isAlreadyPay > 0 {
		return dto.TicketResultResponse{}, apperror.TicketAlreadyPaid
	}
	qrCode := fmt.Sprintf("%d-%d-%d", transactionID, bookingID, time.Now().Unix())

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return dto.TicketResultResponse{}, err
	}
	defer tx.Rollback(ctx)

	if err := s.transactionRepo.UpdateTransactionStatus(ctx, tx, transactionID, qrCode); err != nil {
		return dto.TicketResultResponse{}, err
	}

	if err := s.transactionRepo.UpdateBookingStatus(ctx, tx, bookingID); err != nil {
		return dto.TicketResultResponse{}, err
	}

	if err := tx.Commit(ctx); err != nil {
		return dto.TicketResultResponse{}, err
	}

	ticket, err := s.transactionRepo.GetTicketResult(ctx, s.db, transactionID)
	if err != nil {
		return dto.TicketResultResponse{}, err
	}

	return dto.TicketResultResponse{
		QRCode:        ticket.QRCode,
		TotalPrice:    ticket.TotalPrice,
		PaymentStatus: ticket.PaymentStatus,
		MovieTitle:    ticket.MovieTitle,
		Category:      ticket.Category,
		ShowDate:      ticket.ShowDate,
		ShowTime:      ticket.ShowTime,
		TicketCount:   ticket.TicketCount,
		SeatLabels:    ticket.SeatLabels,
	}, nil
}

func (s *TransactionService) GetTicketResult(ctx context.Context, transactionID int) (dto.TicketResultResponse, error) {
	ticket, err := s.transactionRepo.GetTicketResult(ctx, s.db, transactionID)
	if err != nil {
		return dto.TicketResultResponse{}, apperror.TicketNotFound
	}

	return dto.TicketResultResponse{
		QRCode:        ticket.QRCode,
		TotalPrice:    ticket.TotalPrice,
		PaymentStatus: ticket.PaymentStatus,
		MovieTitle:    ticket.MovieTitle,
		Category:      ticket.Category,
		ShowDate:      ticket.ShowDate,
		ShowTime:      ticket.ShowTime,
		TicketCount:   ticket.TicketCount,
		SeatLabels:    ticket.SeatLabels,
	}, nil
}

func generateVirtualRek() int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	const min = int64(10_000_000_000_000)
	const max = int64(99_999_999_999_999)
	return min + r.Int63n(max-min+1)
}
