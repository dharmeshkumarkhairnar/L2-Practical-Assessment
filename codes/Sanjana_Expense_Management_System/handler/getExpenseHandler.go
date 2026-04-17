package handler

import (
	"net/http"
	"practical-assessment/model"
	"practical-assessment/service"
	"strings"

	"github.com/gin-gonic/gin"
)

type GetExpenseHandler struct {
	getExpenseService *service.GetExpenseService
}

func NewGetExpenseHandler(getExpenseService *service.GetExpenseService) *GetExpenseHandler {
	return &GetExpenseHandler{getExpenseService: getExpenseService}
}

func (controller *GetExpenseHandler) GetUserExpenseHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var bffRequest model.BFFGetExpenseRequest
		var bffResponse model.BFFGetExpenseResponse

		if err := ctx.ShouldBind(&bffRequest); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, "Unexpected Value Error")
			return
		}

		data, err := controller.getExpenseService.GetUserExpenseService(ctx, ctx.Request.Context(), bffRequest)
		if err != nil {
			if strings.Contains(err.Error(), "No data Found") {
				ctx.AbortWithStatusJSON(http.StatusNotFound, "No data Found")
				return
			}
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, "GetExpense Failed")
			return
		}

		bffResponse.Amount = data.Amount
		bffResponse.Description = data.Description
		bffResponse.CreatedAt = data.CreatedAt

		ctx.IndentedJSON(http.StatusOK, bffResponse)
	}
}
