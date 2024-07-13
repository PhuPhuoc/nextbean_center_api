package repository

import (
	"fmt"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/TaskServices/model"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
)

func (store *taskStore) EndTask(proid, taskid, assigneeid string, info *model.DoneTask) error {
	if err_pro_exist := checkTaskIDExistsInProject(store, proid, taskid); err_pro_exist != nil {
		return err_pro_exist
	}
	if err_assignee_exist := checkAssigneeIDExistsInTask(store, taskid, assigneeid); err_assignee_exist != nil {
		return err_assignee_exist
	}
	if err_is_approved := isTaskHasBeenApproved(store, taskid); err_is_approved != nil {
		return err_is_approved
	}

	if err_status := isTaskHasBeenStart(store, taskid); err_status != nil {
		return err_status
	}

	if err_status := isTaskHasBeenDone(store, taskid); err_status != nil {
		return err_status
	}

	if info.ActualEffort <= 0 {
		return fmt.Errorf("you need to fill in the number of hours the task was performed")
	}

	rawsql := `update task set status='done', end_date=?, actual_effort=? where id=?`
	result, err := store.db.Exec(rawsql, utils.CreateDateTimeCurrentFormated(), info.ActualEffort, taskid)
	if err != nil {
		return fmt.Errorf("error when EndTask in store: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error when EndTask in store (check affect): %v", err)
	}
	if rowsAffected == 0 {
		return nil
	}
	return nil
}
