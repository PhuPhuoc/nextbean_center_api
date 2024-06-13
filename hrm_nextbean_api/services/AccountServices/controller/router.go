package controller

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func RegisterAccountRouter(r *mux.Router, db *sql.DB) {
	r.HandleFunc("/login", HandleLogin(db)).Methods("POST")

	account_router := r.PathPrefix("/accounts").Subrouter()
	account_router.HandleFunc("", handleCreateAccount(db)).Methods("POST")
	account_router.HandleFunc("", handleGetAccount(db)).Methods("GET")
	account_router.HandleFunc("/{account-id}", handleUpdateAccount(db)).Methods("PUT")
	account_router.HandleFunc("/{account-id}", handleDeleteAccount(db)).Methods("DELETE")
}
