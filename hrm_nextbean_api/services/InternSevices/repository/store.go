package repository

import "database/sql"

type internStore struct {
	db *sql.DB
}

func NewInternStore(db *sql.DB) *internStore {
	return &internStore{db: db}
}
