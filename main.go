package main

import (
	controllers "ams-back/controllers"
	database "ams-back/database"
	"ams-back/middlewares"
	usecases "ams-back/usecases"
	utils "ams-back/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

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
	controllers.CreateUrlController(r)
	controllers.CreateAdminController(r)
	controllers.CreateResourceController(r)
	controllers.CreatePermissionController(r)
	controllers.CreateConnectionsController(r)
	controllers.CreateWebhookController(r)
	controllers.CreateIssueCredentialsWebhookController(r)
	controllers.CreateContextController(r)
	controllers.CreateIssueCredentialController(r)
	r.Use(middlewares.JSONMiddleware())
	r.Use(middlewares.CORSMiddleware())
	r.Run(":5000")
}
