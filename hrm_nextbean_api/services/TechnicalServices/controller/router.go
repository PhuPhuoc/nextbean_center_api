package controller

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func RegisterTechnicalRouter(r *mux.Router, db *sql.DB) {
	tech_router := r.PathPrefix("/technical").Subrouter()
	tech_router.HandleFunc("/get", handleGetTech(db)).Methods("POST")
	tech_router.HandleFunc("", handleCreateTech(db)).Methods("POST")
	// tech_router.HandleFunc("", handleUpdateOJT(db)).Methods("PUT")
	// tech_router.HandleFunc("/{id}", handleDeleteOJT(db)).Methods("DELETE")
}
