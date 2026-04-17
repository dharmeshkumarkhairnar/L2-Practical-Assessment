package utils

import (
	"errors"
	"practical-assessment/constant"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userId uint64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":       userId,
		"issued_at": time.Now().Unix(),
		"exp":       time.Now().Add(3 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(constant.AccessKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, error) {

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Wrong signing algorithm")
		}
		return []byte(constant.AccessKey), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func VerifyToken(token *jwt.Token) (jwt.MapClaims, error) {

	if !token.Valid {
		return nil, errors.New("token is expired")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, errors.New("claims mapping failed")
	}

	return claims, nil
}
