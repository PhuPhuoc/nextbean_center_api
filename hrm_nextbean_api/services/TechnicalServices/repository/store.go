package repository

import "database/sql"

type techStore struct {
	db *sql.DB
}

func NewTechStore(db *sql.DB) *techStore {
	return &techStore{db: db}
}
