package repository

import (
	"fmt"

	query "github.com/PhuPhuoc/hrm_nextbean_api/rawsql/project_query"
)

func checkProjectIDExist(store *projectStore, pro_id string) error {
	var flag bool = false
	rawsql := query.QueryCheckProjectIDExist()
	if err_query := store.db.QueryRow(rawsql, pro_id).Scan(&flag); err_query != nil {
		return fmt.Errorf("error in store (check project-id exist): %v", err_query)
	}
	if !flag {
		return fmt.Errorf("project'id does not exist in db")
	}
	return nil // project'id exist in db => ready to update
}

func checkPMIDExist(store *projectStore, pm_id string) error {
	var flag bool = false
	rawsql := query.QueryCheckPMIDExist()
	if err_query := store.db.QueryRow(rawsql, pm_id).Scan(&flag); err_query != nil {
		return fmt.Errorf("error in store (check pm-id exist): %v", err_query)
	}
	if !flag {
		return fmt.Errorf("account'id: %v does not exist, or this id does not have the role 'project-manager'", pm_id)
	}
	return nil // pm_id'id exist in db
}

func checkMemIDExist(store *projectStore, mem_id string) error {
	var flag bool = false
	rawsql := query.QueryCheckMemIDExist()
	if err_query := store.db.QueryRow(rawsql, mem_id).Scan(&flag); err_query != nil {
		return fmt.Errorf("error in store (check mem_id exist): %v", err_query)
	}
	if !flag {
		return fmt.Errorf("intern'id: %v does not exist", mem_id)
	}
	return nil // mem'id exist in db
}

func checkMemberExistInProject(store *projectStore, project_id, mem_id string) (bool, error) {
	var flag bool = false
	rawsql := query.QueryCheckMemberInProjectNotExist()
	if err_query := store.db.QueryRow(rawsql, project_id, mem_id).Scan(&flag); err_query != nil {
		return flag, fmt.Errorf("error in store (check mem_id exist in project-member): %v", err_query)
	}
	return flag, nil
}

func checkMemberExistInProjectButHasLeave(store *projectStore, project_id, mem_id string) (bool, error) {
	var flag bool = false
	rawsql := query.QueryCheckMemberInProjectSatusLeave()
	if err_query := store.db.QueryRow(rawsql, project_id, mem_id).Scan(&flag); err_query != nil {
		return flag, fmt.Errorf("error in store (check mem_id exist but leave in project-member): %v", err_query)
	}
	return flag, nil
}
