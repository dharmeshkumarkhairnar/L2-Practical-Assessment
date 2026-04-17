package router

import (
	"practical-assessment/constant"
	"practical-assessment/handler"
	"practical-assessment/repository"
	"practical-assessment/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{constant.AllowOrigins},
		AllowMethods: []string{constant.GET, constant.POST},
		AllowHeaders: []string{constant.AllowOrigins, constant.ContextType, constant.Authorization},
	}))

	logrepo := repository.NewRepository()
	logservice := service.NewLoginService(logrepo)
	controller := handler.NewLoginHandler(logservice)
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", controller.LoginUserHandler())
	}

	getrepo := repository.NewGetExpenseRepository()
	getservice := service.NewGetExpenseService(getrepo)
	getcontroller := handler.NewGetExpenseHandler(getservice)

	delrepo := repository.NewDelExpenseRepository()
	delservice := service.NewDelExpenseService(delrepo)
	delcontroller := handler.NewDelExpenseHandler(delservice)

	v1Group := router.Group("/expenses")
	{
		v1Group.POST("/get", getcontroller.GetUserExpenseHandler())
		v1Group.POST("/del", delcontroller.DelUserExpenseHandler())
	}
	return router
}
