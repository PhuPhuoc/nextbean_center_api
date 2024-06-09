package controller

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func RegisterInterntRouter(r *mux.Router, db *sql.DB) {
	intern_router := r.PathPrefix("/intern").Subrouter()
	intern_router.HandleFunc("", HandleCreateIntern(db)).Methods("POST")
	intern_router.HandleFunc("/get", HandleGetIntern(db)).Methods("POST")
	intern_router.HandleFunc("", HandleUpdateIntern(db)).Methods("PUT")
}
