package utils

import (
	"errors"
	"practical-assessment/constant"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWTSecretKey = []byte(constant.SecretKey)

func GenerateToken(email string, purpose string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":     email,
		"purpose": purpose,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(time.Minute * 3).Unix(),
	})

	tokenString, err := token.SignedString(JWTSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, validSignMethod := token.Method.(*jwt.SigningMethodHMAC); !validSignMethod {
			return nil, errors.New("Invalid Signing Method")
		}
		return JWTSecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("Invalid token")
}

func GetEmailFromToken(tokenString string) (string, error) {
	claims, err := ValidateToken(tokenString)
	if err != nil {
		return "", err
	}

	email, emailFetchOk := claims["sub"].(string)
	if !emailFetchOk {
		return "", errors.New("Email missing or invalid")
	}

	return email, nil
}

func GetExpiryFromToken(tokenString string) (int64, error) {
	claims, err := ValidateToken(tokenString)
	if err != nil {
		return 0, err
	}

	tokenExpiry, tokenExpiryOk := claims["exp"].(float64)
	if !tokenExpiryOk {
		return 0, errors.New("Token expiry missing or invalid")
	}

	return int64(tokenExpiry), nil
}
