package handler

import (
	"net/http"
	"practical-assessment/model"
	"practical-assessment/service"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type LogoutHandler struct {
	Service     service.LogoutService
	DB          *gorm.DB
	RedisClient *redis.Client
}

func NewLogoutHandler(service service.LogoutService, db *gorm.DB, redisClient *redis.Client) *LogoutHandler {
	return &LogoutHandler{
		Service:     service,
		DB:          db,
		RedisClient: redisClient,
	}
}

func (controller *LogoutHandler) HandleLogout(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	token, exists := ctx.Get("token")
	if !exists {
		ctx.IndentedJSON(http.StatusUnauthorized, "no active session")
		return
	}

	err := controller.Service.Logout(ctx.Request.Context(), userId, token.(string))
	if err != nil {
		if strings.Contains(err.Error(), "redis error") {
			ctx.IndentedJSON(http.StatusInternalServerError, model.ErrorAPIResponse{
				Message: model.ErrorMessage{Key: "server", ErrorMessage: "internal server error"},
				Error:   "failed to logout user",
			})
			return
		}
		ctx.IndentedJSON(http.StatusBadRequest, model.ErrorAPIResponse{
			Message: model.ErrorMessage{Key: "", ErrorMessage: "error logging out"},
			Error:   "failed to logout user",
		})
		return
	}
	ctx.IndentedJSON(http.StatusOK, model.LogoutSuccessModel{
		Message: "user logged out successfully",
	})
}
