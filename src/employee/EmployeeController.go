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
		api.GET("/:id", amserr.ErrorWrapper(GetById))
	}
}

func GetById(c *gin.Context) *amserr.ApiError {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err.(*amserr.ApiError)
	}
	result, err := FindEmployeeById(id)
	if err != nil {
		return err.(*amserr.ApiError)
	}
	c.JSON(200, &result)
	return nil
}

func GetAll(c *gin.Context) *amserr.ApiError {
	return nil
}
