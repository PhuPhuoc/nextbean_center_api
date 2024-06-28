package repository

import "database/sql"

type loginGGStore struct {
	db *sql.DB
}

func NewLoginGGStore(db *sql.DB) *loginGGStore {
	return &loginGGStore{db: db}
}
