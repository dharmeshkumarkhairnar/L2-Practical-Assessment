package handler

import (
	"net/http"
	"practical-assessment/model"
	"practical-assessment/service"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UpdateOrderHandler struct {
	service service.UpdateOrderSerivce
	DB      *gorm.DB
}

func NewUpdateOrderHandler(service service.UpdateOrderSerivce, db *gorm.DB) *UpdateOrderHandler {
	return &UpdateOrderHandler{
		service: service,
		DB:      db,
	}
}

func (controller UpdateOrderHandler) HandleUpdateOrder(ctx *gin.Context) {
	var updateReq model.UpdateOrderRequest
	userId := ctx.GetInt64("userId")
	id := ctx.Param("id")

	orderId, err := strconv.Atoi(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, "invalid id")
		return
	}

	if err := ctx.ShouldBindJSON(&updateReq); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, model.ErrorAPIResponse{
			Message: model.ErrorMessage{Key: "request body", ErrorMessage: "invalid request body"},
			Error:   "failed to update order",
		})
		return
	}

	err = controller.service.UpdateOrderService(ctx.Request.Context(), updateReq, uint(orderId), userId)
	if err != nil {
		if strings.Contains(err.Error(), "database query error") {
			ctx.IndentedJSON(http.StatusInternalServerError, model.ErrorAPIResponse{
				Message: model.ErrorMessage{Key: "server", ErrorMessage: "internal server error"},
				Error:   "failed to update order",
			})
			return
		}

		if strings.Contains(err.Error(), "order does not exist") {
			ctx.IndentedJSON(http.StatusBadRequest, model.ErrorAPIResponse{
				Message: model.ErrorMessage{Key: "order", ErrorMessage: "order does not exist"},
				Error:   "failed to update order",
			})
			return
		}
	}

	ctx.IndentedJSON(http.StatusOK, model.UpdateOrderSuccessful{
		Message: "order updated successfully",
	})
}
