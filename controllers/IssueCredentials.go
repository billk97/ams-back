package controllers

import (
	"ams-back/dtos"
	"ams-back/middlewares"
	"ams-back/usecases"
	"ams-back/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

var issueCredentialUrl = ""

func CreateIssueCredentialController(r *gin.Engine) {
	if utils.Config.Aries != "" && AriesHost == "" {
		AriesHost = utils.Config.Aries
	}
	issueCredentialUrl = AriesHost + "/issue-credential-2.0"
	secureApi := r.Group("api/issue-credentials")
	secureApi.Use(middlewares.JwtMiddleware())
	{
		secureApi.POST(":id", handleIssueCredential)
		secureApi.GET("", getCredentialsRecords)
		secureApi.GET(":id", getCredentialsRecordsByConnectionId)
		secureApi.DELETE(":state", deleteAllCredentialOfferByState)
	}
}

func handleIssueCredential(c *gin.Context) {
	id := c.Param("id")
	employeeId, err := strconv.Atoi(id)
	if err != nil {
		apiError := utils.NewApiError("REQUEST_FAILED", err, "details")
		apiError.Enhance(c)
		c.JSON(400, apiError)
		return
	}
	responseDTO, err := usecases.CreateAndSendIssueCredentialRequest(employeeId)
	if err != nil {
		apiError := utils.NewApiError("REQUEST_FAILED", err, "details")
		apiError.Enhance(c)
		c.JSON(400, apiError)
		return
	}
	c.JSON(200, responseDTO)
}

func deleteAllCredentialOfferByState(c *gin.Context) {
	state := c.Param("state")
	result, err := usecases.DeleteAllCredentialsOffersWithState(state)
	if err != nil {
		apiError := utils.NewApiError("REQUEST_FAILED", err, "details")
		apiError.Enhance(c)
		c.JSON(400, apiError)
		return
	}
	dto := dtos.Wrapper{
		Result: result,
	}
	c.JSON(200, &dto)
}

func getCredentialsRecords(c *gin.Context) {
	dto, err := usecases.GetAllCredentialOffers()
	if err != nil {
		apiError := utils.NewApiError("REQUEST_FAILED", err, "details")
		apiError.Enhance(c)
		c.JSON(400, apiError)
		return
	}
	c.JSON(200, dto)
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
