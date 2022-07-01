package controllers

import (
	dtos "ams-back/dtos"
	utils "ams-back/utils"

	"github.com/gin-gonic/gin"
)

func CreateIssueCreddentialController(r *gin.Engine) {
	api := r.Group("api/issue-credentials")
	{
		api.POST("/")
	}
}

func handleIssueCreddential(c *gin.Context) {
	dto := dtos.IssueCredentialDTO{}
	serializationError := c.Bind(&dto)
	if serializationError != nil {
		apiError := utils.NewApiError(
			"INVALID_INPUT",
			serializationError,
			"Error serializing Struc of type IssueCredentiaslDto",
		)
		apiError.Enhance(c)
		c.JSON(400, &apiError)
		return
	}
	// make request to agent
}
