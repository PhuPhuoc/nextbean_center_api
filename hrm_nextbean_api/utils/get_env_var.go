package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadenv() {
	if err_env := godotenv.Load(".env"); err_env != nil {
		log.Fatal("|util| ~ Error loading .env file: ", err_env)
	}
}

func GetPort() string {
	loadenv()
	port := os.Getenv("PORT")
	return port
}

func GetConnStr() string {
	flag := false
	var conn_str string
	if !flag {
		loadenv()
		conn_str = os.Getenv("DB_CONN_STR_DOCKER")
	} else {
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbUser := os.Getenv("DB_USER")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbName := os.Getenv("DB_NAME")
		conn_str = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	}
	log.Println("|util| + connection string: ", conn_str)
	return conn_str
}
