package repository

import (
	"fmt"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/TimetableServices/model"
)

func (store *timetableStore) ApproveTimetable(timetable_id string, info *model.ApproveTimetable) error {
	if err_check_id := checkTimetableIDExists(store, timetable_id); err_check_id != nil {
		return err_check_id
	}
	rawsql := `update timetable set status=? where id=?`
	result, err := store.db.Exec(rawsql, info.Status, timetable_id)
	if err != nil {
		return fmt.Errorf("error in ApproveTimetable: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error in ApproveTimetable (check affect): %v", err)
	}
	if rowsAffected == 1 {
		return nil // approve sucessfully
	} else {
		return fmt.Errorf("error in ApproveTimetable (No user updated): %v", err)
	}
}
