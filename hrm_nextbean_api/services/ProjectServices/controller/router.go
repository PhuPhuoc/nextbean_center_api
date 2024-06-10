package controller

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func RegisterProjectRouter(r *mux.Router, db *sql.DB) {
	project_router := r.PathPrefix("/project").Subrouter()
	project_router.HandleFunc("", handleCreateProject(db)).Methods("POST")
}
