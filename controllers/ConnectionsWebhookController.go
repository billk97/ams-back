package controllers

import (
	"ams-back/dtos"
	utils "ams-back/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateWebhookController(r *gin.Engine) {
	router = r
	api := router.Group("api/webhook/topic/connections")
	{
		api.POST("/", conectionWebhook)
	}
}

func webhook(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(400, err)
		return
	}
	fmt.Println(jsonData)
}

func conectionWebhook(c *gin.Context) {
	dto := dtos.ConnectionDTO{}
	err := c.BindJSON(&dto)
	if err != nil {
		c.JSON(400, err)
		fmt.Printf("err: %s", err.Error())
		return
	}
	utils.PrintWebhook(dto)
	if dto.Rfc23State == "request-received" {
		err := requestReceived(&dto)
		if err != nil {
			c.JSON(400, err)
			return
		}
	}
	c.JSON(200, dto)
}

func requestReceived(dto *dtos.ConnectionDTO) *utils.ApiError {
	resp, err := http.Post(fmt.Sprintf("%s/%s/accept-request", host, dto.ConnectionId), "application/json", nil)
	if err != nil {
		return utils.NewApiError("REQUEST_FAILED", err, "details")
	}
	fmt.Printf("status: %d \n", resp.StatusCode)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return utils.NewApiError("DESIRIALIAZATION_ERROR", err, "details")
	}
	jsonString := string(body)
	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(jsonString), &jsonMap)
	return nil
}
