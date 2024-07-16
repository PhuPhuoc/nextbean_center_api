package repository

import "database/sql"

type dashboardStore struct {
	db *sql.DB
}

func NewDashboardStore(db *sql.DB) *dashboardStore {
	return &dashboardStore{db: db}
}
