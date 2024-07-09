package repository

import (
	"fmt"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/TimetableServices/model"
)

func (store *timetableStore) SetStatusAbsent(timetable_id string, info *model.StatusAbsent) error {
	if err_check_tid_exist := checkTimetableIDExists(store, timetable_id); err_check_tid_exist != nil {
		return err_check_tid_exist
	}
	if err_check_verified := checkStatusVerified(store, timetable_id); err_check_verified != nil {
		return err_check_verified
	}
	if err_check_status := checkStatusAttendanceIsPresent(store, timetable_id); err_check_status != nil {
		return err_check_status
	}

	rawsql := `update timetable set status_attendance=? where id=?`
	if info.Status == "absent" {
		_, err := store.db.Exec(rawsql, "absent", timetable_id)
		if err != nil {
			return fmt.Errorf("error in SetStatusAbsent: %v", err)
		}
	} else if info.Status == "remove" {
		_, err := store.db.Exec(rawsql, "not-yet", timetable_id)
		if err != nil {
			return fmt.Errorf("error in SetStatusAbsent: %v", err)
		}
	} else {
		return fmt.Errorf("error in SetStatusAbsent: invalid status")
	}
	return nil
}
