package controllers

import (
	models "ams-back/models"
	repos "ams-back/repos"
	utils "ams-back/utils"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateResourceConntroller(r *gin.Engine) {
	api := r.Group("api/resource")
	{
		api.GET("/:id", getResourceById)
		api.GET("/", getResources)
		api.POST("/", createResource)
		api.PUT("/:id", updateResource)
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
		databaseError.Enhance(c)
		c.JSON(400, databaseError)
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
		databaseError.Enhance(c)
		c.JSON(400, databaseError)
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
		databaseError.Enhance(c)
		c.JSON(400, databaseError)
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
		databaseError.Enhance(c)
		c.JSON(400, databaseError)
		return
	}
	c.JSON(200, resources)
}
