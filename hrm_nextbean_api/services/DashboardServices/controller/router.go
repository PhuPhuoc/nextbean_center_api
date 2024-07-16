package controller

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func RegisterDashboardRouter(r *mux.Router, db *sql.DB) {
	dashboard_router := r.PathPrefix("/dashboards").Subrouter()
	dashboard_router.HandleFunc("/total-number", handleDashboardGetTotalNumber(db)).Methods("GET")
	dashboard_router.HandleFunc("/inprogress-ojt", handleDashboardGetInProgressOJT(db)).Methods("GET")
}
