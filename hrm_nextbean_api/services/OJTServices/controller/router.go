package controller

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func RegisterOJTRouter(r *mux.Router, db *sql.DB) {
	ojt_router := r.PathPrefix("/ojts").Subrouter()
	ojt_router.HandleFunc("", handleGetOJT(db)).Methods("GET")
	ojt_router.HandleFunc("", handleCreateOJT(db)).Methods("POST")
	ojt_router.HandleFunc("/{ojt-id}", handleUpdateOJT(db)).Methods("PUT")
	ojt_router.HandleFunc("/{ojt-id}", handleDeleteOJT(db)).Methods("DELETE")
}
