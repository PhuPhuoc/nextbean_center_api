package controller

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func RegisterTimetableRouter(r *mux.Router, db *sql.DB) {
	timetable_router := r.PathPrefix("/timetables").Subrouter()
	timetable_router.HandleFunc("", handleGetTimetable(db)).Methods("GET")
	timetable_router.HandleFunc("/{intern-id}", handleCreateTimeTable(db)).Methods("POST")
	timetable_router.HandleFunc("/{timetable-id}/approve", handleApproveInternTimeTable(db)).Methods("POST")
	// for intern: checkin checkout - update
}
