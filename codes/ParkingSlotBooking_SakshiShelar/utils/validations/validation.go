package validations

import (
	"fmt"
	"practical-assessment/constant"
	"practical-assessment/model"
	"strings"

	"github.com/go-playground/validator/v10"
)

var bffValidator *validator.Validate

func FormatValidationErrors(err error) ([]model.ErrorMessage, string) {
	var validationErrors []model.ErrorMessage
	var validationErrorsStr string

	for _, err := range err.(validator.ValidationErrors) {
		var errorMsg string
		fieldName := err.Field()
		if err.Tag() == "required" {
			fieldName = strings.ToLower(fieldName)
			errorMsg = fmt.Sprintf(constant.FieldRequired, fieldName)
		}

		validationErrors = append(validationErrors, model.ErrorMessage{
			Key:      fieldName,
			ErrorMsg: errorMsg,
		})
		validationErrorsStr += fieldName + " is invalid; "
	}

	return validationErrors, validationErrorsStr
}

func init() {
	bffValidator = validator.New()
}

func GetBFFValidator() *validator.Validate {
	return bffValidator
}
