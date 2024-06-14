package controller

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func RegisterProRouter(r *mux.Router, db *sql.DB) {
	project_router := r.PathPrefix("/projects").Subrouter()
	project_router.HandleFunc("", handleGetProject(db)).Methods("GET")
	project_router.HandleFunc("/{project-id}/pm-in-project", handleGetPM(db)).Methods("GET")
	project_router.HandleFunc("/{project-id}/member-in-project", handleGetMember(db)).Methods("GET")
	project_router.HandleFunc("/{project-id}/pm-outside-project", handleGetPMNotInPro(db)).Methods("GET")
	project_router.HandleFunc("/{project-id}/member-outside-project", handleGetMemNotInPro(db)).Methods("GET")
	project_router.HandleFunc("", handleCreateProject(db)).Methods("POST")
	project_router.HandleFunc("/{project-id}", handleUpdateProject(db)).Methods("PUT")
	project_router.HandleFunc("/{project-id}/project-managers", handleMapProjectManager(db)).Methods("POST")
	project_router.HandleFunc("/{project-id}/member", handleMapProjectMember(db)).Methods("POST")
	project_router.HandleFunc("/{project-id}/{member-id}", handleRemoveMemberInProject(db)).Methods("DELETE")
	// pm : get project that this pm with id ... in
	// ...
}
