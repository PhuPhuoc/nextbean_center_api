package controller

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func RegisterAccountRouter(r *mux.Router, db *sql.DB) {
	r.HandleFunc("/login", HandleLogin(db)).Methods("POST")
}
