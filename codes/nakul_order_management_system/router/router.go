package router

import (
	"practical-assessment/handler"
	"practical-assessment/middleware"
	"practical-assessment/repository"
	"practical-assessment/service"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func GetRouter(db *gorm.DB, redisClient *redis.Client) *gin.Engine {
	router := gin.New()
	router.Use(middleware.LoggerMiddleware())
	router.Use(gin.Recovery())

	loginRepo := repository.NewloginRepository(db, redisClient)
	loginService := service.NewLoginService(loginRepo, db, redisClient)
	loginHandler := handler.NewLoginHandler(*loginService)

	logoutService := service.NewLogoutService(db, redisClient)
	logoutHandler := handler.NewLogoutHandler(*logoutService, db, redisClient)

	delOrderRepo := repository.NewdelOrderRepo(db)
	delOrderService := service.NewDelOrderService(delOrderRepo, db)
	delOrderHandler := handler.NewDelOrderHandler(*delOrderService, db)

	createOrderRepo := repository.NewCreateOrderRepo(db)
	createOrderService := service.NewCreateOrder(createOrderRepo, db)
	createOrderHandler := handler.NewcreateOrderHandler(*createOrderService, db)

	getOrderRepo := repository.NewGetOrders(db)
	getOrderService := service.NewGetOrder(getOrderRepo, db)
	getOrderHandler := handler.NewGetOrder(*getOrderService, db)

	updateOrderRepo := repository.NewUpdateOrderRepo(db)
	updateOrderServce := service.NewUpdateOrderSerivce(updateOrderRepo, db)
	updateOrderHandler := handler.NewUpdateOrderHandler(*updateOrderServce, db)

	router.POST("/auth/login", loginHandler.HandleLogin)
	router.POST("/auth/logout", middleware.AuthMiddleware(redisClient), logoutHandler.HandleLogout)
	router.POST("/orders/create", middleware.AuthMiddleware(redisClient), createOrderHandler.HandleCreateOrder)
	router.GET("/orders/list", middleware.AuthMiddleware(redisClient), getOrderHandler.HandleGetOrders)
	router.PUT("/orders/update/:id", middleware.AuthMiddleware(redisClient), updateOrderHandler.HandleUpdateOrder)
	router.DELETE("/orders/cancle/:id", middleware.AuthMiddleware(redisClient), delOrderHandler.HandleDeleteOrder)

	return router
}
