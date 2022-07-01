package controllers

import (
	"ams-back/dtos"
	utils "ams-back/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateIssueCredentialsWebhookController(r *gin.Engine) {
	router = r
	api := router.Group("api/webhook/topic/issue_credential_v2_0")
	{
		api.POST("/", issue_credential)
	}
}

func issue_credential(c *gin.Context) {
	dto := dtos.IssueCredentialWebhookDTO{}
	err := c.BindJSON(&dto)
	if err != nil {
		c.JSON(400, err)
		fmt.Printf("err: %s", err.Error())
		return
	}
	utils.PrintWebhook(dto)
}
