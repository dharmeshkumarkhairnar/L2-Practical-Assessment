package utils

import (
	"fmt"
	"practical-assessment/constant"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// will return accessToken, refreshToken, error
func CreateToken(userId int64) (string, string, error) {

	jti := uuid.New().String()

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":        userId,
		"iat":        time.Now().Unix(),
		"exp":        time.Now().Add(12 * time.Hour).Unix(),
		"jti":        jti,
	})

	accessTokenString, err := accessToken.SignedString([]byte(constant.AccessSecretKey))
	if err != nil {
		return "", "", fmt.Errorf("error creating token")
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":        userId,
		"iat":        time.Now().Unix(),
		"exp":        time.Now().Add(24 * 30 * time.Hour).Unix(),
		"jti":        jti,
	})

	refreshTokenString, err := refreshToken.SignedString([]byte(constant.AccessSecretKey))
	if err != nil {
		return "", "", fmt.Errorf("error creating token")
	}

	return accessTokenString, refreshTokenString, nil
}
