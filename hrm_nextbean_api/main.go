package main

import (
	"log"

	"github.com/PhuPhuoc/hrm_nextbean_api/api"
	"github.com/PhuPhuoc/hrm_nextbean_api/database"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
)

// @title						Intern Resource Management System
// @version					1.0
// @description				A web application to manage interns at the Nextbean Center, designed to oversee their daily tasks and schedules while working at the office. The app aims to streamline the management of internship programs, enabling efficient tracking of intern activities, assignments, and attendance for a well-organized and productive internship experience.
// @host						localhost:8080
// @BasePath					/api/v1
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	conn_str := utils.GetConnStr(false)
	port := utils.GetPort()
	db, err_db := database.InitMySQLStore(conn_str)
	if err_db != nil {
		log.Fatal("|main| ~ Cannot connect to database: ", err_db)
	}
	defer db.Close()

	server := api.InitServer(port, db)
	if err_run_server := server.RunApp(); err_run_server != nil {
		log.Fatal("|main| ~ Cannot run app: ", err_run_server)
	}
}
