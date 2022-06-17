package main

import (
	"ams-back/src/database"
	"ams-back/src/employee"
	"ams-back/src/utils"
)

func main() {
	var config utils.Env
	utils.InitYamlConfig()
	config = utils.Config
	database.InitDbConnection(config.DB)
	// e := employee.Employee{
	// 	FirstName:       "firstName",
	// 	LastName:        "lastName",
	// 	JobTitle:        "jobTitle",
	// 	DirthDate:       time.Now(),
	// 	Email:           "email",
	// 	MobileNumber:    "mobileNumber",
	// 	DidConnectionId: "didConnectionId",
	// }
	// database.GetDb().AutoMigrate(&employee.Employee{})
	// employee.CreateEmploy(&e)
	// em := employee.GetEmployeeById(5)
	// fmt.Println(em)
	// init database pull
	// create
	employee.CreateUrlConntroller()
	employee.Router.Run(":5000")
}
