package repository

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/TimetableServices/model"
)

func (store *timetableStore) GetWeeklyTimetable(date string) ([7]model.Daily, error) {
	var data [7]model.Daily
	if err := getDays(date, data[:]); err != nil {
		return data, err
	}
	if err := queryGetDetailOfWeekDay(store, data[:]); err != nil {
		return data, err
	}
	return data, nil
}

func queryGetDetailOfWeekDay(store *timetableStore, data []model.Daily) error {
	for i := 0; i < 7; i++ {
		var (
			total_app sql.NullInt16
			total_pro sql.NullInt16
			total_den sql.NullInt16
		)
		rawquery := rawsqlGetRecord()
		if err_scan := store.db.QueryRow(rawquery, data[i].Date).Scan(&total_app, &total_pro, &total_den); err_scan != nil {
			return fmt.Errorf("error in GetWeeklyTimetable(queryGetDetailOfWeekDay): %v", err_scan)
		}

		if total_app.Valid {
			data[i].TotalApproved = int(total_app.Int16)
		} else {
			data[i].TotalApproved = 0
		}
		if total_pro.Valid {
			data[i].TotalWaiting = int(total_pro.Int16)
		} else {
			data[i].TotalWaiting = 0
		}
		if total_den.Valid {
			data[i].TotalDenied = int(total_den.Int16)
		} else {
			data[i].TotalDenied = 0
		}
	}
	return nil
}

func rawsqlGetRecord() string {
	var query strings.Builder
	query.WriteString(`select sum(case when status='approved' then 1 else 0 end) as app_re,`)
	query.WriteString(`sum(case when status='processing' then 1 else 0 end) as pro_re,`)
	query.WriteString(`sum(case when status='denied' then 1 else 0 end) as den_re`)
	query.WriteString(` from timetable where office_time=?`)
	return query.String()
}

func getDays(dateStr string, data []model.Daily) error {
	layout := "2006-01-02"
	date, err := time.Parse(layout, dateStr)
	if err != nil {
		return fmt.Errorf("error in GetWeeklyTimetable(getDays): %v", err)
	}

	for i := 0; i < 7; i++ {
		data[i].Date = date.Format(layout)
		data[i].WeekDay = date.Weekday().String()
		date = date.AddDate(0, 0, 1)
	}

	return nil
}
