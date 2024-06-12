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
	project_router.HandleFunc("/remove-member", handleRemoveMemberInProject(db)).Methods("PUT")
	project_router.HandleFunc("/get-pm/{project-id}", handleGetPM(db)).Methods("GET")
	project_router.HandleFunc("/get-mem/{project-id}", handleGetMember(db)).Methods("GET")
	project_router.HandleFunc("/get-pm-not-in/{project-id}", handleGetPMNotInPro(db)).Methods("GET")
	project_router.HandleFunc("/get-mem-not-in/{project-id}", handleGetMemNotInPro(db)).Methods("GET")
}
