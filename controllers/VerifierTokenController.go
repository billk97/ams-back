package controllers

import (
	"ams-back/models"
	"ams-back/repos"
	"ams-back/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CreateVerifierTokenController(r *gin.Engine) {
	api := r.Group("api/token")
	{
		api.GET(":key", getToken)
		api.POST("", addToken)
	}
}

func getToken(c *gin.Context) {
	pathParam := c.Param("key")
	token, err := repos.FindToken(pathParam)
	if err != nil {
		e := utils.NewApiError(
			"INVALID_INPUT",
			err,
			fmt.Sprintf(""),
		)
		e.Enhance(c)
		c.JSON(400, e)
		return
	}
	c.JSON(200, &token)

}

func addToken(c *gin.Context) {
	token := models.VerifierToken{}
	err := c.Bind(&token)
	if err != nil {
		apiErr := utils.NewApiError(
			"INVALID_INPUT",
			err,
			fmt.Sprintf("Erro serializing json to token struct"),
		)
		apiErr.Enhance(c)
		c.JSON(400, apiErr)
		return
	}
	dbErr := repos.CreateToken(&token)
	if dbErr != nil {
		apiError := utils.NewApiError(
			"PERSIST_ENTITY_FAILED",
			dbErr,
			fmt.Sprintf("Could't persist entity of type permision"),
		)
		apiError.Enhance(c)
		c.JSON(400, apiError)
		return
	}
	c.JSON(200, &token)
}
