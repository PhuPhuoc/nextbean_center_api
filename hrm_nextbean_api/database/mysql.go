package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitMySQLStore(conn_str string) (*sql.DB, error) {
	db, err_connect := sql.Open("mysql", conn_str)
	if err_connect != nil {
		log.Println("err when open sql")
		return nil, err_connect
	}

	defer db.Close()

	// if err_ping := db.Ping(); err_ping != nil {
	// 	log.Println("err when ping sql")
	// 	return nil, err_ping

	// }
	return db, nil
}
