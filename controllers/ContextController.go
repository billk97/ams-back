package controllers

import (
	"ams-back/dtos"
	"github.com/gin-gonic/gin"
)

func CreateContextController(r *gin.Engine) {
	api := r.Group("contexts")
	{
		api.GET("rooms/v1", getContext)
		api.GET("rooms", getContext)
		api.GET("alphacorp-employee/v1", getPersonContext)
		api.GET("alphacorp-employee", getPersonContext)
	}
}

func getContext(c *gin.Context) {
	c.JSON(200, dtos.NewContext())
}

func getPersonContext(c *gin.Context) {
	c.JSON(200, dtos.NewPersonContext())
}
