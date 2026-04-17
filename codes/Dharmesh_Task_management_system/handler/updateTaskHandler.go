package handler

import (
	"net/http"
	"practical-assessment/model"
	"practical-assessment/service"
	"practical-assessment/utils/validations"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UpdateTasksHandler struct {
	updateTaskService *service.UpdateTaskService
}

func NewUpdateTasksHandler(updateTaskService *service.UpdateTaskService) *UpdateTasksHandler {
	return &UpdateTasksHandler{
		updateTaskService: updateTaskService,
	}
}

func (ut *UpdateTasksHandler) UpdateTasks(ctx *gin.Context) {
	logger := logrus.New()

	taskIdString := ctx.Param("id")
	taskId, _ := strconv.Atoi(taskIdString)
	userId := ctx.GetInt64("userID")

	var bffUpdateTaskRequest model.BFFUpdateTaskRequest

	if err := ctx.ShouldBind(&bffUpdateTaskRequest); err != nil {
		logger.Error("json binding failed")
		errorMsg := model.ErrorMessage{
			Key:     "json binding",
			Message: "json binding failed",
		}
		ctx.IndentedJSON(http.StatusBadRequest, &model.ErrorAPIResponse{
			ErrorMsg: errorMsg,
			Message:  "update failed",
		})
		return
	}

	err := validations.GetValidator().Struct(&bffUpdateTaskRequest)

	if err != nil {
		errorMsgs := validations.FormatValidationErrors(err)

		logger.Error("validation failed")
		ctx.IndentedJSON(http.StatusBadRequest, errorMsgs)
		return
	}

	err = ut.updateTaskService.UpdateTasks(ctx, ctx.Request.Context(), logger, int64(taskId), bffUpdateTaskRequest,userId)

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

		logger.Error("update failed, database error")
		errorMsg := model.ErrorMessage{
			Key:     "DB",
			Message: err.Error(),
		}
		ctx.IndentedJSON(http.StatusInternalServerError, &model.ErrorAPIResponse{
			ErrorMsg: errorMsg,
			Message:  "update failed",
		})
		return

	}

	ctx.IndentedJSON(http.StatusOK, &model.BFFUpdateTaskResponse{
		Message: "Update  successfull",
	})
}
