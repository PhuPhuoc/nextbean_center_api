package repository

import (
	"fmt"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/TimetableServices/model"
)

func (store *timetableStore) AttendanceTimetable(inid, tid string, info *model.Attendance) error {
	if err_check_tid_exist := checkTimetableIDExists(store, tid); err_check_tid_exist != nil {
		return err_check_tid_exist
	}

	if err_check_verified := checkStatusVerified(store, tid); err_check_verified != nil {
		return err_check_verified
	}

	if err_check_status_attend := checkStatusAttendanceIsAbsentOrPresent(store, tid); err_check_status_attend != nil {
		return err_check_status_attend
	}

	if info.Clockin == "" && info.Clockout == "" {
		return fmt.Errorf("must update clock in or clock out")
	}

	rawsql := `update timetable set `
	params := []interface{}{}

	if info.Clockin != "" {
		rawsql = rawsql + `act_clockin=?, clockin_validated=?`
		params = append(params, info.Clockin, "admin-check")
	}
	if info.Clockout != "" {
		if info.Clockin != "" {
			rawsql = rawsql + `, `
		}
		rawsql = rawsql + `act_clockout=?, clockout_validated=?`
		params = append(params, info.Clockout, "admin-check")
	}
	rawsql = rawsql + ` where id=?`
	params = append(params, tid)
	result, err := store.db.Exec(rawsql, params...)

	if err != nil {
		return fmt.Errorf("error in AttendanceTimetable: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error in AttendanceTimetable (check affect): %v", err)
	}
	if rowsAffected == 0 {
		return nil // updated sucessfully
	}
	return nil
}
