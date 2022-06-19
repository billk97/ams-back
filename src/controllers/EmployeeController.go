package controllers

import (
	"ams-back/src/models"
	"ams-back/src/repos"
	"ams-back/src/utils"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func CreateUrlConntroller(r *gin.Engine) {
	Router = r
	api := Router.Group("api/employees")
	{
		api.GET("/", GetAll)
		api.GET("/:id", GetById)
		api.POST("/", CreateEmployee)
		api.PUT("/:id", UpdateEmployeeData)
	}
}

func GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, err)
		return
	}
	result, err := repos.FindEmployeeById(id)
	if result == nil && err != nil {
		c.JSON(400, err)
		return
	}
	c.JSON(200, result)
}

func GetAll(c *gin.Context) {
	em, err := repos.FindEmployees(2)

	if em == nil && err != nil {
		c.JSON(400, err)
		return
	}
	c.JSON(200, em)
}

func CreateEmployee(c *gin.Context) {
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
		c.JSON(400, err)
		return
	}
	fmt.Println(e)
	c.JSON(200, emp)
}

func UpdateEmployeeData(c *gin.Context) {
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
		c.JSON(400, err)
		return
	}
	fmt.Println(e)
	c.JSON(200, emp)
}
