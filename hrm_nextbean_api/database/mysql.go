package database

import (
	"database/sql"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func InitMySQLStore(conn_str string) (*sql.DB, error) {
	db, err_connect := sql.Open("mysql", conn_str)
	if err_connect != nil {
		log.Println("|database| ~ cannot connect to db: ", err_connect)
		return nil, err_connect
	}

	if err_ping := db.Ping(); err_ping != nil {
		db.Close()
		if strings.Contains(err_ping.Error(), "connection refused") {
			log.Println("|database| ~ Connection refused ~ Wait for the db'container to start ...")
		} else {
			log.Fatalf("|database| ~ Error: %v", err_ping.Error())
		}
	} else {
		log.Println("|database| ~ Ping to database successful ...")
	}
	return db, nil
}
