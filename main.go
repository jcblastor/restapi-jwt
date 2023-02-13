package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jcblastor/restapi-jwt/controllers"
	"github.com/jcblastor/restapi-jwt/initializers"
	"github.com/jcblastor/restapi-jwt/middleware"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()

	r.POST("/register", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
}
