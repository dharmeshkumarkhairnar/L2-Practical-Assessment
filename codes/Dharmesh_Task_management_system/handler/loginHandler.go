package handler

import (
	"net/http"
	"practical-assessment/model"
	"practical-assessment/service"
	"practical-assessment/utils/validations"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type LoginHandler struct {
	loginService *service.LoginService
}

func NewUserlogin(loginService *service.LoginService) *LoginHandler {
	return &LoginHandler{
		loginService: loginService,
	}
}

func (uLogin *LoginHandler) UserLogin(ctx *gin.Context) {
	logger := logrus.New()
	var bffLoginRequest model.BFFLoginRequest

	if err := ctx.ShouldBind(&bffLoginRequest); err != nil {
		logger.Error("json binding failed")
		errorMsg := model.ErrorMessage{
			Key:     "json binding",
			Message: "json binding failed",
		}
		ctx.IndentedJSON(http.StatusBadRequest, &model.ErrorAPIResponse{
			ErrorMsg: errorMsg,
			Message:  "login failed",
		})
		return
	}

	err := validations.GetValidator().Struct(&bffLoginRequest)

	if err != nil {
		errorMsgs:=validations.FormatValidationErrors(err)

		logger.Error("validation failed")
		ctx.IndentedJSON(http.StatusBadRequest, errorMsgs)
		return
	}

	token,err := uLogin.loginService.UserLogin(ctx, ctx.Request.Context(), logger, bffLoginRequest)

	if err != nil {
		if strings.Contains(err.Error(), "user not found") {
			logger.Error("user not found")
			errorMsg := model.ErrorMessage{
				Key:     "credentials",
				Message: "credentials are incorrect",
			}
			ctx.IndentedJSON(http.StatusNotFound, &model.ErrorAPIResponse{
				ErrorMsg: errorMsg,
				Message:  "login failed",
			})
			return
		} else if strings.Contains(err.Error(), "password is incorrect") {
			logger.Error("password is incorrect")
			errorMsg := model.ErrorMessage{
				Key:     "credentials",
				Message: "credentials are incorrect",
			}
			ctx.IndentedJSON(http.StatusUnauthorized, &model.ErrorAPIResponse{
				ErrorMsg: errorMsg,
				Message:  "login failed",
			})
			return
		} else if strings.Contains(err.Error(), "token generation failed") {
			logger.Error("token generation failed")
			errorMsg := model.ErrorMessage{
				Key:     "token",
				Message: "token generation failed",
			}
			ctx.IndentedJSON(http.StatusUnauthorized, &model.ErrorAPIResponse{
				ErrorMsg: errorMsg,
				Message:  "login failed",
			})
			return
		}

		logger.Error("internal server error")
		errorMsg := model.ErrorMessage{
			Key:     "server",
			Message: "internal server error",
		}
		ctx.IndentedJSON(http.StatusInternalServerError, &model.ErrorAPIResponse{
			ErrorMsg: errorMsg,
			Message:  "login failed",
		})
		return

	}

	ctx.IndentedJSON(http.StatusOK, model.BFFLoginResponse{
		Status: "successful",
		Token: token,
	})
}
