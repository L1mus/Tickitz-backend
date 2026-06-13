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
// @Description  Retrieve checkout summary details (ticket quantity, seats, total price, user profile) and available payment options.
// @Tags         Transactions
// @Security     ApiKeyAuth
// @Accept       json
// @Produce      json
// @Param        id           path      int     true  "Booking ID"
// @Success      200          {object}  dto.PaymentPageResponse "Successfully retrieved payment information summary"
// @Failure      401          {object}  dto.ResponseError "Unauthorized - Login session expired or token invalid"
// @Failure      404          {object}  dto.ResponseError "Not Found - Booking record not found or does not belong to user"
// @Failure      500          {object}  dto.ResponseError "Internal Server Error"
// @Router       /transactions/payment/{id} [get]
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

// SubmitPayment
// @Summary      Submit Selected Payment Method
// @Description  Process the chosen payment method to initiate a transaction and generate payment details (e.g., Virtual Account).
// @Tags         Transactions
// @Security     ApiKeyAuth
// @Accept       json
// @Produce      json
// @Param        body         body      dto.SubmitPaymentRequest  true  "Payload containing booking ID and selected payment method ID"
// @Success      200          {object}  dto.TransactionModalResponse "Payment method submitted successfully"
// @Failure      400          {object}  dto.ResponseError "Bad Request - Missing or invalid JSON body parameters"
// @Failure      401          {object}  dto.ResponseError "Unauthorized"
// @Failure      403          {object}  dto.ResponseError "Forbidden - Access denied, this booking does not belong to the authenticated user"
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
// @Summary      Simulate/Confirm Payment Complete
// @Description  Verify and process transaction clearance (simulates payment settlement).
// @Tags         Transactions
// @Security     ApiKeyAuth
// @Accept       json
// @Produce      json
// @Param        body         body      dto.ConfirmPaymentRequest  true  "Payload containing transaction ID and booking ID to be confirmed"
// @Success      200          {object}  dto.TicketResultResponse "Payment settlement verified and confirmed successfully"
// @Failure      400          {object}  dto.ResponseError "Bad Request"
// @Failure      401          {object}  dto.ResponseError "Unauthorized"
// @Failure      406          {object}  dto.ResponseError "Not Acceptable - The transaction has already been paid"
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
// @Summary      Get Digital Ticket Invoice
// @Description  Fetch the finalized digital ticket receipt details after a successful transaction payment.
// @Tags         Transactions
// @Security     ApiKeyAuth
// @Accept       json
// @Produce      json
// @Param        id           path      int     true  "Transaction ID"
// @Success      200          {object}  dto.TicketResultResponse "Successfully retrieved final digital ticket information"
// @Failure      401          {object}  dto.ResponseError "Unauthorized"
// @Failure      500          {object}  dto.ResponseError "Internal Server Error"
// @Router       /transactions/result/{id} [get]
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
// @Summary      Generate Ticket QR Code Image
// @Description  Generates and serves a raw binary PNG QR code image linked to the transaction for entry scanning verification.
// @Tags         Transactions
// @Security     ApiKeyAuth
// @Produce      image/png
// @Param        id           path      int     true  "Transaction ID"
// @Success      200          {file}    binary  "Returns raw binary PNG image of the QR Code"
// @Failure      500          {object}  dto.ResponseError "Internal Server Error - Failed to generate QR code image output"
// @Router       /transactions/qr/{id} [get]
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
// @Summary      Simulate DOKU Sandbox Checkout Invoice
// @Description  Triggers a direct backend-to-backend API request to DOKU Core Sandbox System to simulate a payment invoice creation.
// @Tags         DOKU Integration
// @Security     ApiKeyAuth
// @Accept       json
// @Produce      json
// @Success      200          {object}  map[string]interface{} "Successfully received simulated checkout payload from DOKU Core"
// @Failure      500          {object}  string "Internal Server Error - Failed communicating with external payment gateway API"
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
// @Summary      DOKU Webhook Notification Handler Listener
// @Description  Automated HTTP callback listener endpoint designed to receive status notification updates directly pushed by DOKU gateway systems.
// @Tags         DOKU Integration
// @Accept       json
// @Produce      json
// @Param        body         body      map[string]interface{}  true  "Raw payload notification structure sent from DOKU"
// @Success      200          {object}  map[string]interface{} "Returns status verification token acknowledgement back to DOKU"
// @Failure      400          {object}  map[string]interface{} "Bad Request - Invalid payload structure"
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
