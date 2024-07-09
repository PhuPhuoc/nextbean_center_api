package repository

import (
	"fmt"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/TimetableServices/model"
)

func (store *timetableStore) AttendanceValidated(timetable_id string, info *model.AttendanceValidated) error {
	if err_check_id := checkTimetableIDExists(store, timetable_id); err_check_id != nil {
		return err_check_id
	}
	if err_check_verified := checkStatusVerified(store, timetable_id); err_check_verified != nil {
		return err_check_verified
	}
	if err_check_status := checkStatusAttendanceIsAbsentOrPresent(store, timetable_id); err_check_status != nil {
		return err_check_status
	}
	if err_validate := checkStatusAdminCheck(store, timetable_id, info.ValidateField); err_validate != nil {
		return err_validate
	}

	rawsql := `update timetable set ` + info.ValidateField + `_validated=?` + ` where id=?`
	result, err := store.db.Exec(rawsql, "admin-approve", timetable_id)
	if err != nil {
		return fmt.Errorf("error in AttendanceValidated: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error in AttendanceValidated (check affect): %v", err)
	}
	if rowsAffected == 1 {
		err := changeStatusAttendanceAfterValidatedAllField(store, timetable_id)
		if err != nil {
			return err
		}
		return nil // approve sucessfully
	} else {
		return fmt.Errorf("error in AttendanceValidated (No user updated): %v", err)
	}
}

func changeStatusAttendanceAfterValidatedAllField(store *timetableStore, timetable_id string) error {
	flag, err := checkAllFieldsHaveBeenAprrove(store, timetable_id)
	if err != nil {
		return nil
	}
	if flag {
		rawsql := `update timetable set status_attendance=? where id=?`
		_, err := store.db.Exec(rawsql, "present", timetable_id)
		if err != nil {
			return fmt.Errorf("error in changeStatusAttendanceAfterValidatedAllField: %v", err)
		}
		return nil
	}
	return nil
}
