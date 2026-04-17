package handler

import (
	"errors"
	"net/http"

	// "practical-assessment/constant"
	"practical-assessment/model"
	"practical-assessment/service"

	// "practical-assessment/utils/validation"
	"strings"
	// "time"

	"github.com/gin-gonic/gin"
	// "github.com/go-playground/validator/v10"
)

type LoginHandler struct {
	loginService *service.LoginService
}

func NewLoginHandler(loginService *service.LoginService) *LoginHandler {
	return &LoginHandler{loginService: loginService}
}

func (controller *LoginHandler) LoginUserHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var bffRequest model.BFFLoginRequest
		var bffResponse model.BFFLoginResponse

		if err := ctx.ShouldBind(&bffRequest); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, "Unexpected Value Error")
			return
		}

		// err := validation.GetBFFValidation().Struct(&model.Users{})
		// errMsg, _ := validation.FormatValidationErrors(err.(validator.FieldError))
		// if errMsg != nil {
		// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorMessage{
		// 		Key:      "Validations",
		// 		ErrorMsg: "Error in validations",
		// 	})
		// 	return
		// }

		// username := ctx.GetString(constant.FieldName)
		// token := ctx.GetString(constant.Token)
		// expiry := ctx.GetInt64(constant.Expiry)
		// duration := expiry - time.Now().Unix()
		// ttl := time.Duration(duration)

		token, err := controller.loginService.LoginUserService(ctx, ctx.Request.Context(), bffRequest)
		if err != nil {
			if strings.Contains(err.Error(), "Service Error") {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, errors.New("Service Error"))
				return
			}

			if strings.Contains(err.Error(), "Password Mismatched, Login Failed") {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, errors.New("Password Mismatched, Login Failed"))
				return
			}

			if strings.Contains(err.Error(), "Redis set operation error") {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, errors.New("Redis set operation error"))
				return
			}

			if strings.Contains(err.Error(), "Error while generating token") {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, errors.New("Error while generating token"))
				return
			}

			ctx.AbortWithStatusJSON(http.StatusInternalServerError, "Login Failed")
			return
		}
		bffResponse.Message = "Login Successful"
		bffResponse.Token = token
		ctx.IndentedJSON(http.StatusOK, bffResponse)
	}
}
