package handler

import (
	"net/http"
	"practical-assessment/constant"
	"practical-assessment/model"
	"practical-assessment/service"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AvailableSlotsHandler struct {
	availableSlotsService *service.AvailableSlotsService
}

func NewAvailableSlotsHandler(availableSlotsService *service.AvailableSlotsService) *AvailableSlotsHandler {
	return &AvailableSlotsHandler{
		availableSlotsService: availableSlotsService,
	}
}

func (handler *AvailableSlotsHandler) HandleAvailableSlots(ctx *gin.Context) {
	start := time.Now()
	logger := logrus.New()

	tokenString := ctx.GetString(constant.Token)

	slotNumbers, err := handler.availableSlotsService.AvailableSlots(ctx,tokenString)
	if err != nil {
		errorString := err.Error()

		//404 if no free slots
		if strings.Contains(errorString, constant.NoFreeSlots) {
			errorResponse := model.ErrorMessageResponse{
				Message: model.ErrorMessage{
					Key:      "slots",
					ErrorMsg: constant.NoFreeSlots,
				},
				ErrorMessage: constant.StatusNotFound,
			}

			logger.WithFields(logrus.Fields{
				"user":    "slots",
				"latency": time.Since(start).Seconds(),
			}).Info(constant.StatusNotFound)

			ctx.IndentedJSON(http.StatusNotFound, errorResponse)
			return
		}

		//500 if 1.redis init failed;2.marshal failed;3.unmarshal failed
		if strings.Contains(errorString, constant.RedisInitFailed) ||
			strings.Contains(errorString, constant.JsonMarshalFailed) ||
			strings.Contains(errorString, constant.JsonUnmarshalFailed) {

			var key, errorMsg string

			switch {
			case strings.Contains(errorString, constant.RedisInitFailed):
				key = "redis"
				errorMsg = constant.RedisInitFailed
			case strings.Contains(errorString, constant.JsonMarshalFailed):
				key = "json"
				errorMsg = constant.JsonMarshalFailed
			case strings.Contains(errorString, constant.JsonUnmarshalFailed):
				key = "json"
				errorMsg = constant.JsonUnmarshalFailed
			}

			errorResponse := model.ErrorMessageResponse{
				Message: model.ErrorMessage{
					Key:      key,
					ErrorMsg: errorMsg,
				},
				ErrorMessage: constant.InternalServer,
			}

			logger.WithFields(logrus.Fields{
				"user":    "redis",
				"latency": time.Since(start).Seconds(),
			}).Info(constant.InternalServer)

			ctx.IndentedJSON(http.StatusInternalServerError, errorResponse)
			return
		}
	}

	logger.WithFields(logrus.Fields{
		"latency": time.Since(start).Seconds(),
	}).Info(constant.SLotsFetchedSuccess)

	ctx.IndentedJSON(http.StatusOK, model.BFFAvailableSlotsReponse{
		SlotNumbers: slotNumbers,
	})
}
