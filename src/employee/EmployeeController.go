package employee

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func GetById(c *gin.Context) *ApiError {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err.(*ApiError)
	}
	result, err := FindEmployeeById(id)
	if err != nil {
		return err.(*ApiError)
	}
	c.JSON(200, &result)
	return nil
}

type appHandler func(*gin.Context) *ApiError

func (a appHandler) HandlerFunc(c *gin.Context) {
	fmt.Println("test")
	err := a(c)
	fmt.Print("skata-> ")
	fmt.Println(err.Err)
	if err != nil {
		err.Enhance(c)
		c.AbortWithStatusJSON(400, err)
	}
}

func wrapper(a appHandler) gin.HandlerFunc {
	return a.HandlerFunc
}

func CreateUrlConntroller() {
	Router = gin.Default()
	api := Router.Group("api/employees")
	{
		api.GET("/:id", wrapper(GetById))
	}
}
