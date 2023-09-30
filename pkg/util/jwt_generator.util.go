package util

import (
	"time"

	"github.com/Ahmad940/dropify/pkg/config"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken() (string, error) {
	// Create the Claims
	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * time.Duration(config.GetEnv().JWT_DURATION)).Unix(),
		"id":  "",
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	encodedToken, err := token.SignedString(config.GetEnv().JWT_SECRET)
	if err != nil {
		return "", err
	}
	return encodedToken, nil
}
