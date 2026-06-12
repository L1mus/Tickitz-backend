package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	apperror "github.com/L1mus/Tickitz-backend/internal/appError"
	"github.com/L1mus/Tickitz-backend/internal/config"
	"github.com/L1mus/Tickitz-backend/internal/dto"
	"github.com/L1mus/Tickitz-backend/internal/response"
	"github.com/L1mus/Tickitz-backend/internal/service"
	"github.com/L1mus/Tickitz-backend/pkg"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/skip2/go-qrcode"
)

type TransactionController struct {
	transactionService *service.TransactionService
}

func NewTransactionController(transactionService *service.TransactionService) *TransactionController {
	return &TransactionController{
		transactionService: transactionService,
	}
}

// GetPaymentInformation
// @Summary      Get Payment Page Information
// @Description  Retrieve payment summary information (ticket data, total price, user data) along with payment method options..
// @Tags         Transactions
// @Security     ApiKeyAuth
// @Accept       json
// @Produce      json
// @Param        booking_id   query     int     true  "ID Booking ticket"
// @Success      200          {object}  dto.PaymentPageResponse "Success get information payment" | "booking data not found" | "user is not the same as the user who made the booking
// @Failure      401          {object}  dto.ResponseError "Unauthorized"
// @Failure      500          {object}  dto.ResponseError "Internal Server Error"
// @Router       /transactions/payment [get]
func (c TransactionController) GetPaymentInformation(ctx *gin.Context) {
	token, exist := ctx.Get("claims")
	if !exist {
		response.Error(ctx, http.StatusUnauthorized, "unauthorized, please login")
	}

	claims := token.(*pkg.Claims)

	bookingIdString := ctx.Param("id")
	bookingId, _ := strconv.Atoi(bookingIdString)
	res, err := c.transactionService.GetPaymentPage(ctx.Request.Context(), bookingId, claims.Id)
	if err != nil {
		if errors.Is(err, apperror.BookingNotFound) || errors.Is(err, apperror.UserIstMatch) {
			response.Error(ctx, http.StatusNotFound, err.Error())
			return
		}
		response.Error(ctx, http.StatusInternalServerError, "internal server error")
		return
	}
	response.Success(ctx, http.StatusOK, "Success get information payment", res)
}

// SubmitPayment godoc
// @Summary      Submit Selected Payment Method
// @Tags         Transactions
// @Security     ApiKeyAuth
// @Accept       json
// @Produce      json
// @Param        body         body      dto.SubmitPaymentRequest  true  "Payload booking ID and payment method ID"
// @Success      200          {object}  dto.TransactionModalResponse "Submit payment success"
// @Failure      400          {object}  dto.ResponseError "bad request"
// @Failure      401          {object}  dto.ResponseError "Unauthorized"
// @Failure      403          {object}  dto.ResponseError "forbidden, booking does not belong to user"
// @Failure      500          {object}  dto.ResponseError "Internal Server Error"
// @Router       /transactions/submit [post]
func (c *TransactionController) SubmitPayment(ctx *gin.Context) {
	token, exist := ctx.Get("claims")
	if !exist {
		response.Error(ctx, http.StatusUnauthorized, "unauthorized, please login")
	}
	claims := token.(*pkg.Claims)
	var payload dto.SubmitPaymentRequest
	if err := ctx.ShouldBindWith(&payload, binding.JSON); err != nil {
		response.Error(ctx, http.StatusBadRequest, "bad request")
	}
	res, err := c.transactionService.SubmitPayment(ctx.Request.Context(), claims.Id, payload)
	if err != nil {
		if errors.Is(err, apperror.ForbiddenBooking) {
			response.Error(ctx, http.StatusForbidden, err.Error())
			return
		}
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, "Submit payment success", res)
}

// ConfirmPayment
// @Summary      Confirm Ticket Payment
// @Description  Confirm ticket payment (payment simulation complete).
// @Tags         Transactions
// @Security     ApiKeyAuth
// @Accept       json
// @Produce      json
// @Param        body         body      map[string]int  true  "example: {'transaction_id': 1, 'booking_id': 1}"
// @Success      200          {object}  dto.TicketResultResponse "Payment confirmed successfully"
// @Failure      400          {object}  dto.ResponseError "bad request"
// @Failure      401          {object}  dto.ResponseError "Unauthorized"
// @Failure      500          {object}  dto.ResponseError "Internal Server Error"
// @Router       /transactions/confirm [post]
func (c *TransactionController) ConfirmPayment(ctx *gin.Context) {
	_, exist := ctx.Get("claims")
	if !exist {
		response.Error(ctx, http.StatusUnauthorized, "please login first !")
		return
	}
	var payload dto.ConfirmPaymentRequest
	if err := ctx.ShouldBindWith(&payload, binding.JSON); err != nil {
		response.Error(ctx, http.StatusBadRequest, "bad request")
		return
	}

	res, err := c.transactionService.CheckPayment(ctx, payload.TransactionID, payload.BookingID)
	if err != nil {
		if errors.Is(err, apperror.TicketAlreadyPaid) {
			response.Error(ctx, http.StatusNotAcceptable, err.Error())
			return
		}
		response.Error(ctx, http.StatusInternalServerError, "internal server error")
		return
	}
	response.Success(ctx, http.StatusOK, "Payment confirmed successfully", res)
}

// GetResultTicket
// @Summary      Get Digital Ticket Result
// @Description  Retrieve final invoice data for digital tickets after successful payment.
// @Tags         Transactions
// @Security     ApiKeyAuth
// @Accept       json
// @Produce      json
// @Param        transaction_id  query   int     true  "ID Transaksi"
// @Success      200          {object}  dto.TicketResultResponse "Success get information ticket"
// @Failure      404          {object}  dto.ResponseError "Ticket not found"
// @Failure      500          {object}  dto.ResponseError "internal server error"
// @Router       /transactions/result [get]
func (c *TransactionController) GetResultTicket(ctx *gin.Context) {
	_, exist := ctx.Get("claims")
	if !exist {
		response.Error(ctx, http.StatusUnauthorized, "please login first !")
		return
	}
	transactionIDString := ctx.Param("id")
	transactionId, _ := strconv.Atoi(transactionIDString)
	res, err := c.transactionService.GetTicketResult(ctx.Request.Context(), transactionId)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, "Success get information ticket", res)
}

// GetQrCodeImage
// @Summary      Get QR Code Image PNG
// @Description  Generates a PNG format QR Code banner image based on the transaction ID for studio entry scanning needs..
// @Tags         Transactions
// @Security     ApiKeyAuth
// @Produce      image/png
// @Param        transaction_id  query   int     true  "ID Transaction"
// @Success      200          {file}    binary "File Image QR Code PNG"
// @Failure      500          {object}  dto.ResponseError "failed to generate QR Code"
// @Router       /transactions/qr [get]
func (c *TransactionController) GetQrCodeImage(ctx *gin.Context) {
	transactionIDString := ctx.Param("id")
	transactionId, _ := strconv.Atoi(transactionIDString)
	result, err := c.transactionService.GetTicketResult(ctx, transactionId)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	var QRPng []byte
	QRPng, err = qrcode.Encode(result.QRCode, qrcode.Medium, 256)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to generate qr code")
		return
	}
	ctx.Data(http.StatusOK, "image/png", QRPng)
}

// TryCheckOutDoku
// @Summary      Simulate DOKU Checkout
// @Description  Triggers an HTTP Request hit for invoice creation directly to DOKU's core system sandbox API.
// @Tags         DOKU Integration
// @Security     ApiKeyAuth
// @Accept       json
// @Produce      json
// @Success      200          {object}  map[string]interface{} "Response success from Core System DOKU"
// @Router       /transactions/checkout [post]
func (c *TransactionController) TryCheckOutDoku(ctx *gin.Context) {
	{
		requestID := "REQ-" + fmt.Sprintf("%d", time.Now().Unix())
		timestamp := time.Now().UTC().Format("2006-01-02T15:04:05Z")

		// Data transaksi latihan
		payload := map[string]interface{}{
			"order": map[string]interface{}{
				"invoice_number": "INV-" + fmt.Sprintf("%d", time.Now().Unix()),
				"amount":         150000,
			},
			"payment": map[string]interface{}{
				"payment_due_date": 60,
			},
		}

		jsonBody, _ := json.Marshal(payload)
		signature := config.GenerateSignature(config.ClientID, config.SecretKey, timestamp, requestID, string(jsonBody))

		// Eksekusi HTTP Request (Fetch) ke DOKU
		client := &http.Client{Timeout: 10 * time.Second}
		req, _ := http.NewRequest("POST", config.BaseURL+config.RequestPath, bytes.NewBuffer(jsonBody))

		req.Header.Set("Client-Id", config.ClientID)
		req.Header.Set("Request-Id", requestID)
		req.Header.Set("Request-Timestamp", timestamp)
		req.Header.Set("Signature", signature)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		defer resp.Body.Close()

		respBody, _ := io.ReadAll(resp.Body)
		// Mengirimkan respons balik dari DOKU ke Frontend Anda
		var result map[string]interface{}
		json.Unmarshal(respBody, &result)

		ctx.JSON(resp.StatusCode, result)
	}
}

// TryCallback
// @Summary      DOKU Webhook Notification Listener
// @Description  An automated endpoint listener that receives callbacks/notifications of transaction status updates from DOKU.
// @Tags         DOKU Integration
// @Accept       json
// @Produce      json
// @Param        body         body      map[string]interface{}  true  "Data payload from DOKU"
// @Success      200          {object}  map[string]interface{} "{'status': 'SUCCESS'}"
// @Router       /transactions/doku-callback [post]
func (c *TransactionController) TryCallback(ctx *gin.Context) {
	// DOKU akan mengirimkan data hasil transaksi ke sini
	var callbackData map[string]interface{}

	if err := ctx.ShouldBindJSON(&callbackData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "INVALID_REQUEST"})
		return
	}

	// Cetak di terminal backend untuk melihat data pembayaran simulator
	fmt.Println("=== NOTIFIKASI WEBHOOK DOKU DITERIMA ===")
	fmt.Printf("Data: %+v\n", callbackData)

	// WAJIB: Berikan respons balik ini agar DOKU tahu backend Anda sukses menerima data
	ctx.JSON(http.StatusOK, gin.H{"status": "SUCCESS"})
}
