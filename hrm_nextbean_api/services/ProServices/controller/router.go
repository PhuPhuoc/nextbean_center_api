package controller

import (
	"database/sql"

	"github.com/PhuPhuoc/hrm_nextbean_api/middleware"
	"github.com/gorilla/mux"
)

func RegisterProRouter(r *mux.Router, db *sql.DB) {
	project_router := r.PathPrefix("/projects").Subrouter()

	project_router.Use(middleware.AuthMiddleware(db))

	project_router.HandleFunc("", middleware.ProjectAccessMiddleware(db, true, true, true, true)(handleGetProject(db))).Methods("GET")
	project_router.HandleFunc("/{project-id}/detail", middleware.ProjectAccessMiddleware(db, true, true, false, false)(handleGetProjectDetail(db))).Methods("GET")
	project_router.HandleFunc("/{project-id}/pm-in-project", middleware.ProjectAccessMiddleware(db, true, true, false, false)(handleGetPM(db))).Methods("GET")
	project_router.HandleFunc("/{project-id}/member-in-project", middleware.ProjectAccessMiddleware(db, true, true, false, false)(handleGetMember(db))).Methods("GET")
	project_router.HandleFunc("/{project-id}/pm-outside-project", middleware.ProjectAccessMiddleware(db, false, false, false, false)(handleGetPMNotInPro(db))).Methods("GET")
	project_router.HandleFunc("/{project-id}/member-outside-project", middleware.ProjectAccessMiddleware(db, false, false, false, false)(handleGetMemNotInPro(db))).Methods("GET")
	project_router.HandleFunc("", middleware.ProjectAccessMiddleware(db, false, false, false, false)(handleCreateProject(db))).Methods("POST")
	project_router.HandleFunc("/{project-id}", middleware.ProjectAccessMiddleware(db, false, false, false, false)(handleUpdateProject(db))).Methods("PUT")
	project_router.HandleFunc("/{project-id}/project-managers", middleware.ProjectAccessMiddleware(db, false, false, false, false)(handleMapProjectManager(db))).Methods("POST")
	project_router.HandleFunc("/{project-id}/member", middleware.ProjectAccessMiddleware(db, false, false, false, false)(handleMapProjectMember(db))).Methods("POST")
	project_router.HandleFunc("/{project-id}/{member-id}", middleware.ProjectAccessMiddleware(db, false, false, false, false)(handleRemoveMemberInProject(db))).Methods("DELETE")

}
