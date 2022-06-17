package employee

import (
	"ams-back/src/amserr"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func CreateUrlConntroller(r *gin.Engine) {
	Router = r
	api := Router.Group("api/employees")
	{
		api.GET("/", amserr.ErrorWrapper(GetAll))
		api.GET("/:id", (GetById))
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

func GetAll(c *gin.Context) *amserr.ApiError {
	return nil
}
