package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/handharbeni/annos-api/controllers"
)

func RegisterPublicRoutes(r *gin.Engine) {

	r.GET("/publicmessage", controllers.GetPublicText)
}
