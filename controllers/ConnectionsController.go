package controllers

import (
	"ams-back/dtos"
	"ams-back/middlewares"
	"ams-back/repos"
	"ams-back/usecases"
	utils "ams-back/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

var connectionUrl = ""

func CreateConnectionsController(r *gin.Engine) {
	if utils.Config.Aries != "" && AriesHost == "" {
		AriesHost = utils.Config.Aries
	}
	connectionUrl = utils.Config.Aries + "/connections"
	secureApi := r.Group("api/connections")
	api := r.Group("api/connections")
	secureApi.Use(middlewares.JwtMiddleware())
	{
		secureApi.GET("", getConnections)
	}
	{
		api.POST("/create-invitation/:uuid", createInvitation)
		api.GET(":id", getConnectionById)
		api.DELETE("/:id", deleteConnections)
	}
}

func getConnections(c *gin.Context) {
	resp, err := http.Get(connectionUrl)
	fmt.Println(connectionUrl)
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

func getConnectionById(c *gin.Context) {
	id := c.Param("id")
	connection, err := usecases.GetConnectionDetails(id)
	if err != nil {
		apiError := utils.NewApiError("DESIRIALIAZATION_ERROR", err, "details")
		c.JSON(400, apiError)
		return
	}
	c.JSON(200, connection)
}

func createInvitation(c *gin.Context) {
	employeeInvitationId := c.Param("uuid")
	if employeeInvitationId == "" {
		apiError := utils.NewApiError("INVITATION_NOT_FOUND", nil, "missing invitation")
		c.JSON(500, apiError)
		return
	}
	fmt.Printf("employeeInvitation: %s \n", employeeInvitationId)
	resp, err := http.Post(fmt.Sprintf("%s/create-invitation", connectionUrl), "application/json", nil)
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
		apiErr := utils.NewApiError(
			"NOT_FOUND",
			err,
			fmt.Sprintf("Could't find employee with invitaion: %s", employeeInvitationId),
		)
		apiErr.Enhance(c)
		c.JSON(400, apiErr)
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
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s", connectionUrl, c.Param("id")), nil)
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
