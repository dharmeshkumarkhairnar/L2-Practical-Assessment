package handler

import (
	"net/http"
	"practical-assessment/constant"
	"practical-assessment/model"
	"practical-assessment/service"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CancelBookingsHandler struct {
	cancelBookingsService *service.CancelBookingsService
}

func NewCancelBookingsHandler(cancelBookingsService *service.CancelBookingsService) *CancelBookingsHandler {
	return &CancelBookingsHandler{
		cancelBookingsService: cancelBookingsService,
	}
}

func (handler *CancelBookingsHandler) HandleCancelBookings(ctx *gin.Context) {
	start := time.Now()
	logger := logrus.New()

	bookingsId, _ := strconv.Atoi(ctx.Param("id"))
	userId := ctx.GetInt("user_id")

	err := handler.cancelBookingsService.CancelBookings(ctx, userId, bookingsId)
	if err != nil {
		errorString := err.Error()

		//404 if bookking not found
		if strings.Contains(errorString, constant.BookingNotFound) {
			errorResponse := model.ErrorMessageResponse{
				Message: model.ErrorMessage{
					Key:      "bookingID",
					ErrorMsg: constant.BookingNotFound,
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

		//401 if user access diff booking id
		if strings.Contains(errorString, constant.UserUnauthorized) {
			errorResponse := model.ErrorMessageResponse{
				Message: model.ErrorMessage{
					Key:      "bookingID",
					ErrorMsg: constant.UserIsUnauthorized,
				},
				ErrorMessage: constant.UserUnauthorized,
			}

			logger.WithFields(logrus.Fields{
				"user":    userId,
				"latency": time.Since(start).Seconds(),
			}).Info(constant.UserIsUnauthorized)

			ctx.IndentedJSON(http.StatusUnauthorized, errorResponse)
			return
		}

		//400 is already cancelled
		if strings.Contains(errorString, constant.AlreadyCancelled) {
			errorResponse := model.ErrorMessageResponse{
				Message: model.ErrorMessage{
					Key:      "bookingID",
					ErrorMsg: constant.AlreadyCancelled,
				},
				ErrorMessage: constant.NoUpdate,
			}

			logger.WithFields(logrus.Fields{
				"user":    userId,
				"latency": time.Since(start).Seconds(),
			}).Info(constant.AlreadyCancelled)

			ctx.IndentedJSON(http.StatusBadRequest, errorResponse)
			return
		}

		//500 if failed to cancel booking from server side
		if strings.Contains(errorString, constant.FailedToCancel) {
			errorResponse := model.ErrorMessageResponse{
				Message: model.ErrorMessage{
					Key:      "bookingID",
					ErrorMsg: constant.FailedToCancel,
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


	logger.WithFields(logrus.Fields{
		"user":    userId,
		"latency": time.Since(start).Seconds(),
	}).Info(constant.BookingCancelledSuccess)

	ctx.IndentedJSON(http.StatusOK, model.BFFCancelBookingsResponse{
		Message: constant.BookingCancelledSuccess,
	})
}
