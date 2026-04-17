package handler

import (
	"net/http"
	"practical-assessment/model"
	"practical-assessment/service"
	"strings"

	"github.com/gin-gonic/gin"
)

type DelExpenseHandler struct {
	delExpenseService *service.DelExpenseService
}

func NewDelExpenseHandler(delExpenseService *service.DelExpenseService) *DelExpenseHandler {
	return &DelExpenseHandler{delExpenseService: delExpenseService}
}

func (controller *DelExpenseHandler) DelUserExpenseHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var bffRequest model.BFFDelExpenseRequest
		var bffResponse model.BFFDelExpenseResponse

		if err := ctx.ShouldBind(&bffRequest); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, "Unexpected Value Error")
			return
		}

		err := controller.delExpenseService.DelUserExpenseService(ctx, ctx.Request.Context(), bffRequest)
		if err != nil {
			if strings.Contains(err.Error(), "No data Found") {
				ctx.AbortWithStatusJSON(http.StatusNotFound, "No data Found")
				return
			}
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, "GetExpense Failed")
			return
		}

		bffResponse.Message = "Expense Deleted Successfully"

		ctx.IndentedJSON(http.StatusOK, bffResponse)
	}
}
