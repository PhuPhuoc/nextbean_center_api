package controller

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func RegisterInterntRouter(r *mux.Router, db *sql.DB) {
	intern_router := r.PathPrefix("/intern").Subrouter()
	intern_router.HandleFunc("/{account-id}", handleGetDetailIntern(db)).Methods("GET")
	intern_router.HandleFunc("", handleCreateIntern(db)).Methods("POST")
	intern_router.HandleFunc("/get", handleGetIntern(db)).Methods("POST")
	intern_router.HandleFunc("", handleUpdateIntern(db)).Methods("PUT")
	intern_router.HandleFunc("/skill", handleMapInternSkill(db)).Methods("POST")

}
