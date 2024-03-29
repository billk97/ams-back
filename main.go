package main

import (
	"ams-back/controllers"
	"ams-back/database"
	"ams-back/middlewares"
	"ams-back/usecases"
	"ams-back/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	// TODO get the public did?
	var config utils.Env
	utils.InitEnv()
	config = utils.Config

	database.InitDbConnection(config.DB)
	database.Synchronize(database.GetDb())
	usecases.CreateSuperAdminIfNotExists()
	r := gin.Default()
	r.Use(middlewares.JSONMiddleware())
	r.Use(middlewares.CORSMiddleware())
	controllers.CreateUrlController(r)
	controllers.CreateAdminController(r)
	controllers.CreateResourceController(r)
	controllers.CreatePermissionController(r)
	controllers.CreateConnectionsController(r)
	controllers.CreateWebhookController(r)
	controllers.CreateIssueCredentialsWebhookController(r)
	controllers.CreateContextController(r)
	controllers.CreateIssueCredentialController(r)
	controllers.CreateWalletController(r)
	controllers.CreateVerifierTokenController(r)

	r.Run(":5000")
}
