package validations

import (
	"practical-assessment/model"
	"strings"

	"github.com/dlclark/regexp2"
	"github.com/go-playground/validator/v10"
)

var bffValidator *validator.Validate

func PasswordValidator(f1 validator.FieldLevel) bool {
	pattern := regexp2.MustCompile(`^(?=.*[A-Z])(?=.*[a-z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`, 0)
	matched, _ := pattern.MatchString(f1.Field().String())
	return matched
}

func FormatValidationErrors(err error) []*model.ErrorMessage {

	var errosMsgs []*model.ErrorMessage
	var key string
	var message string

	for _, err := range err.(validator.ValidationErrors) {
		fieldName := err.Field()
		tagName := err.Tag()

		if tagName == "required" {
			key = fieldName
			message = "this field is required"
		} else if tagName == "max" && fieldName=="Title" {
			key = fieldName
			message = "size should be less than 200 chars"
		} else if tagName == "max" && fieldName=="Description" {
			key = fieldName
			message = "size should be less than 2000 chars"
		} else if fieldName == "Email" {
			key = fieldName
			message = "error in email format"
		} else if fieldName == "Password" {
			key = fieldName
			message = "password must have at least one uppercase, one lowercase, Digit and special character"
		} else if fieldName == "Status" {
			key = fieldName
			message = "only Pending, InProgress and Completed allowed for this field"
		} else if fieldName == "Priority" {
			key = fieldName
			message = "only Low, Medium and High allowed for this field"
		}

		errosMsgs = append(errosMsgs, &model.ErrorMessage{
			Key:     key,
			Message: message,
		})
	}
	return errosMsgs
}


func GetValidator() *validator.Validate {
	return bffValidator
}

func ValidateStatus(f1 validator.FieldLevel) bool {
	status:=f1.Field().String()
	status=strings.ToUpper(status)
	switch status {
	case "PENDING","INPROGRESS","COMPLETED","":
		return true
	}
	return false
}

func ValidatePriority(f1 validator.FieldLevel) bool {
	priority:=f1.Field().String()
	priority=strings.ToUpper(priority)
	switch priority {
	case "LOW","HIGH","MEDIUM":
		return true
	}
	return false
}

func init() {
	bffValidator = validator.New()
	bffValidator.RegisterValidation("checkPassword", PasswordValidator)
	bffValidator.RegisterValidation("checkStatus", ValidateStatus)
	bffValidator.RegisterValidation("checkPriority", ValidatePriority)
}
