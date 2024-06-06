package repository

import "database/sql"

type InternStore struct {
	db *sql.DB
}

func NewInternStore(db *sql.DB) *InternStore {
	return &InternStore{db: db}
}
