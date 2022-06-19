package controllers

import (
	"ams-back/src/dtos"
	"ams-back/src/usecases"
	"ams-back/src/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func CreateAdminConntroller(r *gin.Engine) {
	router = r
	api := router.Group("api/admin")
	{
		api.POST("/login", AdminLogin)
	}
}

func AdminLogin(c *gin.Context) {
	dto := dtos.CredentialDto{}
	err := c.BindJSON(&dto)
	if err != nil {
		apiErr := utils.NewApiError(
			"INVALID_INPUT",
			err,
			fmt.Sprintf("Erro serializing json to credential struct"),
		)
		apiErr.Enhance(c)
		c.JSON(400, apiErr)
		return
	}
	access, loginError := usecases.Login(dto)
	if loginError != nil {
		loginError.Enhance(c)
		c.JSON(401, loginError)
		return
	}
	c.JSON(200, &access)
}