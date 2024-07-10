package repository

import "database/sql"

type reportStore struct {
	db *sql.DB
}

func NewReportStore(db *sql.DB) *reportStore {
	return &reportStore{db: db}
}
