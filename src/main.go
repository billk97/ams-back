package main

import (
	"ams-back/src/database"
	"ams-back/src/employee"
	"ams-back/src/utils"
	"time"
)

func main() {
	// get config from yaml
	var config utils.Env

	utils.InitYamlConfig()
	config = utils.Config
	database.InitDbConnection(config.DB)
	e := employee.Employee{
		FirstName:       "firstName",
		LastName:        "lastName",
		JobTitle:        "jobTitle",
		DirthDate:       time.Now(),
		Email:           "email",
		MobileNumber:    "mobileNumber",
		DidConnectionId: "didConnectionId",
	}
	employee.CreateEmploy(&e)
	// init database pull
	// create
}
