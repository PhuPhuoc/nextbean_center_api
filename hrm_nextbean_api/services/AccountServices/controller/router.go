package controller

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func RegisterAccountRouter(r *mux.Router, db *sql.DB) {
	r.HandleFunc("/login", HandleLogin(db)).Methods("POST")

	account_router := r.PathPrefix("/account").Subrouter()
	account_router.HandleFunc("", HandleCreateAccount(db)).Methods("POST")
}
