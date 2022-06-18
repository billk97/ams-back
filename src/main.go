package main

import (
	"ams-back/src/database"
	"ams-back/src/employee"
	"ams-back/src/utils"

	"github.com/gin-gonic/gin"
)

func JSONMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
	}
}

func main() {
	var config utils.Env
	utils.InitYamlConfig()
	config = utils.Config
	database.InitDbConnection(config.DB)
	r := gin.Default()
	r.Use(JSONMiddleware())
	employee.CreateUrlConntroller(r)
	employee.Router.Run(":5000")
}
