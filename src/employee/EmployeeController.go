package employee

import (
	"ams-back/src/amserr"
	"ams-back/src/models"
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
	result, err := FindEmployeeById(id)
	if result == nil && err != nil {
		c.JSON(400, err)
		return
	}
	c.JSON(200, result)
}

func GetAll(c *gin.Context) {
	em, err := FindEmployees(2)

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
		apiErr := amserr.NewApiError(
			"INVALID_INPUT",
			err,
			fmt.Sprintf("Erro serializing json to employee struct"),
		)
		apiErr.Enhance(c)
		c.JSON(400, apiErr)
		return
	}
	e, err := SaveEmploy(&emp)
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
		apiErr := amserr.NewApiError(
			"INVALID_INPUT",
			err,
			fmt.Sprintf("Erro serializing json to employee struct"),
		)
		apiErr.Enhance(c)
		c.JSON(400, apiErr)
		return
	}
	emp.ID = uint(id)
	e, err := UpdateEmployee(&emp)
	if e == nil && err != nil {
		c.JSON(400, err)
		return
	}
	fmt.Println(e)
	c.JSON(200, emp)
}
