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
		return fmt.Errorf("invalid-request: project'id %v is not exists or this project has been deleted", proid)
	}
	return nil
}

func checkTaskIDExistsInProject(store *taskStore, proid, taskid string) error {
	var flag bool
	query := `select exists(select 1 from task where project_id=? and id = ? and deleted_at is null)`
	err := store.db.QueryRow(query, proid, taskid).Scan(&flag)
	if err != nil {
		return fmt.Errorf("error in checkTaskIDExistsInProject: %v", err)
	}
	if !flag {
		return fmt.Errorf("invalid-request: task'id %v is not exists in project or this project has been deleted", taskid)
	}
	return nil
}

func checkAssigneeIDExistsInTask(store *taskStore, taskid, inid string) error {
	var flag bool = false
	rawsql := `select exists(select 1 from task where id=? and assigned_to=? and deleted_at is null)`
	if err_query := store.db.QueryRow(rawsql, taskid, inid).Scan(&flag); err_query != nil {
		return fmt.Errorf("error in checkInternIDExistsInProject: %v", err_query)
	}
	if !flag {
		return fmt.Errorf("invalid-request: member (id: %v) is not the person who assigned this task or this task has been deleted", inid)
	}
	return nil
}

func isTaskHasBeenApproved(store *taskStore, taskid string) error {
	var flag bool = false
	rawsql := `select exists(select 1 from task where is_approved=1 and id=? and deleted_at is null)`
	if err_query := store.db.QueryRow(rawsql, taskid).Scan(&flag); err_query != nil {
		return fmt.Errorf("error in isTaskHasBeenApproved: %v", err_query)
	}
	if flag {
		return nil
	}
	return fmt.Errorf("invalid-request: This task has not been approved by PM")
}

func isTaskHasBeenStartedOrDone(store *taskStore, taskid string) error {
	var flag bool = false
	rawsql := `select exists(select 1 from task where (status='inprogress' or status='done') and id=? and deleted_at is null)`
	if err_query := store.db.QueryRow(rawsql, taskid).Scan(&flag); err_query != nil {
		return fmt.Errorf("error in isTaskHasBeenApproved: %v", err_query)
	}
	if !flag {
		return nil
	}
	return fmt.Errorf("invalid-request: This task has been started or done")
}

func isTaskHasBeenDone(store *taskStore, taskid string) error {
	var flag bool = false
	rawsql := `select exists(select 1 from task where status='done' and id=? and deleted_at is null)`
	if err_query := store.db.QueryRow(rawsql, taskid).Scan(&flag); err_query != nil {
		return fmt.Errorf("error in isTaskHasBeenDone: %v", err_query)
	}
	if !flag {
		return nil
	}
	return fmt.Errorf("invalid-request: This task has been completed")
}

func isTaskHasBeenStart(store *taskStore, taskid string) error {
	var flag bool = false
	rawsql := `select exists(select 1 from task where status='inprogress' and id=? and deleted_at is null)`
	if err_query := store.db.QueryRow(rawsql, taskid).Scan(&flag); err_query != nil {
		return fmt.Errorf("error in isTaskHasBeenDone: %v", err_query)
	}
	if flag {
		return nil
	}
	return fmt.Errorf("invalid-request: This task has not started, so completion cannot be confirmed")
}
