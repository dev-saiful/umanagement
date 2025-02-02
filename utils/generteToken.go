package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Payload struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

var JwtKey = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWT(email,role string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	payload := &Payload{
		Email: email,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyJWT(tokenString string) (*Payload, error) {
	if tokenString == "" {
		return nil, errors.New("token cannot be empty")
	}

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &Payload{}, func(token *jwt.Token) (interface{}, error) {
		//  validate the signing method
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return JwtKey, nil
	})

	if err != nil {
		return nil, errors.New("failed to parse token: " + err.Error())
	}

	// Validate the token
	payload, ok := token.Claims.(*Payload)
	if ok && token.Valid {
		return payload, nil
	}
	return nil, errors.New("invalid token")
}
