package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/handharbeni/annos-api/controllers"
	"github.com/handharbeni/annos-api/middlewares"
)

func RegisterProtectedRoutes(r *gin.Engine) {

	authGroup := r.Group("/auth")

	authGroup.Use(middlewares.AuthHandler("admin"))
	{
		authGroup.GET("/getmessage", controllers.GetSecretText)
	}
}
