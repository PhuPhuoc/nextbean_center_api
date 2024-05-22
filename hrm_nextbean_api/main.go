package main

import (
	"fmt"
	"log"
	"os"

	"github.com/PhuPhuoc/hrm_nextbean_api/api"
	"github.com/PhuPhuoc/hrm_nextbean_api/database"
	"github.com/joho/godotenv"
)

func main() {
	if err_env := godotenv.Load(".env"); err_env != nil {
		log.Fatal("#HRM-nextbean-api#|main| ~ Error loading .env file: ", err_env)
	}

	port := os.Getenv("PORT")
	// conn_str_local := os.Getenv("DB_CONN_STR_DOCKER")

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	conn_str_docker := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err_db := database.InitMySQLStore(conn_str_docker)
	if err_db != nil {
		//log.Println("@ conn-str: ", conn_str_docker)
		log.Fatal("#HRM-nextbean-api#|main| ~ Cannot connect to database: ", err_db)
	}

	server := api.InitServer(port, db)
	if err_run_server := server.RunApp(); err_run_server != nil {
		log.Fatal("#HRM-nextbean-api#|main| ~ Cannot run app: ", err_run_server)
	}
}
