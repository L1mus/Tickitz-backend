package config

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

const (
	BaseURL     = "https://api-sandbox.doku.com"
	ClientID    = "BRN-0257-1781100218257"
	SecretKey   = "SK-DA1psEG1xRcHLTNXQ2YI"
	RequestPath = "/checkout/v1/payment"
	APIkey      = "doku_key_sandbox_2cfdd1416de2439f94f9da8530262aba"
)

func GenerateSignature(clientID, secretKey, timestamp, requestID, requestBody string) string {
	// 	Lakukan hashing SHA-256 pada body request (minified)
	hasher := sha256.New()
	hasher.Write([]byte(requestBody))
	bodyHash := base64.StdEncoding.EncodeToString(hasher.Sum(nil))

	// Susun komponen string sesuai aturan DOKU (String-To-Sign)
	stringToSign := fmt.Sprintf("Client-Id:%s\nRequest-Id:%s\nRequest-Timestamp:%s\nRequest-Target:%s\nDigest:%s",
		clientID, requestID, timestamp, RequestPath, bodyHash)

	// Buat enkripsi HMAC-SHA256 menggunakan Secret Key
	key := []byte(secretKey)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(stringToSign))

	// 4. Encode hasil akhir ke format Base64
	return "HMACSHA256=" + base64.StdEncoding.EncodeToString(h.Sum(nil))
}
