package repository

import (
	"fmt"

	query "github.com/PhuPhuoc/hrm_nextbean_api/rawsql/project_query"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/model"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
)

func (store *projectStore) RemoveMemInPro(mapInfo *model.MapProMem) error {
	if err_check_pro_id := checkProjectIDExist(store, mapInfo.ProjectId); err_check_pro_id != nil {
		return err_check_pro_id
	}

	if err_check_mem_id := checkMemIDExist(store, mapInfo.MemId); err_check_mem_id != nil {
		return err_check_mem_id
	}

	flag_idNotExistInTable, err_check := checkMemberExistInProject(store, mapInfo.ProjectId, mapInfo.MemId)
	if err_check != nil {
		return err_check
	}

	if !flag_idNotExistInTable {
		return fmt.Errorf("intern'id: %v not does not exist in project", mapInfo.MemId)

	} else {
		flag_idExistButHasLeave, err_check_leave := checkMemberExistInProjectButHasLeave(store, mapInfo.ProjectId, mapInfo.MemId)
		if err_check_leave != nil {
			return err_check_leave
		}

		if flag_idExistButHasLeave {
			return fmt.Errorf("intern'id: %v already remove in project", mapInfo.MemId)
		} else {
			if err_check_task := checkMemHasTask(store, mapInfo.ProjectId, mapInfo.MemId); err_check_task != nil {
				return err_check_task
			}
			rawsql := query.QueryDeleteMemberInProject()
			result, err := store.db.Exec(rawsql, utils.CreateDateTimeCurrentFormated(), mapInfo.ProjectId, mapInfo.MemId)
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
