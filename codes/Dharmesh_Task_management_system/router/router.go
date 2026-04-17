package router

import (
	"practical-assessment/handler"
	"practical-assessment/middleware"
	"practical-assessment/repository"
	"practical-assessment/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"POST", "GET", "PUT", "DELETE"},
		AllowHeaders:    []string{"Authorization", "Content-type", "Origin"},
	}))

	loginRepo := repository.Newlogin()
	loginService := service.NewUserlogin(loginRepo)
	loginHandler := handler.NewUserlogin(loginService)

	logoutService := service.NewUserLogout()
	logoutHandler := handler.NewUserlogout(logoutService)

	createTaskRepo:=repository.NewCreateTask()
	createTaskService:=service.NewCreateTask(createTaskRepo)
	createTaskHandler:=handler.NewCreateTask(createTaskService)

	getTasksRepo:=repository.NewGetTasks()
	getTasksService:=service.NewGetTaskService(getTasksRepo)
	getTasksHandler:=handler.NewGetTasksHandler(getTasksService)

	deleteTasksRepo:=repository.NewDeleteTasks()
	deleteTasksService:=service.NewDeleteTaskService(deleteTasksRepo)
	deleteTasksHandler:=handler.NewDeleteTasksHandler(deleteTasksService)

	updateTasksRepo:=repository.NewUpdateTasks()
	updateTasksService:=service.NewUpdateTaskService(updateTasksRepo)
	updateTasksHandler:=handler.NewUpdateTasksHandler(updateTasksService)

	router.POST("/auth/login", loginHandler.UserLogin)
	router.POST("/auth/logout", middleware.AuthMiddleware(), logoutHandler.UserLogout)
	router.POST("/tasks", middleware.AuthMiddleware(), createTaskHandler.CreateTask)
	router.GET("/tasks", middleware.AuthMiddleware(), getTasksHandler.GetTasks)
	router.DELETE("/tasks/:id", middleware.AuthMiddleware(), deleteTasksHandler.DeleteTasks)
	router.PUT("/tasks/:id", middleware.AuthMiddleware(), updateTasksHandler.UpdateTasks)

	return router
}
