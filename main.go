package main

import (
	"github.com/gin-gonic/gin"
	"goJwt/handlers"
	"goJwt/initializers"
	"goJwt/middleware"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
	initializers.SyncDB()
}

func main() {
	r := gin.Default()
	r.POST("/sign-up", handlers.SignUp)
	r.POST("/sign-in", handlers.SignIn)
	r.GET("/validate", middleware.RequireAuth, handlers.Validate)
	r.Run()
}
