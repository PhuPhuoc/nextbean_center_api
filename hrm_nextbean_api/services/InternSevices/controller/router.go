package controller

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func RegisterInterntRouter(r *mux.Router, db *sql.DB) {
	intern_router := r.PathPrefix("/intern").Subrouter()
	intern_router.HandleFunc("", HandleCreateIntern(db)).Methods("POST")
	intern_router.HandleFunc("/get/{page}/{psize}", HandleGetIntern(db)).Methods("POST")
}
