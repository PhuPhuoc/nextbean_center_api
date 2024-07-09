package repository

import "fmt"

func checkInternIDExists(store *timetableStore, inid string) error {
	var flag bool
	query := `select exists(select 1 from intern i join account acc on i.account_id = acc.id where i.id = ? and acc.deleted_at is null)`
	err := store.db.QueryRow(query, inid).Scan(&flag)
	if err != nil {
		return fmt.Errorf("error in checkInternIDExists: %v", err)
	}
	if !flag {
		return fmt.Errorf("invalid-request: intern'id (%v) is not exists", inid)
	}
	return nil
}

func checkTimetableIDExists(store *timetableStore, timetable_id string) error {
	var flag bool
	query := `select exists(select 1 from timetable where id=? and deleted_at is null)`
	err := store.db.QueryRow(query, timetable_id).Scan(&flag)
	if err != nil {
		return fmt.Errorf("error in checkTimetableIDExists: %v", err)
	}
	if !flag {
		return fmt.Errorf("invalid-request: timetable'id %v is not exists", timetable_id)
	}
	return nil
}

func checkStatusVerified(store *timetableStore, timetable_id string) error {
	var flag bool
	query := `select exists(select 1 from timetable where id=? and verified='approved' and deleted_at is null)`
	err := store.db.QueryRow(query, timetable_id).Scan(&flag)
	if err != nil {
		return fmt.Errorf("error in checkStatusVerified: %v", err)
	}
	if !flag {
		return fmt.Errorf("invalid-request: Application not approved or has been rejected")
	}
	return nil
}

func checkStatusAttendanceIsAbsentOrPresent(store *timetableStore, timetable_id string) error {
	var flag bool
	query := `select exists(select 1 from timetable where id=? and status_attendance in ('absent','present') and deleted_at is null)`
	err := store.db.QueryRow(query, timetable_id).Scan(&flag)
	if err != nil {
		return fmt.Errorf("error in checkStatusAttendanceIsAbsentOrPresent: %v", err)
	}
	if flag {
		return fmt.Errorf("invalid-request: Status has been recorded, no further editing is possible")
	}
	return nil
}

func checkStatusAttendanceIsPresent(store *timetableStore, timetable_id string) error {
	var flag bool
	query := `select exists(select 1 from timetable where id=? and status_attendance='present' and deleted_at is null)`
	err := store.db.QueryRow(query, timetable_id).Scan(&flag)
	if err != nil {
		return fmt.Errorf("error in checkStatusAttendanceIsAbsentOrPresent: %v", err)
	}
	if flag {
		return fmt.Errorf("invalid-request: Status has been recorded ('present'), no further editing is possible")
	}
	return nil
}

func checkStatusAdminCheck(store *timetableStore, timetable_id, field string) error {
	var flag bool
	var query string
	if field == "clockin" {
		query = `select exists(select 1 from timetable where id=? and clockin_validated='admin-check' and deleted_at is null)`
	} else if field == "clockout" {
		query = `select exists(select 1 from timetable where id=? and clockout_validated='admin-check' and deleted_at is null)`
	} else {
		return fmt.Errorf("error in checkStatusAttendanceValidated: field must be 'clockin' or 'clockout'")

	}
	err := store.db.QueryRow(query, timetable_id).Scan(&flag)
	if err != nil {
		return fmt.Errorf("error in checkStatusAttendanceValidated: %v", err)
	}
	if !flag {
		return fmt.Errorf("invalid-request: the current state is not a state that the administrator can authenticate")
	}
	return nil
}

func checkAllFieldsHaveBeenAprrove(store *timetableStore, timetable_id string) (bool, error) {
	var flag bool
	query := `select exists(select 1 from timetable where id=? and clockin_validated='admin-approve' and clockout_validated='admin-approve' and deleted_at is null)`
	err := store.db.QueryRow(query, timetable_id).Scan(&flag)
	if err != nil {
		return false, fmt.Errorf("error in checkStatusAttendanceIsAbsentOrPresent: %v", err)
	}
	return flag, nil
}
