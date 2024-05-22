package main

import (
	"log"

	"github.com/PhuPhuoc/hrm_nextbean_api/api"
	"github.com/PhuPhuoc/hrm_nextbean_api/database"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
)

func main() {
	conn_str := utils.GetConnStr()
	port := utils.GetPort()
	db, err_db := database.InitMySQLStore(conn_str)
	if err_db != nil {
		log.Fatal("|main| ~ Cannot connect to database: ", err_db)
	}

	server := api.InitServer(port, db)
	if err_run_server := server.RunApp(); err_run_server != nil {
		log.Fatal("|main| ~ Cannot run app: ", err_run_server)
	}
}
