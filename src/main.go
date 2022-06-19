package main

import (
	"ams-back/src/controllers"
	"ams-back/src/database"
	"ams-back/src/usecases"
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
	database.Synchronize(database.GetDb())
	usecases.CreateSuperAdminIfNotExists()
	r := gin.Default()
	r.Use(JSONMiddleware())
	controllers.CreateUrlConntroller(r)
	controllers.CreateAdminConntroller(r)
	controllers.Router.Run(":5000")
}
