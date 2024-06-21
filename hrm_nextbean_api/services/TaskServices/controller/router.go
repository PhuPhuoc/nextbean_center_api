package controller

import (
	"database/sql"

	"github.com/PhuPhuoc/hrm_nextbean_api/middleware"
	"github.com/gorilla/mux"
)

func RegisterTaskRouter(r *mux.Router, db *sql.DB) {
	task_router := r.PathPrefix("/projects/{project-id}/tasks").Subrouter()
	task_router.Use(middleware.AuthMiddleware(db))
	task_router.HandleFunc("", middleware.TaskAccessMiddleware(db, false, true, true)(handleCreateTask(db))).Methods("POST")
	// manager vs pm -> get all task in project
	task_router.HandleFunc("", middleware.TaskAccessMiddleware(db, true, true, false)(handleGetTask(db))).Methods("GET")
	// intern -> get all my task in project
	task_router.HandleFunc("/my-task", middleware.TaskAccessMiddleware(db, false, false, true)(handleGetMyTask(db))).Methods("GET")
	// pm -> update task'details
	// intern -> update task'status vs actual_effort

}
