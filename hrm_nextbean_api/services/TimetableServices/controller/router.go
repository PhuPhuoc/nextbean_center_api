package controller

import (
	"database/sql"

	"github.com/PhuPhuoc/hrm_nextbean_api/middleware"
	"github.com/gorilla/mux"
)

func RegisterTimetableRouter(r *mux.Router, db *sql.DB) {
	timetable_router := r.PathPrefix("/timetables").Subrouter()
	timetable_router.Use(middleware.AuthMiddleware(db))
	timetable_router.HandleFunc("", middleware.TimetableAccessMiddleware(db, true, false)(handleGetTimetable(db))).Methods("GET")
	timetable_router.HandleFunc("/weekly", middleware.TimetableAccessMiddleware(db, true, false)(handleGetWeeklyTimetable(db))).Methods("GET")
	timetable_router.HandleFunc("", middleware.TimetableAccessMiddleware(db, false, true)(handleCreateTimeTable(db))).Methods("POST")
	timetable_router.HandleFunc("/{timetable-id}/approve", middleware.TimetableAccessMiddleware(db, true, false)(handleApproveInternTimeTable(db))).Methods("POST")
	// for intern: checkin checkout - update
}
