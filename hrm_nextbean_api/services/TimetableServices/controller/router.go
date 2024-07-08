package controller

import (
	"database/sql"

	"github.com/PhuPhuoc/hrm_nextbean_api/middleware"
	"github.com/gorilla/mux"
)

func RegisterTimetableRouter(r *mux.Router, db *sql.DB) {
	timetable_router := r.PathPrefix("/timetables").Subrouter()
	timetable_router.Use(middleware.AuthMiddleware(db))
	// for admin: get all timetable in db
	timetable_router.HandleFunc("", middleware.TimetableAccessMiddleware(db, true, true, false)(handleGetTimetable(db))).Methods("GET")
	// for admin: get timetable in a week
	timetable_router.HandleFunc("/weekly", middleware.TimetableAccessMiddleware(db, true, false, false)(handleGetWeeklyTimetable(db))).Methods("GET")
	// for intern: create timetable
	timetable_router.HandleFunc("", middleware.TimetableAccessMiddleware(db, false, true, false)(handleCreateTimeTable(db))).Methods("POST")
	// for admin: approve intern's timetable
	timetable_router.HandleFunc("/{timetable-id}/approve", middleware.TimetableAccessMiddleware(db, true, false, false)(handleApproveInternTimeTable(db))).Methods("POST")
	// for intern: checkin checkout
	timetable_router.HandleFunc("/{timetable-id}/attendance", middleware.TimetableAccessMiddleware(db, false, true, true)(handleClockinClockout(db))).Methods("PATCH")

}
