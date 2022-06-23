package main

import (
	controllers "ams-back/controllers"
	database "ams-back/database"
	usecases "ams-back/usecases"
	utils "ams-back/utils"

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
	controllers.CreateResourceConntroller(r)
	controllers.CreatePermissionController(r)
	r.Run(":5000")
}
