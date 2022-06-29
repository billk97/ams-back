package main

import (
	controllers "ams-back/controllers"
	database "ams-back/database"
	usecases "ams-back/usecases"
	utils "ams-back/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func JSONMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		// c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func main() {
	var config utils.Env
	utils.InitEnv()
	config = utils.Config
	fmt.Println(config.DB.Name)
	fmt.Println(config.DB.Username)
	fmt.Println(config.DB.Password)
	fmt.Println(config.DB.Host)
	fmt.Println(config.DB.Port)
	database.InitDbConnection(config.DB)
	database.Synchronize(database.GetDb())
	usecases.CreateSuperAdminIfNotExists()
	r := gin.Default()
	controllers.CreateUrlConntroller(r)
	controllers.CreateAdminConntroller(r)
	controllers.CreateResourceConntroller(r)
	controllers.CreatePermissionController(r)
	controllers.CreateConnectionsController(r)
	r.Use(JSONMiddleware())
	r.Use(CORSMiddleware())
	r.Run(":5000")
}
