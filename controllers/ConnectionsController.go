package controllers

import (
	"ams-back/dtos"
	"ams-back/repos"
	utils "ams-back/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

var host = "http://acapy:8031/connections"

func CreateConnectionsController(r *gin.Engine) {
	if utils.Config.Aries != "" {
		host = utils.Config.Aries + "/connections"
	}
	router = r
	api := router.Group("api/connections")
	{
		api.POST("/create-invitation/:uuid", createInvitation)
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
	employeeInvitationId := c.Param("uuid")
	if employeeInvitationId == "" {
		apiError := utils.NewApiError("INVITATION_NOT_FOUND", nil, "missing invitation")
		c.JSON(500, apiError)
		return
	}
	fmt.Printf("employeeInvitation: %s \n", employeeInvitationId)
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
	dto := dtos.CreateInvitationDTO{}
	json.Unmarshal(body, &dto)
	employee, err := repos.FindEmployeeByInvitation(employeeInvitationId)
	if err != nil && employee == nil {
		c.JSON(400, err)
		return
	}
	employee.Status = "INVITATION-OPENED"
	employee.DidConnectionId = dto.ConnectionId
	_, databaseError := repos.UpdateEmployee(employee)
	if databaseError != nil {
		c.JSON(400, databaseError)
		return
	}
	c.JSON(200, &dto)
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
