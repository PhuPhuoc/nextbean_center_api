package repository

import "fmt"

func checkTaskIDExists(store *commentStore, taskid string) error {
	var flag bool
	query := `select exists(select 1 from task where id = ? and deleted_at is null)`
	err := store.db.QueryRow(query, taskid).Scan(&flag)
	if err != nil {
		return fmt.Errorf("error in checkTaskIDExists: %v", err)
	}
	if !flag {
		return fmt.Errorf("invalid-request: task'id %v is not exists or has been deleted", taskid)
	}
	return nil
}
