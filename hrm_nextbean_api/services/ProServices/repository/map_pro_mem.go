package repository

import (
	"fmt"

	query "github.com/PhuPhuoc/hrm_nextbean_api/rawsql/project_query"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/model"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
)

func (store *projectStore) MapProMem(mapInfo *model.MapProMem) error {
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
		// create new map
		rawsql := query.QueryAddMemberToProject()
		result, err := store.db.Exec(rawsql, mapInfo.ProjectId, mapInfo.MemId, utils.CreateDateTimeCurrentFormated(), "inprogress")
		if err != nil {
			return fmt.Errorf("error when MapProMem in store: %v", err)
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return fmt.Errorf("error when MapProMem in store (check affect): %v", err)
		}
		if rowsAffected == 1 {
			return nil // created sucessfully
		} else {
			return fmt.Errorf("error when MapProMem in store (No member added into project): %v", err)
		}

	} else {
		flag_idExistButHasLeave, err_check_leave := checkMemberExistInProjectButHasLeave(store, mapInfo.ProjectId, mapInfo.MemId)
		if err_check_leave != nil {
			return err_check_leave
		}

		if flag_idExistButHasLeave {
			rawsql := query.QueryReJoinProject()
			result, err := store.db.Exec(rawsql, mapInfo.ProjectId, mapInfo.MemId)
			if err != nil {
				return fmt.Errorf("error when Re-MapProMem in store: %v", err)
			}
			rowsAffected, err := result.RowsAffected()
			if err != nil {
				return fmt.Errorf("error when Re-MapProMem in store (check affect): %v", err)
			}
			if rowsAffected == 1 {
				return nil // created sucessfully
			} else {
				return fmt.Errorf("error when Re-MapProMem in store (No user re-mapped): %v", err)
			}
		} else {
			return fmt.Errorf("member with id: %v already exist in project: %v", mapInfo.MemId, mapInfo.ProjectId)
		}
	}
}
