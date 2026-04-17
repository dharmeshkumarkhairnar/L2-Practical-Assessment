package handler

import (
	"fmt"
	"net/http"
	"practical-assessment/model"
	"practical-assessment/service"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DelOrderHandler struct {
	service service.DelOrderSerivce
	DB      *gorm.DB
}

func NewDelOrderHandler(service service.DelOrderSerivce, db *gorm.DB) *DelOrderHandler {
	return &DelOrderHandler{
		service: service,
		DB:      db,
	}
}

func (controller DelOrderHandler) HandleDeleteOrder(ctx *gin.Context) {
	id := ctx.Param("id")

	orderId, err := strconv.Atoi(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, "invalid id")
		return
	}
	userId := ctx.GetInt64("userId")

	err = controller.service.DeleteOrder(ctx.Request.Context(), uint(orderId), userId)
	if err != nil {
		if strings.Contains(err.Error(), "database query error") {
			ctx.IndentedJSON(http.StatusInternalServerError, model.ErrorAPIResponse{
				Message: model.ErrorMessage{Key: "server", ErrorMessage: "internal server error"},
				Error:   "failed to cancle order",
			})
			return
		}

		if strings.Contains(err.Error(), "order does not exist") || strings.Contains(err.Error(), "order not found") {
			fmt.Println("cancelling non existing order:", err)
			ctx.IndentedJSON(http.StatusBadRequest, model.ErrorAPIResponse{
				Message: model.ErrorMessage{Key: "order", ErrorMessage: "order does not exist"},
				Error:   "cannot cancle non-existing order",
			})
			return
		}
	}

	ctx.IndentedJSON(http.StatusOK, model.DeleteOrderResponse{
		Message: "order cancelled successfully",
	})
}
