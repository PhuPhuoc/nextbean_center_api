package repository

import "database/sql"

type taskStore struct {
	db *sql.DB
}

func NewTaskStore(db *sql.DB) *taskStore {
	return &taskStore{db: db}
}
