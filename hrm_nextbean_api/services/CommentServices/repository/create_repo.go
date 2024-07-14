package repository

import (
	"fmt"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/CommentServices/model"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
)

func (store *commentStore) CreateComment(taskid, role, accID, internID string, info *model.CommentCreation) error {
	if err_taskid_valid := checkTaskIDExists(store, taskid); err_taskid_valid != nil {
		return err_taskid_valid
	}
	if info.Type == "report" && role != "user" {
		return fmt.Errorf("only the person assigned the task is allowed to create the report")
	}

	rawsql := `insert into report(task_id, account_id, type, content, created_at) values (?,?,?,?,?)`
	result, err := store.db.Exec(rawsql, taskid, accID, info.Type, info.Content, utils.CreateDateTimeCurrentFormated())
	if err != nil {
		return fmt.Errorf("error in CreateTask: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error in CreateTask (check affect): %v", err)
	}
	if rowsAffected == 1 {
		return nil // created sucessfully
	} else {
		return fmt.Errorf("error in CreateTask (No task created): %v", err)
	}
}

