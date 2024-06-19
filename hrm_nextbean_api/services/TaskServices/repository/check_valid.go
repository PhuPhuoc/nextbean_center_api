package repository

import "fmt"

func checkInternIDExistsInProject(store *taskStore, proid, inid string) error {
	var flag bool = false
	rawsql := `select exists(select 1 from project_intern pin join intern i on pin.intern_id=i.id join account a on i.account_id=a.id where pin.project_id=? and pin.intern_id=? and a.deleted_at is null)`
	if err_query := store.db.QueryRow(rawsql, proid, inid).Scan(&flag); err_query != nil {
		return fmt.Errorf("error in checkInternIDExistsInProject: %v", err_query)
	}
	if !flag {
		return fmt.Errorf("invalid-request: member (id: %v) is not part of the project or the member's account has been deleted", inid)
	}
	return nil
}

func checkProjectIDExists(store *taskStore, proid string) error {
	var flag bool
	query := `select exists(select 1 from project where id = ? and deleted_at is null)`
	err := store.db.QueryRow(query, proid).Scan(&flag)
	if err != nil {
		return fmt.Errorf("error in checkProjectIDExists: %v", err)
	}
	if !flag {
		return fmt.Errorf("invalid-request: project'id %v is not exists", proid)
	}
	return nil
}
