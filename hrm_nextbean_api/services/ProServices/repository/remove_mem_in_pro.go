package repository

import (
	"fmt"

	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
)

func (store *projectStore) RemoveMemInPro(proid string, memid string) error {
	if err_check_pro_id := checkProjectIDExists(store, proid); err_check_pro_id != nil {
		return err_check_pro_id
	}

	if err_check_mem_id := checkMemIDExists(store, memid); err_check_mem_id != nil {
		return err_check_mem_id
	}

	flag_idNotExistInTable, err_check := checkMemberExistsInProject(store, proid, memid)
	if err_check != nil {
		return err_check
	}

	if !flag_idNotExistInTable {
		return fmt.Errorf("invalid-request: intern'id '%v' is not exists in project", memid)

	} else {
		flag_idExistButHasLeave, err_check_leave := checkMemberExistsInProjectButHasLeave(store, proid, memid)
		if err_check_leave != nil {
			return err_check_leave
		}

		if flag_idExistButHasLeave {
			return fmt.Errorf("invalid-request: intern'id '%v' already remove in project", memid)
		} else {
			if err_check_task := checkMemHasTask(store, proid, memid); err_check_task != nil {
				return err_check_task
			}
			rawsql := `update project_intern set leave_at=?, status='leave' where project_id=? and intern_id=?`
			result, err := store.db.Exec(rawsql, utils.CreateDateTimeCurrentFormated(), proid, memid)
			if err != nil {
				return fmt.Errorf("error in RemoveMemInPro: delete member in project: %v", err)
			}
			rowsAffected, err := result.RowsAffected()
			if err != nil {
				return fmt.Errorf("error when RemoveMemInPro in store (check affect): %v", err)
			}
			if rowsAffected == 1 {
				return nil // created sucessfully
			} else {
				return fmt.Errorf("error when RemoveMemInPro in store (member removed out of project): %v", err)
			}
		}
	}
}

func checkMemHasTask(store *projectStore, pro_id, mem_id string) error {
	var flag bool = false
	rawsql := `select exists(select 1 from task where project_id=? and assigned_to=? and status!='done')`
	if err_query := store.db.QueryRow(rawsql, pro_id, mem_id).Scan(&flag); err_query != nil {
		return fmt.Errorf("error in checkMemHasTask: %v", err_query)
	}
	if flag {
		return fmt.Errorf("invalid-request: intern'id '%v' still has task so cannot delete", mem_id)
	}
	return nil
}
