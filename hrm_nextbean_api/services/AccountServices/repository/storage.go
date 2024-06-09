package repository

import "database/sql"

type accountStore struct {
	db *sql.DB
}

func NewAccountStore(db *sql.DB) *accountStore {
	return &accountStore{db: db}
}
