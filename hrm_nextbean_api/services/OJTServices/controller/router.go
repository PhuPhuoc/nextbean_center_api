package controller

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func RegisterOJTRouter(r *mux.Router, db *sql.DB) {
	ojt_router := r.PathPrefix("/ojt").Subrouter()
	ojt_router.HandleFunc("/get", handleGetOJT(db)).Methods("POST")
	ojt_router.HandleFunc("", handleCreateOJT(db)).Methods("POST")
	ojt_router.HandleFunc("", handleUpdateOJT(db)).Methods("PUT")
	ojt_router.HandleFunc("/{id}", handleDeleteOJT(db)).Methods("DELETE")
}
