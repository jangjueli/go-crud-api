package routes

import (
	"go-crud-api/config"
	"go-crud-api/controllers"
	"go-crud-api/repositories"
	"go-crud-api/services"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	repo := repositories.NewUserRepository(config.DB)
	service := services.NewUserService(repo)
	controller := controllers.NewUserController(service)

	users := router.Group("/api/user")
	{
		users.GET("", controller.GetUsers)
		users.GET("/:id", controller.GetUser)
		users.POST("", controller.CreateUser)
		users.PUT("/:id", controller.UpdateUser)
		users.DELETE("/:id", controller.DeleteUser)
	}
}
