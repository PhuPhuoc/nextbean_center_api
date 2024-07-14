package repository

import (
	"database/sql"
	"time"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/TimetableServices/model"
)

func (store *timetableStore) GetTodayTimetable(inid string) (model.Today, error) {
	obj := model.Today{}
	now := time.Now()
	currentdate := now.Format("2006-01-02")

	var checkin, checkout *string
	rawsql := "select act_clockin, act_clockout from timetable where office_time=? and verified='approved' and intern_id=?"
	err_query := store.db.QueryRow(rawsql, currentdate, inid).Scan(&checkin, &checkout)
	if err_query != nil {
		if err_query == sql.ErrNoRows {
			return obj, nil
		}
		return obj, err_query
	}
	if checkin != nil && *checkin != "" {
		obj.Checkin = "checked in"
	} else {
		obj.Checkin = "not yet"
	}
	if checkout != nil && *checkout != "" {
		obj.Checkout = "checked out"
	} else {
		obj.Checkout = "not yet"
	}
	return obj, nil
}
