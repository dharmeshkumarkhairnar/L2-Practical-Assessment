package validation

import (
	"practical-assessment/constant"
	"regexp"

	"github.com/go-playground/validator/v10"
)

var (
	uprReg   = regexp.MustCompile(constant.UpperCaseRegex)
	lwrRegex = regexp.MustCompile(constant.LowerCaseRegex)
	spclReg  = regexp.MustCompile(constant.SpecialSymRegex)
	digReg   = regexp.MustCompile(constant.DigitRegex)
)

// func StrongPassword(password string) bool {
// 	matched := uprReg.MatchString(password)
// }

func FormatValidationErrors(err error) []string {
	var ValidationErrs []string

	for _, err := range err.(validator.ValidationErrors) {
		if err.Field() == "password" {
			if err.Tag() == "required" {
				ValidationErrs = append(ValidationErrs, "password is required\n")
			}
			ValidationErrs = append(ValidationErrs, "password must be atleast 8 characters long and must contain atleast one lowercase letter, one uppercase letter, one special symbol and a digit\n")
		}

		if err.Field() == "email" {
			if err.Tag() == "required" {
				ValidationErrs = append(ValidationErrs, "email is required\n")
			}
			ValidationErrs = append(ValidationErrs, "email is invalid\n")
		}

		if err.Field() == "product_name" {
			if err.Tag() == "required" {
				ValidationErrs = append(ValidationErrs, "product_name is required\n")
			}
			ValidationErrs = append(ValidationErrs, "product name must be a string\n")
		}

		if err.Field() == "quantity" {
			if err.Tag() == "required" {
				ValidationErrs = append(ValidationErrs, "quantity is required\n")
			}
			ValidationErrs = append(ValidationErrs, "quantity must be a positive number\n")
		}

		if err.Field() == "price" {
			if err.Tag() == "required" {
				ValidationErrs = append(ValidationErrs, "price is required\n")
			}
			ValidationErrs = append(ValidationErrs, "price must be a positive number\n")
		}

	}

	return ValidationErrs
}
