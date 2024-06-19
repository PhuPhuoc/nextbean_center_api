package repository

import (
	"fmt"
	"time"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/TimetableServices/model"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
)

func (store *timetableStore) CreateTimetable(inid string, info *model.TimtableCreation) error {
	if err_check_id := checkInternIDExists(store, inid); err_check_id != nil {
		return err_check_id
	}
	if err_date := checkDateTodayOrFuture(info.OfficeTime); err_date != nil {
		return err_date
	}
	if err_date_exist := checkDateAlreadyRegister(store, inid, info.OfficeTime); err_date_exist != nil {
		return err_date_exist
	}

	rawsql := `insert into timetable(intern_id, office_time, est_start, est_end, created_at) values (?,?,?,?,?)`
	result, err := store.db.Exec(rawsql, inid, info.OfficeTime, info.EstStart, info.EstEnd, utils.CreateDateTimeCurrentFormated())
	if err != nil {
		return fmt.Errorf("error in CreateTimetable: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error in CreateTimetable (check affect): %v", err)
	}
	if rowsAffected == 1 {
		return nil // created sucessfully
	} else {
		return fmt.Errorf("error in CreateTimetable (No user created): %v", err)
	}
}

func checkDateAlreadyRegister(store *timetableStore, inid, dateStr string) error {
	var flag bool
	query := `select exists(select 1 from timetable where intern_id=? and office_time=?)`
	err := store.db.QueryRow(query, inid, dateStr).Scan(&flag)
	if err != nil {
		return fmt.Errorf("error di in checkDateAlreadyRegister: %v", err)
	}
	if flag {
		return fmt.Errorf("invalid-request: office-time (%v) is already applied", dateStr)
	}
	return nil
}

func checkDateTodayOrFuture(dateStr string) error {
	// Parse the date string
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return fmt.Errorf("invalid-request: invalid date format %v", err)
	}
	// Get the current date
	now := time.Now()
	// Extract just the date part (without time)
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	// Compare the parsed date with today's date
	if date.Before(today) {
		return fmt.Errorf("invalid-request: Don't live in the past")
	}
	return nil
}
