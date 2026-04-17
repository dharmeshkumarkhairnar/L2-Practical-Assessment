package handler

import (
	"net/http"
	"practical-assessment/model"
	"practical-assessment/service"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GetOrder struct {
	Service service.GetOrder
	DB      *gorm.DB
}

func NewGetOrder(service service.GetOrder, db *gorm.DB) *GetOrder {
	return &GetOrder{
		Service: service,
		DB:      db,
	}
}

func (controller GetOrder) HandleGetOrders(ctx *gin.Context) {
	// var orders model.Orders

	userId := ctx.GetInt64("userId")
	userOrders, err := controller.Service.GetOrder(ctx.Request.Context(), userId)
	if err != nil {
		if strings.Contains(err.Error(), "database query error") {
			ctx.IndentedJSON(http.StatusBadRequest, model.ErrorAPIResponse{
				Message: model.ErrorMessage{Key: "server", ErrorMessage: "internal server error"},
				Error:   "get orders request failed",
			})
			return
		}
		if strings.Contains(err.Error(), "user does not have any orders") {
			ctx.IndentedJSON(http.StatusBadRequest, model.ErrorAPIResponse{
				Message: model.ErrorMessage{Key: "order", ErrorMessage: "user does not have any orders"},
				Error:   "no orders found for user",
			})
			return
		}

	}

	ctx.IndentedJSON(http.StatusOK, model.GetOrderSuccessful{
		Message: "fetched user orders",
		Orders:  userOrders,
	})
}
