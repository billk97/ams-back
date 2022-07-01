package controllers

import (
	utils "ams-back/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

var host = "http://acapy:8031/connections"

func CreateConnectionsController(r *gin.Engine) {
	if utils.Config.Aries != "" {
		host = utils.Config.Aries + "/connections"
	}
	router = r
	api := router.Group("api/connections")
	{
		api.POST("/create-invitation", createInvitation)
		api.GET("/", getConnections)
		api.DELETE("/:id", deleteConnections)
	}
}

func getConnections(c *gin.Context) {
	fmt.Println(host)
	fmt.Println("====")
	resp, err := http.Get(host)
	if err != nil {
		apiError := utils.NewApiError("REQUEST_FAILED", err, "details")
		c.JSON(500, apiError)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		apiError := utils.NewApiError("DESIRIALIAZATION_ERROR", err, "details")
		c.JSON(400, apiError)
		return
	}
	jsonString := string(body)
	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(jsonString), &jsonMap)
	c.JSON(200, jsonMap)
}

func createInvitation(c *gin.Context) {
	resp, err := http.Post(fmt.Sprintf("%s/create-invitation", host), "application/json", nil)
	if err != nil {
		apiError := utils.NewApiError("REQUEST_FAILED", err, "details")
		c.JSON(500, apiError)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		apiError := utils.NewApiError("DESIRIALIAZATION_ERROR", err, "details")
		c.JSON(400, apiError)
		return
	}
	jsonString := string(body)
	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(jsonString), &jsonMap)
	c.JSON(200, jsonMap)
}

func deleteConnections(c *gin.Context) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s", host, c.Param("id")), nil)
	if err != nil {
		apiError := utils.NewApiError("REQUEST_FAILED", err, "details")
		c.JSON(500, apiError)
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		apiError := utils.NewApiError("REQUEST_FAILED", err, "details")
		c.JSON(500, apiError)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		apiError := utils.NewApiError("DESIRIALIAZATION_ERROR", err, "details")
		c.JSON(400, apiError)
		return
	}
	jsonString := string(body)
	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(jsonString), &jsonMap)
	c.JSON(200, jsonMap)
}
