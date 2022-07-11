package controllers

import (
	models "ams-back/models"
	repos "ams-back/repos"
	utils "ams-back/utils"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePermissionController(r *gin.Engine) {
	api := r.Group("api/permissions")
	{
		api.GET("", getPermissions)
		api.POST("", addPermission)
	}
}

func addPermission(c *gin.Context) {
	permission := models.Permission{}
	serializationError := c.Bind(&permission)
	if serializationError != nil {
		apiErr := utils.NewApiError(
			"INVALID_INPUT",
			serializationError,
			fmt.Sprintf("Erro serializing json to permission struct"),
		)
		apiErr.Enhance(c)
		c.JSON(400, apiErr)
		return
	}
	databaseError := repos.CreatePermission(&permission)
	if databaseError != nil {
		apiError := utils.NewApiError(
			"PERSIST_ENTITY_FAILED",
			databaseError,
			fmt.Sprintf("Could't persist entity of type permision"),
		)
		apiError.Enhance(c)
		c.JSON(400, apiError)
		return
	}
	c.JSON(200, &permission)
}

func getPermissions(c *gin.Context) {
	queryParam := c.Query("page")
	if queryParam == "" {
		queryParam = "0"
	}
	page, serializationError := strconv.Atoi(queryParam)
	if serializationError != nil {
		e := utils.NewApiError(
			"INVALID_INPUT",
			serializationError,
			fmt.Sprintf("Erro serializing page"),
		)
		e.Enhance(c)
		c.JSON(400, e)
		return
	}
	permissions, databaseError := repos.FindPermissions(page)
	if databaseError != nil {
		apiError := utils.NewApiError(
			"QUERY_EXECUTION_FAILED",
			databaseError,
			fmt.Sprintf("Could not execute query"),
		)
		apiError.Enhance(c)
		c.JSON(400, apiError)
		return
	}
	c.JSON(200, &permissions)
}
