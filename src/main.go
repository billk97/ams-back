package main

import (
	"ams-back/src/database"
	"ams-back/src/employee"
	"ams-back/src/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	var config utils.Env
	utils.InitYamlConfig()
	config = utils.Config
	database.InitDbConnection(config.DB)
	r := gin.Default()
	employee.CreateUrlConntroller(r)
	employee.Router.Run(":5000")
}
