package controllers

import (
	dtos "ams-back/dtos"
	usecases "ams-back/usecases"
	utils "ams-back/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func CreateAdminConntroller(r *gin.Engine) {
	router = r
	api := router.Group("api/admin")
	{
		api.POST("/login", adminLogin)
	}
}

func adminLogin(c *gin.Context) {
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
