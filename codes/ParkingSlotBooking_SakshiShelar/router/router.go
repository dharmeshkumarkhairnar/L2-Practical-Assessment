package router

import (
	"practical-assessment/constant"
	"practical-assessment/handler"
	"practical-assessment/middleware"
	"practical-assessment/repository"
	"practical-assessment/service"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())

	loginRepository := repository.LoginUserRepository()
	loginService := service.NewLoginUserService(loginRepository)
	loginHandler := handler.NewLoginUserHandler(loginService)

	logoutService := service.NewLogoutUserService()
	logoutHandler := handler.NewLogoutUserHandler(logoutService)

	availableSlotsRepository := repository.AvailableSlotsRepository()
	availableSlotsService := service.NewAvailableSlotsService(availableSlotsRepository)
	availableSlotsHandler := handler.NewAvailableSlotsHandler(availableSlotsService)

	userBookingsRepository := repository.UserBookingsRepository()
	userBookingsService := service.NewUserBookingsService(userBookingsRepository)
	userBookingsHandler := handler.NewUserBookingsHandler(userBookingsService)

	cancelBookingsRepository := repository.CancelBookingsRepository()
	cancelBookingsService := service.NewCancelBookingsService(cancelBookingsRepository)
	cancelBookingsHandler := handler.NewCancelBookingsHandler(cancelBookingsService)

	authGroup := router.Group(constant.AuthPrefix)
	{
		authGroup.POST(constant.Login, loginHandler.HandleLoginUser)
		authGroup.POST(constant.Logout, middleware.AuthMiddleware(), logoutHandler.HandleLogoutUser)
	}

	router.GET(constant.AvailableSlots, middleware.AuthMiddleware(), availableSlotsHandler.HandleAvailableSlots)

	authGroup = router.Group(constant.BookingsPrefix)
	{
		authGroup.GET(constant.UserBookings, middleware.AuthMiddleware(), userBookingsHandler.HandleUserBookings)
		authGroup.DELETE(constant.CancelBookings, middleware.AuthMiddleware(), cancelBookingsHandler.HandleCancelBookings)
	}

	return router
}
