package routes

import (
	controller "app/controllers"
	"app/middleware"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.RouterGroup) {
	router.Use(middleware.Authenticate())
	router.GET("/users/:user_id", controller.GetUser())
	router.GET("/users", controller.GetUsers())

}
