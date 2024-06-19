package repository

import (
	"fmt"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/TaskServices/model"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
	"github.com/google/uuid"
)

func (store *taskStore) CreateTask(role, proid, internIDWhoRegisteredTheTask string, info *model.TaskCreation) error {
	if role != "pm" && role != "user" {
		return fmt.Errorf("invalid-request: something went wrong - can not role cannot be determined")
	}

	if err_pro_id_not_exist := checkProjectIDExists(store, proid); err_pro_id_not_exist != nil {
		return err_pro_id_not_exist
	}

	if err_inid_not_exist_in_project := checkInternIDExistsInProject(store, proid, info.AssignedTo); err_inid_not_exist_in_project != nil {
		return err_inid_not_exist_in_project
	}
	var is_aprroved int
	if role == "pm" {
		is_aprroved = 1
	} else if role == "user" {
		is_aprroved = 0
		if internIDWhoRegisteredTheTask != info.AssignedTo {
			return fmt.Errorf("invalid-request: you can only create tasks for yourself")
		}
	}
	newUUID := uuid.New()
	rawsql := `insert into task(id, project_id, assigned_to, is_approved, name, description, estimated_effort, created_at) values (?,?,?,?,?,?,?,?)`
	result, err := store.db.Exec(rawsql, newUUID, proid, info.AssignedTo, is_aprroved, info.Name, info.Description, info.EstimatedEffort, utils.CreateDateTimeCurrentFormated())
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
