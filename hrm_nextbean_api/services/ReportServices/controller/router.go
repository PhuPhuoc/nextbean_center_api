package controller

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func RegisterReportRouter(r *mux.Router, db *sql.DB) {
	report_router := r.PathPrefix("/reports").Subrouter()
	report_router.HandleFunc("/{ojt-id}/project-intern", handleGetProjectReport(db)).Methods("GET")
}
