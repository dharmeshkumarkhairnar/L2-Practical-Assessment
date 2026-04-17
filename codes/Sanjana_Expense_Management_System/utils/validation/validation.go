package validation

import (
	"fmt"
	"practical-assessment/constant"
	"practical-assessment/model"
	"regexp"

	"github.com/go-playground/validator/v10"
)

var bffValidator *validator.Validate

func FormatValidationErrors(err validator.FieldError) ([]model.ErrorMessage, string) {
	var ValidationErrors []model.ErrorMessage
	var ValidationErrorsStr string
	fieldName := err.Field()
	var errMsg string
	if err.Tag() == "required" {
		errMsg = fmt.Sprintf(constant.RequiredFieldError, fieldName)
	} else {
		switch err.Field() {
		case constant.FieldName:
			errMsg = fmt.Sprintf(constant.RequiredFieldError, constant.FieldName)
		case constant.FieldPassword:
			switch err.Tag() {
			case "required":
				errMsg = fmt.Sprintf(constant.RequiredFieldError, constant.FieldPassword)
			case "passwordFormat":
				errMsg = constant.PasswordFomatContraintError
			}
		}
		ValidationErrors = append(ValidationErrors, model.ErrorMessage{
			Key:      fieldName,
			ErrorMsg: errMsg,
		})

		ValidationErrorsStr += fieldName + " is invalid;"
	}
	return ValidationErrors, ValidationErrorsStr
}

func ValidatePassword(f1 validator.FieldLevel) bool {
	pass := f1.Field().String()
	matched, _ := regexp.MatchString(constant.PasswordRegexp, pass)
	return matched
}

func init() {
	bffValidator = validator.New()
	bffValidator.RegisterValidation("passwordFormat", ValidatePassword)
}

func GetBFFValidation() *validator.Validate {
	return bffValidator
}
