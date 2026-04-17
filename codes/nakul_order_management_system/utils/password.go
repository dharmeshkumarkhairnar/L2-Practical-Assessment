package utils

import "golang.org/x/crypto/bcrypt"

func CompareHashPassword(hashPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}
