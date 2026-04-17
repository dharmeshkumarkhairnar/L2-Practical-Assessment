package handler

import (
	"net/http"
	"practical-assessment/model"
	"practical-assessment/service"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateOrderHandler struct {
	Service service.CreateOrder
	DB      *gorm.DB
}

func NewcreateOrderHandler(service service.CreateOrder, db *gorm.DB) *CreateOrderHandler {
	return &CreateOrderHandler{
		Service: service,
		DB:      db,
	}
}

func (controller CreateOrderHandler) HandleCreateOrder(ctx *gin.Context) {
	var createReq model.CreateOrderRequest
	userId := ctx.GetInt64("userId")

	if err := ctx.ShouldBindJSON(&createReq); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, model.ErrorAPIResponse{
			Message: model.ErrorMessage{Key: "request body", ErrorMessage: "invalid request body"},
			Error:   "create order request failed",
		})
		return
	}

	err := controller.Service.CreateOrderService(ctx.Request.Context(), createReq, userId)
	if err != nil {
		if strings.Contains(err.Error(), "database query error") {
			ctx.IndentedJSON(http.StatusInternalServerError, model.ErrorAPIResponse{
				Message: model.ErrorMessage{Key: "server", ErrorMessage: "internal server error"},
				Error:   "create order request failed",
			})
			return
		}

		ctx.IndentedJSON(http.StatusInternalServerError, model.ErrorAPIResponse{
			Message: model.ErrorMessage{Key: "server", ErrorMessage: "internal server error"},
			Error:   "create order request failed",
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, model.CreateOrderResponse{
		Message: "order created successfully",
	})
}
