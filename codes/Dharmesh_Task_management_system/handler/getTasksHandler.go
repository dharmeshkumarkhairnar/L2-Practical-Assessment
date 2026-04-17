package handler

import (
	"net/http"
	"practical-assessment/model"
	"practical-assessment/service"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type GetTasksHandler struct {
	getTaskService *service.GetTaskService
}

func NewGetTasksHandler(getTaskService *service.GetTaskService) *GetTasksHandler {
	return &GetTasksHandler{
		getTaskService: getTaskService,
	}
}

func (gt *GetTasksHandler) GetTasks(ctx *gin.Context) {
	logger := logrus.New()

	userId := ctx.GetInt64("userID")

	status := ctx.Query("status")
	limit, _ := strconv.ParseInt(ctx.Query("limit"), 10, 64)
	offset, _ := strconv.ParseInt(ctx.Query("offset"), 10, 64)

	tasks, err := gt.getTaskService.GetTasks(ctx, ctx.Request.Context(), logger, userId, status, limit, offset)

	if err != nil {
		if strings.Contains(err.Error(), "tasks not found") {
			logger.Error("tasks not found")
			errorMsg := model.ErrorMessage{
				Key:     "tasks",
				Message: "tasks not found",
			}
			ctx.IndentedJSON(http.StatusNotFound, &model.ErrorAPIResponse{
				ErrorMsg: errorMsg,
				Message:  "no tasks available",
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
			Message:  "operation failed",
		})
		return

	}

	ctx.IndentedJSON(http.StatusOK, tasks)
}
