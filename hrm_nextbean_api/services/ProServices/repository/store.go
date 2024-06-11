package repository

import "database/sql"

type projectStore struct {
	db *sql.DB
}

func NewProjectStore(db *sql.DB) *projectStore {
	return &projectStore{db: db}
}
