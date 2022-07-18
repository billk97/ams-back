package controllers

import (
	"ams-back/middlewares"
	models "ams-back/models"
	repos "ams-back/repos"
	utils "ams-back/utils"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateResourceController(r *gin.Engine) {
	secureApi := r.Group("api/resource")
	secureApi.Use(middlewares.JwtMiddleware())
	{
		secureApi.GET("/:id", getResourceById)
		secureApi.GET("", getResources)
		secureApi.POST("", createResource)
		secureApi.PUT("/:id", updateResource)
	}
}

func createResource(c *gin.Context) {
	resource := models.Resource{}
	serializationError := c.Bind(&resource)
	if serializationError != nil {
		apiErr := utils.NewApiError(
			"INVALID_INPUT",
			serializationError,
			fmt.Sprintf("Erro serializing json to resource struct"),
		)
		apiErr.Enhance(c)
		c.JSON(400, apiErr)
		return
	}
	databaseError := repos.SaveResource(&resource)
	if databaseError != nil {
		apiErr := utils.NewApiError(
			"PERSIST_ENTITY_FAILED",
			databaseError,
			fmt.Sprintf("Could't persist entity of type Resource"),
		)

		apiErr.Enhance(c)
		c.JSON(400, apiErr)
		return
	}
	c.JSON(200, &resource)
}

func updateResource(c *gin.Context) {
	id, serializationError := strconv.Atoi(c.Param("id"))
	if serializationError != nil {
		apiErr := utils.NewApiError(
			"INVALID_INPUT",
			serializationError,
			fmt.Sprintf("Erro serializing id"),
		)
		apiErr.Enhance(c)
		c.JSON(400, apiErr)
		return
	}
	resource := models.Resource{}
	serializationError = c.BindJSON(&resource)
	if serializationError != nil {
		apiErr := utils.NewApiError(
			"INVALID_INPUT",
			serializationError,
			fmt.Sprintf("Erro serializing json entity resource"),
		)
		apiErr.Enhance(c)
		c.JSON(400, apiErr)
		return
	}
	resource.ID = uint(id)
	databaseError := repos.UpdateResource(&resource)
	if databaseError != nil {
		apiError := utils.NewApiError(
			"PERSIST_ENTITY_FAILED",
			databaseError,
			fmt.Sprintf("Could't persist entity of type Resource"),
		)
		apiError.Enhance(c)
		c.JSON(400, apiError)
		return
	}
	c.JSON(200, &resource)
}

func getResourceById(c *gin.Context) {
	id, serializationError := strconv.Atoi(c.Param("id"))
	if serializationError != nil {
		apiErr := utils.NewApiError(
			"INVALID_INPUT",
			serializationError,
			fmt.Sprintf("Erro serializing id"),
		)
		apiErr.Enhance(c)
		c.JSON(400, apiErr)
		return
	}
	resource, databaseError := repos.FindOneResourceById(id)
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
	if resource == nil {
		apiErr := utils.NewApiError(
			"NOT_FOUND",
			nil,
			fmt.Sprintf("Resource with ID %d not found", id),
		)
		apiErr.Enhance(c)
		c.JSON(400, apiErr)
		return
	}
	c.JSON(200, resource)
}

func getResources(c *gin.Context) {
	query := c.Query("page")
	if query == "" {
		query = "0"
	}
	page, serializationError := strconv.Atoi(query)
	if serializationError != nil {
		apiErr := utils.NewApiError(
			"INVALID_INPUT",
			serializationError,
			fmt.Sprintf("Erro serializing page"),
		)
		apiErr.Enhance(c)
		c.JSON(400, apiErr)
		return
	}
	resources, databaseError := repos.FindResource(page)
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
	c.JSON(200, resources)
}
