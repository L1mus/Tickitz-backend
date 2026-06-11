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

func NewTransactionService(
	transactionRepo *repository.TransactionRepository,
	OrderRepo *repository.OrderRepository,
	db *pgxpool.Pool,
) *TransactionService {
	return &TransactionService{
		transactionRepo: transactionRepo,
		orderRepo:       OrderRepo,
		db:              db,
	}
}

func (s *TransactionService) GetPaymentPage() {
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
}

func (s *TransactionService) SubmitPayment() {
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
}

func (s *TransactionService) CheckPayment() {
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
}
