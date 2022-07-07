package controllers

import (
	"ams-back/dtos"
	"ams-back/models"
	"ams-back/repos"
	"ams-back/utils"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateUrlController(r *gin.Engine) {
	api := r.Group("api/employees")
	{
		api.GET("/", getAll)
		api.GET("/:id", getById)
		api.POST("/", createEmployee)
		api.PUT("/:id", updateEmployeeData)
	}
}

func getById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, err)
		return
	}
	result, err := repos.FindEmployeeById(id)
	if result == nil && err != nil {
		apiErr := utils.NewApiError(
			"NOT_FOUND",
			err,
			fmt.Sprintf("Could't find employee withId: %d", id),
		)

		apiErr.Enhance(c)
		c.JSON(400, apiErr)
		return
	}
	c.JSON(200, result)
}

func getAll(c *gin.Context) {
	em, err := repos.FindEmployees(2)

	if em == nil && err != nil {
		apiErr := utils.NewApiError(
			"NOT_FOUND",
			err,
			fmt.Sprintf("Could't find employees: "),
		)
		apiErr.Enhance(c)
		c.JSON(400, apiErr)
		return
	}
	c.JSON(200, dtos.Wrapper{
		Result: em,
	})
}

func createEmployee(c *gin.Context) {
	emp := models.Employee{}
	err := c.BindJSON(&emp)
	if err != nil {
		apiErr := utils.NewApiError(
			"INVALID_INPUT",
			err,
			fmt.Sprintf("Erro serializing json to employee struct"),
		)
		apiErr.Enhance(c)
		c.JSON(400, apiErr)
		return
	}
	e, err := repos.SaveEmploy(&emp)
	if e == nil && err != nil {
		apiError := utils.NewApiError(
			"PERSIST_ENTITY_FAILED",
			err,
			fmt.Sprintf("Could't persist entity of type employee"),
		)
		apiError.Enhance(c)
		c.JSON(400, apiError)
		return
	}
	c.JSON(200, emp)
}

func updateEmployeeData(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, err)
		return
	}
	emp := models.Employee{}
	err = c.BindJSON(&emp)
	if err != nil {
		apiErr := utils.NewApiError(
			"INVALID_INPUT",
			err,
			fmt.Sprintf("Erro serializing json to employee struct"),
		)
		apiErr.Enhance(c)
		c.JSON(400, apiErr)
		return
	}
	emp.ID = uint(id)
	e, err := repos.UpdateEmployee(&emp)
	if e == nil && err != nil {
		apiErr := utils.NewApiError(
			"PERSIST_ENTITY_FAILED",
			err,
			fmt.Sprintf("Could't persist entity of type employee"),
		)
		apiErr.Enhance(c)
		c.JSON(400, apiErr)
		return
	}
	fmt.Println(e)
	c.JSON(200, emp)
}
