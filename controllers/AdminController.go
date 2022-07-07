package controllers

import (
	"ams-back/dtos"
	"ams-back/usecases"
	"ams-back/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateAdminController(r *gin.Engine) {
	api := r.Group("api/admin")
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
		apiErr := utils.NewApiError(
			"INVALID_CREDENTIALS",
			loginError,
			fmt.Sprintf("Credentials are invalid!"),
		)
		apiErr.Enhance(c)
		c.JSON(401, apiErr)
		return
	}
	c.JSON(200, &access)
}
