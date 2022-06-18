package employee

import (
	"ams-back/src/amserr"
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
	}
	c.JSON(200, em)
}

func CreateEmployee(c *gin.Context) {
	emp := Employee{}
	err := c.BindJSON(&emp)
	if err != nil {
		apiErr := amserr.NewApiError(
			"INVALID_INPUT",
			err,
			fmt.Sprintf("Erro serializing json to employee struct"),
		)
		c.JSON(400, apiErr)
	}
	// c.Request.Body
}
