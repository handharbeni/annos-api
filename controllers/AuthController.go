package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/handharbeni/annos-api/middlewares"
	"github.com/handharbeni/annos-api/shared"
)

type AuthController struct{}

func (this AuthController) HandleLogin(c *gin.Context) {
	userId := "123"
	username := "Beast"
	roles := []string{shared.RoleAdmin, shared.RoleProUser}

	// do user auth here

	//issue token
	token, err := middlewares.GenerateToken([]byte(middlewares.SigningKey), userId, username, roles)
	if err != nil {

	}
	c.JSON(200, token)
}
