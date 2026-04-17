package handler

import (
	"net/http"
	"practical-assessment/constant"
	"practical-assessment/model"
	"practical-assessment/service"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type LogoutUserHandler struct {
	logoutUserService *service.LogoutUserService
}

func NewLogoutUserHandler(logoutUserService *service.LogoutUserService) *LogoutUserHandler {
	return &LogoutUserHandler{
		logoutUserService: logoutUserService,
	}
}

func (handler *LogoutUserHandler) HandleLogoutUser(ctx *gin.Context) {
	start := time.Now()
	logger := logrus.New()

	email := ctx.GetString(constant.Email)
	tokenString := ctx.GetString(constant.Token)

	err := handler.logoutUserService.LogoutUser(ctx, tokenString)
	if err != nil {
		errorString := err.Error()

		//401 if password is wrong
		if strings.Contains(errorString, constant.FailedToSetInRedis) {
			errorResponse := model.ErrorMessageResponse{
				Message: model.ErrorMessage{
					Key:      "password",
					ErrorMsg: constant.FailedToSetInRedis,
				},
				ErrorMessage: constant.UserUnauthorized,
			}

			logger.WithFields(logrus.Fields{
				"user":    email,
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
				"user":    email,
				"latency": time.Since(start).Seconds(),
			}).Info(constant.UserNotFound)

			ctx.IndentedJSON(http.StatusNotFound, errorResponse)
			return
		}

		//500 error
		if strings.Contains(errorString, constant.RedisInitFailed) {
			errorResponse := model.ErrorMessageResponse{
				Message: model.ErrorMessage{
					Key:      "redis",
					ErrorMsg: constant.RedisInitFailed,
				},
				ErrorMessage: constant.InternalServer,
			}

			logger.WithFields(logrus.Fields{
				"user":    "redis",
				"latency": time.Since(start).Seconds(),
			}).Info(constant.InternalServer)

			ctx.IndentedJSON(http.StatusInternalServerError, errorResponse)
			return
		}

		errorResponse := model.ErrorMessageResponse{
			Message: model.ErrorMessage{
				Key:      "server",
				ErrorMsg: constant.InternalServer,
			},
			ErrorMessage: constant.InternalServer,
		}

		logger.WithFields(logrus.Fields{
			"user":    email,
			"latency": time.Since(start).Seconds(),
		}).Info(constant.InternalServer)

		ctx.IndentedJSON(http.StatusInternalServerError, errorResponse)
		return

	}

	logger.WithFields(logrus.Fields{
		"user":    email,
		"latency": time.Since(start).Seconds(),
	}).Info(constant.UserLogoutSuccess)

	ctx.IndentedJSON(http.StatusOK, model.BFFLogoutUserResponse{
		Message: constant.UserLogoutSuccess,
	})
}
