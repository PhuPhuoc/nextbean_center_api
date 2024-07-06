package repository

import (
	"database/sql"
	"strconv"
	"strings"

	"github.com/PhuPhuoc/hrm_nextbean_api/common"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/TimetableServices/model"
)

func (store *timetableStore) GetTimetable(pagin *common.Pagination, filter *model.TimeTableFilter) ([]model.Timtable, error) {
	var total_record int64 = 0
	data := []model.Timtable{}
	rawsql, param := queryGetTimetable(pagin, filter)

	rows, err_query := store.db.Query(rawsql, param...)
	if err_query != nil {
		return data, err_query
	}
	defer rows.Close()

	for rows.Next() {
		t := new(model.Timtable)

		var (
			actClockin  sql.NullString
			actClockout sql.NullString
		)
		if err_scan := rows.Scan(&t.Id, &t.InternName, &t.StudentCode, &t.OfficeTime, &t.Verified, &t.EstStartTime, &t.EstEndTime, &actClockin, &t.ClockinValidated, &actClockout, &t.ClockoutValidated, &t.StatusAttendance, &total_record); err_scan != nil {
			return data, err_scan
		}

		if actClockin.Valid {
			t.ActClockin = actClockin.String
		} else {
			t.ActClockin = ""
		}
		if actClockout.Valid {
			t.ActClockout = actClockout.String
		} else {
			t.ActClockout = ""
		}
		data = append(data, *t)
	}

	pagin.Items = total_record
	per := pagin.Items % int64(pagin.PSize)
	if per > 0 {
		pagin.Pages = pagin.Items/int64(pagin.PSize) + 1
	} else {
		pagin.Pages = pagin.Items / int64(pagin.PSize)
	}

	return data, nil
}

func queryGetTimetable(pagin *common.Pagination, filter *model.TimeTableFilter) (string, []interface{}) {
	where, param := createConditionClause(filter)
	query := rawSqlSelectAccount(where, filter.OrderBy, pagin)

	// double param because in this code has 2 part of where clause ( 1 in cte - other in main select )
	doubledParams := make([]interface{}, len(param)*2)
	copy(doubledParams, param)
	copy(doubledParams[len(param):], param)

	return query, doubledParams
}

func rawSqlSelectAccount(where string, order string, pagin *common.Pagination) string {
	var query strings.Builder
	if order == "" {
		order = "t.office_time desc"
	} else {
		order = "t." + order
	}
	join := ` from timetable t join intern i on t.intern_id=i.id join account a on i.account_id=a.id `
	query.WriteString(`with cte as (select count(*) as total_record` + join + where + `) `)
	query.WriteString(`select t.id, a.user_name, i.student_code, t.office_time, t.verified, t.est_start_time, t.est_end_time, t.act_clockin, t.clockin_validated, t.act_clockout, t.clockout_validated, t.status_attendance ,cte.total_record`)
	query.WriteString(join + ` join cte`)
	query.WriteString(where)
	query.WriteString(` order by ` + order)
	query.WriteString(` limit ` + strconv.Itoa(pagin.PSize))
	query.WriteString(` offset ` + strconv.Itoa((pagin.Page-1)*pagin.PSize))
	return query.String()
}

func createConditionClause(filter *model.TimeTableFilter) (string, []interface{}) {
	param := []interface{}{}
	var query strings.Builder
	query.WriteString(` where `)

	if filter.Role == "user" {
		query.WriteString(`a.id = ? and `)
		param = append(param, filter.Account_id)
	}

	if filter.Id != "" {
		query.WriteString(`t.id = ? and `)
		param = append(param, filter.Id)
	}
	if filter.InternName != "" {
		query.WriteString(`a.user_name like ? and `)
		p := `%` + filter.InternName + `%`
		param = append(param, p)
	}
	if filter.StudentCode != "" {
		query.WriteString(`i.student_code like ? and `)
		p := `%` + filter.InternName + `%`
		param = append(param, p)
	}
	if filter.Verified != "" {
		query.WriteString(`t.verified in (`)
		parts := strings.Split(filter.Verified, "-")
		for i, part := range parts {
			if i > 0 {
				query.WriteString(`,`)
			}
			query.WriteString(`?`)
			param = append(param, part)
		}
		query.WriteString(`) and `)
	}

	if filter.StatusAttendance != "" {
		query.WriteString(`t.status-attendance in (`)
		parts := strings.Split(filter.StatusAttendance, "-")
		for i, part := range parts {
			if i > 0 {
				query.WriteString(`,`)
			}
			query.WriteString(`?`)
			param = append(param, part)
		}
		query.WriteString(`) and `)
	}

	if filter.OfficeTimeFrom != "" {
		query.WriteString(`t.office_time >= ? and `)
		param = append(param, filter.OfficeTimeFrom)
	}

	if filter.OfficeTimeTo != "" {
		query.WriteString(`t.office_time <= ? and `)
		param = append(param, filter.OfficeTimeTo)
	}

	query.WriteString(`t.deleted_at is null`)
	return query.String(), param
}
