package routes

import (
	"github.com/VoidArtanis/go-rest-boilerplate/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterPublicRoutes(r *gin.Engine) {

	r.GET("/publicmessage", controllers.GetPublicText)
}
