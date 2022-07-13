package controllers

import (
	"ams-back/dtos"
	"ams-back/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

var issueCredentialUrl = ""

func CreateIssueCredentialController(r *gin.Engine) {
	if utils.Config.Aries != "" && AriesHost == "" {
		AriesHost = utils.Config.Aries
	}
	issueCredentialUrl = AriesHost + "/issue-credential-2.0"
	api := r.Group("api/issue-credentials")
	{
		api.POST("", handleIssueCredential)
		api.GET("", getCredentialsRecords)
		api.GET("/:id", getCredentialsRecordsByConnectionId)
	}
}

func handleIssueCredential(c *gin.Context) {
	dto := dtos.IssueCredentialDTO{}
	serializationError := c.Bind(&dto)
	if serializationError != nil {
		apiError := utils.NewApiError(
			"INVALID_INPUT",
			serializationError,
			"Error serializing Struc of type IssueCredentialsDto",
		)
		apiError.Enhance(c)
		c.JSON(400, &apiError)
		return
	}
	// todo make request to agent to issue credentials
}

func getCredentialsRecords(c *gin.Context) {
	resp, err := http.Get(issueCredentialUrl + "/records")
	if err != nil {
		apiError := utils.NewApiError("REQUEST_FAILED", err, "details")
		apiError.Enhance(c)
		c.JSON(400, apiError)
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

func getCredentialsRecordsByConnectionId(c *gin.Context) {
	id := c.Param("id")
	resp, err := http.Get(issueCredentialUrl + "/records")
	if err != nil {
		apiError := utils.NewApiError("REQUEST_FAILED", err, "details")
		apiError.Enhance(c)
		c.JSON(400, apiError)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		apiError := utils.NewApiError("DESIRIALIAZATION_ERROR", err, "details")
		c.JSON(400, apiError)
		return
	}
	jsonString := string(body)
	dto := dtos.CredentialExchangeRecordDTO{}
	json.Unmarshal([]byte(jsonString), &dto)
	credentialExchanges := dto.Results
	var credentialExchangesForId []dtos.CredentialExchangeRecord
	for i := range credentialExchanges {
		if credentialExchanges[i].CredExRecord.ConnectionId != id {
			continue
		}
		credentialExchangesForId = append(credentialExchangesForId, credentialExchanges[i])
	}
	result := dtos.CredentialExchangeRecordDTO{
		credentialExchangesForId,
	}
	c.JSON(200, &result)
}
