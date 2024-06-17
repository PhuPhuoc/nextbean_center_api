package repository

import "database/sql"

type timetableStore struct {
	db *sql.DB
}

func NewTimeTableStore(db *sql.DB) *timetableStore {
	return &timetableStore{db: db}
}
