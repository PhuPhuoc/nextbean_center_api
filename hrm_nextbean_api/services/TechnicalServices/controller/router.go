package controller

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func RegisterTechnicalRouter(r *mux.Router, db *sql.DB) {
	tech_router := r.PathPrefix("/technicals").Subrouter()
	tech_router.HandleFunc("", handleGetTech(db)).Methods("GET")
	tech_router.HandleFunc("", handleCreateTech(db)).Methods("POST")
	// tech_router.HandleFunc("", handleUpdateOJT(db)).Methods("PUT")
	// tech_router.HandleFunc("/{id}", handleDeleteOJT(db)).Methods("DELETE")
}
