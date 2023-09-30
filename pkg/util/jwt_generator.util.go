package util

import (
	"fmt"
	"time"

	"github.com/Ahmad940/dropify/pkg/config"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(id string) (string, error) {
	// Create the Claims
	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * time.Duration(config.GetEnv().JWT_DURATION)).Unix(),
		"id":  id,
	}

	fmt.Println("Key", config.GetEnv().JWT_SECRET)

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	encodedToken, err := token.SignedString([]byte(config.GetEnv().JWT_SECRET))
	if err != nil {
		return "", err
	}
	return encodedToken, nil
}
