package handler

import (
	"net/http"
	"practical-assessment/model"
	"practical-assessment/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type DeleteTasksHandler struct {
	deleteTaskService *service.DeleteTaskService
}

func NewDeleteTasksHandler(deleteTaskService *service.DeleteTaskService) *DeleteTasksHandler {
	return &DeleteTasksHandler{
		deleteTaskService: deleteTaskService,
	}
}

func (dt *DeleteTasksHandler) DeleteTasks(ctx *gin.Context) {
	logger := logrus.New()

	taskIdString := ctx.Param("id")
	taskId, _ := strconv.Atoi(taskIdString)
	userId := ctx.GetInt64("userID")

	err := dt.deleteTaskService.DeleteTasks(ctx, ctx.Request.Context(), logger, int64(taskId),userId)

	if err != nil {

		logger.Error("deletion failed, database error")
		errorMsg := model.ErrorMessage{
			Key:     "DB",
			Message: err.Error(),
		}
		ctx.IndentedJSON(http.StatusInternalServerError, &model.ErrorAPIResponse{
			ErrorMsg: errorMsg,
			Message:  "deletion failed",
		})
		return

	}

	ctx.IndentedJSON(http.StatusOK, &model.BFFDeleteTaskResponse{
		Message: "Deletion successfull",
	})
}
