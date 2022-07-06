package controllers

import (
	"ams-back/dtos"
	"github.com/gin-gonic/gin"
)

func CreateContextController(r *gin.Engine) {
	api := r.Group("contexts/rooms")
	{
		api.GET("/v1", getContext)
		api.GET("/", getContext)
	}
}

func getContext(c *gin.Context) {
	c.JSON(200, dtos.NewContext())
}
