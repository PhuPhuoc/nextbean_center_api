package repository

import (
	"fmt"
)

func (store *accountStore) checkAccountIdExists(id string) error {
	var flag bool = false
	rawsql := `select exists(select 1 from account where id = ? and deleted_at is null)`
	if err_query := store.db.QueryRow(rawsql, id).Scan(&flag); err_query != nil {
		return fmt.Errorf("error in checkIdExist: %v", err_query)
	}
	if !flag {
		return fmt.Errorf("invalid-request: id '%v' not exist", id)
	}
	return nil // user'id exist in db => ready to update/delete
}

func (store *accountStore) checkEmailExistWithID(email, id string) error {
	var flag bool = false
	rawsql := `select exists(select 1 from account where email = ? and id != ?)`
	if err_query := store.db.QueryRow(rawsql, email, id).Scan(&flag); err_query != nil {
		return fmt.Errorf("error in checkEmailExistWithID: %v", err_query)
	}
	if flag {
		return fmt.Errorf("invalid-request: email '%v' already exists", email)
	}
	return nil // user'email not exist in db
}
