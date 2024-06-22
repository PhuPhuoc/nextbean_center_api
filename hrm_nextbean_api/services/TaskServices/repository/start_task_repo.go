package repository

import (
	"fmt"

	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
)

func (store *taskStore) StartTask(proid, taskid, assigneeid string) error {
	if err_pro_exist := checkTaskIDExistsInProject(store, proid, taskid); err_pro_exist != nil {
		return err_pro_exist
	}
	if err_assignee_exist := checkAssigneeIDExistsInTask(store, taskid, assigneeid); err_assignee_exist != nil {
		return err_assignee_exist
	}

	if err_is_approved := isTaskHasBeenApproved(store, taskid); err_is_approved != nil {
		return err_is_approved
	}

	if err_status := isTaskHasBeenStartedOrDone(store, taskid); err_status != nil {
		return err_status
	}

	rawsql := `update task set status='inprogress', start_date=? where id=?`
	result, err := store.db.Exec(rawsql, utils.CreateDateTimeCurrentFormated(), taskid)
	if err != nil {
		return fmt.Errorf("error when StartTask in store: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error when StartTask in store (check affect): %v", err)
	}
	if rowsAffected == 0 {
		return nil
	}
	return nil
}
