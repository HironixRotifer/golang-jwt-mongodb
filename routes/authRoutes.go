package routes

import (
	controller "app/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(userRouter *gin.RouterGroup) {
	userRouter.POST("user/signup", controller.Signup())
	userRouter.POST("user/signin", controller.Signin())
}
