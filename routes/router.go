package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/handharbeni/annos-api/controllers"
	"github.com/handharbeni/annos-api/middlewares"
)

func InitRouter(engine *gin.Engine) {
	InitMiddleware(engine)
	authController := new(controllers.AuthController)
	engine.POST("/login", authController.HandleLogin)
	RegisterProtectedRoutes(engine)
	RegisterPublicRoutes(engine)
	RegisterUtilityRoutes(engine)
}

func InitMiddleware(engine *gin.Engine) {
	engine.Use(middlewares.CORSMiddleware())
}
