package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xerofenix/gin-jwt/controllers"
	"github.com/xerofenix/gin-jwt/initializers"
	"github.com/xerofenix/gin-jwt/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	initializers.SyncDatabse()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.File("./index.html")
	})
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.Run()
}
