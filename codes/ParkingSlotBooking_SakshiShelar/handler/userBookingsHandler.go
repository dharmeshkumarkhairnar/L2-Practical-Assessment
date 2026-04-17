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

type UserBookingsHandler struct {
	userBookingsService *service.UserBookingsService
}

func NewUserBookingsHandler(userBookingsService *service.UserBookingsService) *UserBookingsHandler {
	return &UserBookingsHandler{
		userBookingsService: userBookingsService,
	}
}

func (handler *UserBookingsHandler) HandleUserBookings(ctx *gin.Context) {
	start := time.Now()
	logger := logrus.New()

	userId := ctx.GetInt("user_id")

	userBookings, err := handler.userBookingsService.UserBookings(ctx, userId)
	if err != nil {
		errorString := err.Error()

		//404 no bookings found for user
		if strings.Contains(errorString, constant.NoBookingsFound) {
			errorResponse := model.ErrorMessageResponse{
				Message: model.ErrorMessage{
					Key:      "user",
					ErrorMsg: constant.NoBookingsFound,
				},
				ErrorMessage: constant.StatusNotFound,
			}

			logger.WithFields(logrus.Fields{
				"user":    userId,
				"latency": time.Since(start).Seconds(),
			}).Info(constant.StatusNotFound)

			ctx.IndentedJSON(http.StatusNotFound, errorResponse)
			return
		}

		errorResponse := model.ErrorMessageResponse{
			Message: model.ErrorMessage{
				Key:      "server",
				ErrorMsg: constant.InternalServer,
			},
			ErrorMessage: constant.InternalServer,
		}

		logger.WithFields(logrus.Fields{
			"user":    userId,
			"latency": time.Since(start).Seconds(),
		}).Info(constant.InternalServer)

		ctx.IndentedJSON(http.StatusInternalServerError, errorResponse)
		return
	}

	var bookingsResponse []model.BFFUserBookingsResponse

	for _, val := range userBookings {
		booking := model.BFFUserBookingsResponse{
			Id:            val.Id,
			SlotNumber:    val.SlotNumber,
			VehicleNumber: val.VehicleNumber,
			Status:        val.Status,
		}
		bookingsResponse = append(bookingsResponse, booking)
	}

	logger.WithFields(logrus.Fields{
		"user":    userId,
		"latency": time.Since(start).Seconds(),
	}).Info(constant.InternalServer)

	ctx.IndentedJSON(http.StatusOK, bookingsResponse)

}
