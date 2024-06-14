package repository

import (
	"fmt"

	query "github.com/PhuPhuoc/hrm_nextbean_api/rawsql/project_query"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
)

func (store *projectStore) RemoveMemInPro(proid string, memid string) error {
	if err_check_pro_id := checkProjectIDExist(store, proid); err_check_pro_id != nil {
		return err_check_pro_id
	}

	if err_check_mem_id := checkMemIDExist(store, memid); err_check_mem_id != nil {
		return err_check_mem_id
	}

	flag_idNotExistInTable, err_check := checkMemberExistInProject(store, proid, memid)
	if err_check != nil {
		return err_check
	}

	if !flag_idNotExistInTable {
		return fmt.Errorf("intern'id: %v not does not exist in project", memid)

	} else {
		flag_idExistButHasLeave, err_check_leave := checkMemberExistInProjectButHasLeave(store, proid, memid)
		if err_check_leave != nil {
			return err_check_leave
		}

		if flag_idExistButHasLeave {
			return fmt.Errorf("intern'id: %v already remove in project", memid)
		} else {
			if err_check_task := checkMemHasTask(store, proid, memid); err_check_task != nil {
				return err_check_task
			}
			rawsql := query.QueryDeleteMemberInProject()
			result, err := store.db.Exec(rawsql, utils.CreateDateTimeCurrentFormated(), proid, memid)
			if err != nil {
				return fmt.Errorf("error when RemoveMemInPro in store: %v", err)
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
	rawsql := query.QueryCheckMemTaskBeforeDelete()
	if err_query := store.db.QueryRow(rawsql, pro_id, mem_id).Scan(&flag); err_query != nil {
		return fmt.Errorf("error in store (check mem task before remove): %v", err_query)
	}
	if flag {
		return fmt.Errorf("intern'id: %v still has task so cannot delete", mem_id)
	}
	return nil
}
