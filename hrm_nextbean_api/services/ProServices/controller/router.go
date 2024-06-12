package controller

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func RegisterProRouter(r *mux.Router, db *sql.DB) {
	project_router := r.PathPrefix("/project").Subrouter()
	project_router.HandleFunc("/get", handleGetProject(db)).Methods("POST")
	project_router.HandleFunc("", handleCreateProject(db)).Methods("POST")
	project_router.HandleFunc("", handleUpdateProject(db)).Methods("PUT")
	project_router.HandleFunc("/manager", handleMapProjectManager(db)).Methods("POST")
	project_router.HandleFunc("/member", handleMapProjectMember(db)).Methods("POST")
}
