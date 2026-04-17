package handler

import (
	"net/http"
	"practical-assessment/model"
	"practical-assessment/service"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type LogoutHandler struct {
	logoutService *service.LogoutService
}

func NewUserlogout(logoutService *service.LogoutService) *LogoutHandler {
	return &LogoutHandler{
		logoutService: logoutService,
	}
}

func (uLogout *LogoutHandler) UserLogout(ctx *gin.Context) {
	logger := logrus.New()

	userId := ctx.GetInt64("userID")

	err := uLogout.logoutService.UserLogout(ctx.Request.Context(), logger, userId)

	if err != nil {
		if strings.Contains(err.Error(), "error in expiring the data in radis") {
			logger.Error("error in expiring the data in radis")
			errorMsg := model.ErrorMessage{
				Key:     "redis",
				Message: "error in expiring the data in radis",
			}
			ctx.IndentedJSON(http.StatusNotFound, &model.ErrorAPIResponse{
				ErrorMsg: errorMsg,
				Message:  "logout failed",
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
			Message:  "logout failed",
		})
		return

	}

	ctx.IndentedJSON(http.StatusOK, model.BFFLogoutResponse{
		Status: "successful",
	})
}
