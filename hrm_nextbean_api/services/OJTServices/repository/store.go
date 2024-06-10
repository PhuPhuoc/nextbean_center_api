package repository

import "database/sql"

type ojtStore struct {
	db *sql.DB
}

func NewOjtStore(db *sql.DB) *ojtStore {
	return &ojtStore{db: db}
}
