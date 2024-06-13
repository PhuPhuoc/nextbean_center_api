package repository

import (
	"fmt"

	query "github.com/PhuPhuoc/hrm_nextbean_api/rawsql/account_query"
)

func (store *accountStore) checkIdExist(id string) error {
	var flag bool = false
	rawsql := query.QueryIdExist()
	if err_query := store.db.QueryRow(rawsql, id).Scan(&flag); err_query != nil {
		return fmt.Errorf("error when (check Id exist): %v", err_query)
	}
	if !flag {
		return fmt.Errorf("not_exist_id")
	}
	return nil // user'id exist in db => ready to update/delete
}

func (store *accountStore) checkEmailExistWithID(email, id string) error {
	var flag bool = false
	rawsql := query.QueryCheckExistEmailWithID()
	if err_query := store.db.QueryRow(rawsql, email, id).Scan(&flag); err_query != nil {
		return fmt.Errorf("error when UpdateAccount in store (check exist email): %v", err_query)
	}
	if flag {
		return fmt.Errorf("duplicate_data_email")
	}
	return nil // user'email not exist in db
}
