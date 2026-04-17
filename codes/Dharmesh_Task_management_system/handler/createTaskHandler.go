package handler

import (
	"net/http"
	"practical-assessment/model"
	"practical-assessment/service"
	"practical-assessment/utils/validations"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CreateTaskHandler struct {
	createTaskService *service.CreateTaskService
}

func NewCreateTask(createTaskService *service.CreateTaskService) *CreateTaskHandler {
	return &CreateTaskHandler{
		createTaskService: createTaskService,
	}
}

func (ct *CreateTaskHandler) CreateTask(ctx *gin.Context) {
	logger := logrus.New()
	var bffCreateTaskRequest model.BFFCreateTaskRequest
	userId := ctx.GetInt64("userID")

	if err := ctx.ShouldBind(&bffCreateTaskRequest); err != nil {
		logger.Error("json binding failed")
		errorMsg := model.ErrorMessage{
			Key:     "json binding",
			Message: "json binding failed",
		}
		ctx.IndentedJSON(http.StatusBadRequest, &model.ErrorAPIResponse{
			ErrorMsg: errorMsg,
			Message:  "login failed",
		})
		return
	}

	err := validations.GetValidator().Struct(&bffCreateTaskRequest)

	if err != nil {
		errorMsgs:=validations.FormatValidationErrors(err)

		logger.Error("validation failed")
		ctx.IndentedJSON(http.StatusBadRequest, errorMsgs)
		return
	}

	err=ct.createTaskService.CreateTask(ctx,ctx.Request.Context(),logger,bffCreateTaskRequest,userId)

	if err != nil {
		if strings.Contains(err.Error(), "error in adding data in DB") {
			logger.Error("error in adding data in DB")
			errorMsg := model.ErrorMessage{
				Key:     "DB",
				Message: "error in adding data in DB",
			}
			ctx.IndentedJSON(http.StatusNotFound, &model.ErrorAPIResponse{
				ErrorMsg: errorMsg,
				Message:  "task creation failed",
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
			Message:  "task creation failed",
		})
		return

	}

	ctx.IndentedJSON(http.StatusCreated, model.BFFCreateTaskResponse{
		Message: "successful",
	})
}
