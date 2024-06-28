package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadenv() error {
	if err_env := godotenv.Load(".env"); err_env != nil {
		return fmt.Errorf("|util| ~ error loading .env file: %v", err_env)
	}
	return nil
}

func GetPort() string {
	port := os.Getenv("APP_PORT")
	return ":" + port
}

func GetSecretKey() string {
	secret := os.Getenv("SECRET_KEY")
	return secret
}

func GetConnStr(flag bool) string {
	var conn_str string
	if !flag {
		loadenv()
		conn_str = os.Getenv("DB_CONN_STR")
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

func GetClientID() string {
	secret := os.Getenv("CLIENT_ID")
	return secret
}

func GetClientSecret() string {
	secret := os.Getenv("CLIENT_SECRET")
	return secret
}

func GetURLCallback() string {
	secret := os.Getenv("URL_CALLBACK_OAUTH")
	return secret
}
