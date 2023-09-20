package main

import (
	"github.com/gin-gonic/gin"
	// "net/http"

	"app/middleware"
	"app/routes"
)

func main() {
	app := gin.Default()

	authRoutes := app.Group("/")
	userRouter := app.Group("/")

	userRouter.Use(middleware.Authenticate())

	routes.AuthRoutes(authRoutes)
	routes.UserRouter(userRouter)

	app.Run(":8080")
}
