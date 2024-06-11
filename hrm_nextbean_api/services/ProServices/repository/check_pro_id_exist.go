package repository

import (
	"fmt"

	query "github.com/PhuPhuoc/hrm_nextbean_api/rawsql/project_query"
)

func checkProjectIDExist(store *projectStore, pro_id string) error {
	var flag bool = false
	rawsql := query.QueryCheckProjectIDExist()
	if err_query := store.db.QueryRow(rawsql, pro_id).Scan(&flag); err_query != nil {
		return fmt.Errorf("error when UpdateProject in store (check Id exist): %v", err_query)
	}
	if !flag {
		return fmt.Errorf("project'id does not exist in db")
	}
	return nil // project'id exist in db => ready to update
}
