package repository

import (
	"fmt"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/TaskServices/model"
)

func (store *taskStore) UpdateTask(proid, taskid string, info *model.TaskUpdate) error {
	if err := checkTaskIDExistsInProject(store, proid, taskid); err != nil {
		return err
	}

	newAssignee, err_check := isNewAssignee(store, taskid, info.AssignedTo)
	if err_check != nil {
		return err_check
	}

	status := info.Status
	if newAssignee {
		status = "todo"
		if err_exist := checkInternIDExistsInProject(store, proid, info.AssignedTo); err_exist != nil {
			return err_exist
		}
	}

	rawsql := `update task set assigned_to=?, is_approved=?, status=?, name=?, description=?, estimated_effort=? where id=?`
	result, err := store.db.Exec(rawsql, info.AssignedTo, info.IsApproved, status, info.Name, info.Description, info.EstimatedEffort, taskid)
	if err != nil {
		return fmt.Errorf("error when UpdateTask in store: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error when UpdateTask in store (check affect): %v", err)
	}
	if rowsAffected == 0 {
		return nil
	}
	return nil
}

func isNewAssignee(store *taskStore, taskid, assignee string) (bool, error) {
	var flag bool
	query := `select exists(select 1 from task where assigned_to=? and id = ? and deleted_at is null)`
	err := store.db.QueryRow(query, assignee, taskid).Scan(&flag)
	if err != nil {
		return false, fmt.Errorf("error in isNewAssignee: %v", err)
	}
	if flag {
		return false, nil
	}
	return true, nil
}
