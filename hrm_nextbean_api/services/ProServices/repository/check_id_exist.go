package repository

import (
	"fmt"
)

func checkProjectIDExists(store *projectStore, pro_id string) error {
	var flag bool = false
	rawsql := `select exists(select 1 from project where id = ? and deleted_at is null)`
	if err_query := store.db.QueryRow(rawsql, pro_id).Scan(&flag); err_query != nil {
		return fmt.Errorf("error in checkProjectIDExists: %v", err_query)
	}
	if !flag {
		return fmt.Errorf("invalid-request: project'id '%v' does not exists", pro_id)
	}
	return nil // project'id exist in d
}

func checkPMIDExists(store *projectStore, pm_id string) error {
	var flag bool = false
	rawsql := `select exists(select 1 from account where id = ? and role='pm' and deleted_at is null)`
	if err_query := store.db.QueryRow(rawsql, pm_id).Scan(&flag); err_query != nil {
		return fmt.Errorf("error in checkPMIDExist: %v", err_query)
	}
	if !flag {
		return fmt.Errorf("invalid-request: pm'id '%v' does not exists'", pm_id)
	}
	return nil // pm_id'id exist in db
}

func checkMemIDExists(store *projectStore, mem_id string) error {
	var flag bool = false
	rawsql := `select exists(select 1 from intern i join account acc on i.account_id=acc.id where i.id = ? and acc.deleted_at is null)`
	if err_query := store.db.QueryRow(rawsql, mem_id).Scan(&flag); err_query != nil {
		return fmt.Errorf("error in checkMemIDExist: %v", err_query)
	}
	if !flag {
		return fmt.Errorf("invalid-request: intern'id '%v' does not exists", mem_id)
	}
	return nil // mem'id exist in db
}

func checkMemberExistsInProject(store *projectStore, project_id, mem_id string) (bool, error) {
	var flag bool = false
	rawsql := `select exists(select 1 from project_intern where project_id = ? and intern_id = ?)`
	if err_query := store.db.QueryRow(rawsql, project_id, mem_id).Scan(&flag); err_query != nil {
		return flag, fmt.Errorf("error in checkMemberExistInProject: %v", err_query)
	}
	return flag, nil
}

func checkMemberExistsInProjectButHasTerminated(store *projectStore, project_id, mem_id string) (bool, error) {
	var flag bool = false
	rawsql := `select exists(select 1 from project_intern where project_id = ? and intern_id = ? and status = 'terminated')`
	if err_query := store.db.QueryRow(rawsql, project_id, mem_id).Scan(&flag); err_query != nil {
		return flag, fmt.Errorf("error in checkMemberExistInProjectButHasTerminated: %v", err_query)
	}
	return flag, nil
}
