package repository

import "fmt"

func checkInternIDExists(store *timetableStore, inid string) error {
	var flag bool
	query := `select exists(select 1 from intern i join account acc on i.account_id = acc.id where i.id = ? and acc.deleted_at is null)`
	err := store.db.QueryRow(query, inid).Scan(&flag)
	if err != nil {
		return fmt.Errorf("error di in checkInternIDExists: %v", err)
	}
	if !flag {
		return fmt.Errorf("invalid-request: intern'id %v is not exists", inid)
	}
	return nil
}

func checkTimetableIDExists(store *timetableStore, timetable_id string) error {
	var flag bool
	query := `select exists(select 1 from timetable where id=? and deleted_at is null)`
	err := store.db.QueryRow(query, timetable_id).Scan(&flag)
	if err != nil {
		return fmt.Errorf("error di in checkInternIDExists: %v", err)
	}
	if !flag {
		return fmt.Errorf("invalid-request: intern'id %v is not exists", timetable_id)
	}
	return nil
}
