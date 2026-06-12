package pkg

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Id       int    `json:"id"`
	FullName string `json:"full_name"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func NewClaims(id int, fullName string, role string) *Claims {
	return &Claims{
		Id:       id,
		FullName: fullName,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    os.Getenv("JWT_ISSUER"),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
}

func (c *Claims) GenJWT() (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return "", errors.New("missing jwt secret")
	}
	uToken := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return uToken.SignedString([]byte(jwtSecret))
}

func (c *Claims) VerifyJWT(tokenString string) (*Claims, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return nil, errors.New("missing jwt secret")
	}

	jwtToken, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (any, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := jwtToken.Claims.(*Claims)
	if !ok || !jwtToken.Valid {
		return nil, errors.New("invalid token claims")
	}

	if claims.Issuer != os.Getenv("JWT_ISSUER") {
		return nil, jwt.ErrTokenInvalidIssuer
	}

	return claims, nil
}
