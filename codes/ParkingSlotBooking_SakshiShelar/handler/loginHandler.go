package handler

import (
	"net/http"
	"practical-assessment/constant"
	"practical-assessment/model"
	"practical-assessment/service"
	"practical-assessment/utils/validations"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type LoginUserHandler struct {
	loginUserService *service.LoginUserService
}

func NewLoginUserHandler(loginUserService *service.LoginUserService) *LoginUserHandler {
	return &LoginUserHandler{
		loginUserService: loginUserService,
	}
}

func (handler *LoginUserHandler) HandleLoginUser(ctx *gin.Context) {
	start := time.Now()
	logger := logrus.New()

	var bffLoginUserRequest model.BFFLoginUserRequest

	if err := ctx.ShouldBind(&bffLoginUserRequest); err != nil {
		errorResponse := model.ErrorMessageResponse{
			Message: model.ErrorMessage{
				Key:      err.Error(),
				ErrorMsg: constant.UnexpectedValue,
			},
			ErrorMessage: constant.InvalidInputPayload,
		}

		logger.WithFields(logrus.Fields{
			"user":    bffLoginUserRequest.Email,
			"latency": time.Since(start).Seconds(),
		}).Info(constant.UnexpectedValue)

		ctx.IndentedJSON(http.StatusBadRequest, errorResponse)
		return
	}
	// 400 error
	if err := validations.GetBFFValidator().Struct(&bffLoginUserRequest); err != nil {
		validationErrors, _ := validations.FormatValidationErrors(err)

		logger.WithFields(logrus.Fields{
			"user":    bffLoginUserRequest.Email,
			"latency": time.Since(start).Milliseconds(),
		}).Info(constant.UnexpectedValue)

		ctx.JSON(http.StatusBadRequest, validationErrors)
		return
	}

	tokenString, err := handler.loginUserService.LoginUser(ctx, ctx.Request.Context(), bffLoginUserRequest)
	if err != nil {
		errorString := err.Error()

		//401 if password is wrong
		if strings.Contains(errorString, constant.PasswordMismatch) {
			errorResponse := model.ErrorMessageResponse{
				Message: model.ErrorMessage{
					Key:      "password",
					ErrorMsg: constant.PasswordMismatch,
				},
				ErrorMessage: constant.InvalidInputPayload,
			}

			logger.WithFields(logrus.Fields{
				"user":    bffLoginUserRequest.Email,
				"latency": time.Since(start).Seconds(),
			}).Info(constant.PasswordMismatch)

			ctx.IndentedJSON(http.StatusUnauthorized, errorResponse)
			return
		}

		//404 error
		if strings.Contains(errorString, constant.UserNotFound) {
			errorResponse := model.ErrorMessageResponse{
				Message: model.ErrorMessage{
					Key:      "email",
					ErrorMsg: constant.UserNotFound,
				},
				ErrorMessage: constant.AuthenticationFailed,
			}

			logger.WithFields(logrus.Fields{
				"user":    bffLoginUserRequest.Email,
				"latency": time.Since(start).Seconds(),
			}).Info(constant.UserNotFound)

			ctx.IndentedJSON(http.StatusNotFound, errorResponse)
			return
		}

		//500 error
		errorResponse := model.ErrorMessageResponse{
			Message: model.ErrorMessage{
				Key:      "server",
				ErrorMsg: constant.InternalServer,
			},
			ErrorMessage: constant.InternalServer,
		}

		logger.WithFields(logrus.Fields{
			"user":    bffLoginUserRequest.Email,
			"latency": time.Since(start).Seconds(),
		}).Info(constant.InternalServer)

		ctx.IndentedJSON(http.StatusInternalServerError, errorResponse)
		return

	}

	logger.WithFields(logrus.Fields{
		"user":    bffLoginUserRequest.Email,
		"latency": time.Since(start).Seconds(),
	}).Info(constant.UserLoginSuccess)

	ctx.IndentedJSON(http.StatusOK, model.BFFLoginUserResponse{
		Message: constant.UserLoginSuccess,
		Token:   tokenString,
	})
}
