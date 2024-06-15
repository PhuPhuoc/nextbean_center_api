package repository

import (
	"fmt"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/model"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
)

func (store *projectStore) MapProMem(proid string, mapInfo *model.MapProMem) error {
	if err_check_pro_id := checkProjectIDExists(store, proid); err_check_pro_id != nil {
		return err_check_pro_id
	}

	if err_check_mem_id := checkMemIDExists(store, mapInfo.MemId); err_check_mem_id != nil {
		return err_check_mem_id
	}

	flag_idNotExistInTable, err_check := checkMemberExistsInProject(store, proid, mapInfo.MemId)
	if err_check != nil {
		return err_check
	}

	if !flag_idNotExistInTable {
		// create new map
		rawsql := `insert into project_intern(project_id, intern_id, join_at, status) values (?,?,?,?)`
		result, err := store.db.Exec(rawsql, proid, mapInfo.MemId, utils.CreateDateTimeCurrentFormated(), "inprogress")
		if err != nil {
			return fmt.Errorf("error in MapProMem add member to project: %v", err)
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return fmt.Errorf("error in MapProMem (check affect): %v", err)
		}
		if rowsAffected == 1 {
			return nil // created sucessfully
		} else {
			return fmt.Errorf("error in MapProMem (No member added into project): %v", err)
		}

	} else {
		flag_idExistButHasLeave, err_check_leave := checkMemberExistsInProjectButHasLeave(store, proid, mapInfo.MemId)
		if err_check_leave != nil {
			return err_check_leave
		}

		if flag_idExistButHasLeave {
			rawsql := `update project_intern set leave_at=null, status='inprogress' where project_id=? and intern_id=?`
			result, err := store.db.Exec(rawsql, proid, mapInfo.MemId)
			if err != nil {
				return fmt.Errorf("error in Re-MapProMem: %v", err)
			}
			rowsAffected, err := result.RowsAffected()
			if err != nil {
				return fmt.Errorf("error in Re-MapProMem (check affect): %v", err)
			}
			if rowsAffected == 1 {
				return nil // created sucessfully
			} else {
				return fmt.Errorf("error in Re-MapProMem (No user re-mapped): %v", err)
			}
		} else {
			return fmt.Errorf("invalid-request: member with id '%v' already exist in project: %v", mapInfo.MemId, proid)
		}
	}
}
