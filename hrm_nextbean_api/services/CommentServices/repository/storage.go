package repository

import "database/sql"

type commentStore struct {
	db *sql.DB
}

func NewCommentStore(db *sql.DB) *commentStore {
	return &commentStore{db: db}
}
