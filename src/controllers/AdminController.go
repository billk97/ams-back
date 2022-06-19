package controllers

import (
	"ams-back/src/dtos"
	"ams-back/src/usecases"
	"ams-back/src/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

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
	}
	c.JSON(200, access)
}
