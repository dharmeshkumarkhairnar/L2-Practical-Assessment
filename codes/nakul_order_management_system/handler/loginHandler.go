package handler

import (
	"fmt"
	"net/http"
	"practical-assessment/model"
	"practical-assessment/service"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type LoginHandler struct {
	Service     service.LoginService
	DB          *gorm.DB
	RedisClient *redis.Client
}

func NewLoginHandler(service service.LoginService) *LoginHandler {
	return &LoginHandler{
		Service: service,
	}
}

func (controller *LoginHandler) HandleLogin(ctx *gin.Context) {
	var loginRequest model.LoginRequest

	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, model.ErrorAPIResponse{
			Message: model.ErrorMessage{Key: "request body", ErrorMessage: "invalid request body"},
			Error:   "user login request failed",
		})
		return
	}

	//validation

	// validator.Validate.Struct(&loginRequest)

	token, _, err := controller.Service.Login(ctx.Request.Context(), loginRequest)
	if err != nil {
		fmt.Println(err)
		ctx.IndentedJSON(http.StatusBadRequest, model.ErrorAPIResponse{
			Message: model.ErrorMessage{Key: "", ErrorMessage: "invalid credentials"},
			Error:   "user login request failed",
		})
		return
	}
	ctx.IndentedJSON(http.StatusOK, model.LoginSuccessModel{
		Message:     "user logged in successfully",
		AccessToken: token,
	})
}
