package controller

import (
	"database/sql"

	"github.com/PhuPhuoc/hrm_nextbean_api/middleware"
	"github.com/gorilla/mux"
)

func RegisterTaskRouter(r *mux.Router, db *sql.DB) {
	task_router := r.PathPrefix("/tasks").Subrouter()
	task_router.Use(middleware.AuthMiddleware(db))
	task_router.HandleFunc("/{project-id}", middleware.TaskAccessMiddleware(db, false, true, true)(handleCreateTask(db))).Methods("POST")
	// timetable_router.HandleFunc("", handleGetTimetable(db)).Methods("GET")
	// timetable_router.HandleFunc("/{timetable-id}/approve", handleApproveInternTimeTable(db)).Methods("POST")
	// for intern: checkin checkout - update
}
