package utils

import (
	"errors"
	// "log"
	"practical-assessment/model"
	// "strconv"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

var secretKey *model.JWT

func GenerateToken(username string) (string, string, error) {

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Minute * 2).Unix(),
	})

	accessTokenString, err := accessToken.SignedString([]byte("abc"))
	if err != nil {
		return "", "", err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Minute * 1).Unix(),
	})

	refreshTokenString, err := refreshToken.SignedString([]byte("abc"))
	if err != nil {
		return "", "", err
	}
	return accessTokenString, refreshTokenString, nil
}

// func ValidateToken(tokenString string) (string, error) {
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, errors.New("invalid signing method")
// 		}
// 		return []byte(secretKey.AccessSecreteKey), nil
// 	})
// 	log.Println(err)
// 	if err != nil {
// 		return "", err
// 	}

// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		expiry, ok := claims["exp"].(float64)
// 		if !ok {
// 			return "", errors.New("expiry missing in token")
// 		}
// 		exp := strconv.FormatFloat(expiry, 'f', -1, 64)
// 		return exp, nil
// 	}

// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		username, ok := claims["name"].(string)
// 		if !ok {
// 			return "", errors.New("username missing in token")
// 		}
// 		return username, nil
// 	}
// 	return "", errors.New("Invalid token")
// }

func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, validSignMethod := token.Method.(*jwt.SigningMethodHMAC); !validSignMethod {
			return nil, errors.New("Invalid Signing Method")
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("Invalid token")
}

func GetExpiry(tokenString string) (int64, error) {
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
