package repository

import (
	"database/sql"
	"strconv"
	"strings"

	"github.com/PhuPhuoc/hrm_nextbean_api/common"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/TaskServices/model"
)

func (store *taskStore) GetTask(pagin *common.Pagination, filter *model.TaskFilter) ([]model.Task, error) {
	var total_record int64 = 0
	data := []model.Task{}
	rawsql, param := rawSqlGetProject(pagin, filter)

	rows, err_query := store.db.Query(rawsql, param...)
	if err_query != nil {
		return data, err_query
	}
	defer rows.Close()

	var (
		des     sql.NullString
		est_eff sql.NullString
		act_eff sql.NullString
		is_app  sql.NullBool
	)

	for rows.Next() {
		task := new(model.Task)
		if err_scan := rows.Scan(&task.Id, &task.ProjectId, &task.AssignedTo, &task.AssignedName, &task.AssignedCode, &is_app, &task.Status, &task.Name, &des, &est_eff, &act_eff, &total_record); err_scan != nil {
			return data, err_scan
		}

		if des.Valid {
			task.Description = des.String
		} else {
			task.Description = ""
		}
		if est_eff.Valid {
			task.EstimatedEffort = est_eff.String
		} else {
			task.EstimatedEffort = ""
		}
		if act_eff.Valid {
			task.ActualEffort = act_eff.String
		} else {
			task.ActualEffort = ""
		}
		if is_app.Valid {
			if is_app.Bool {
				task.IsApproved = "approved"
			} else {
				task.IsApproved = "waiting"
			}
		} else {
			task.IsApproved = "unknown"
		}

		data = append(data, *task)
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

func rawSqlGetProject(pagin *common.Pagination, filter *model.TaskFilter) (string, []interface{}) {
	where, param := whereClause(filter)
	main := mainClause(where, filter.OrderBy, pagin)

	// double param because in this code has 2 part of where clause ( 1 in cte - other in main select )
	doubledParams := make([]interface{}, len(param)*2)
	copy(doubledParams, param)
	copy(doubledParams[len(param):], param)

	return main, doubledParams
}

func mainClause(where, order string, pagin *common.Pagination) string {
	join := ` from task t join intern i on t.assigned_to = i.id join account a on i.account_id=a.id`
	if order == "" {
		order = "t.created_at desc"
	} else {
		order = "t." + order
	}
	var query strings.Builder
	query.WriteString(`with cte as ( select count(*) as total_record` + join + where + `) `)
	query.WriteString(`select t.id, t.project_id, t.assigned_to, a.user_name as assignee_name, i.student_code, t.is_approved, t.status, t.name, t.description, t.estimated_effort, t.actual_effort, cte.total_record`)
	query.WriteString(join + ` join cte` + where)
	query.WriteString(` order by ` + order)
	query.WriteString(` limit ` + strconv.Itoa(pagin.PSize))
	query.WriteString(` offset ` + strconv.Itoa((pagin.Page-1)*pagin.PSize))
	return query.String()
}

func whereClause(filter *model.TaskFilter) (string, []interface{}) {
	param := []interface{}{}
	var query strings.Builder
	query.WriteString(` where `)

	if filter.AssigneeId != "" {
		query.WriteString(`t.assigned_to = ? and `)
		param = append(param, filter.AssigneeId)
	}

	if filter.AssigneeName != "" {
		query.WriteString(`a.user_name like ? and `)
		param = append(param, filter.AssigneeName)
	}

	if filter.AssigneeCode != "" {
		query.WriteString(`i.student_code like ? and `)
		param = append(param, filter.AssigneeCode)
	}

	if filter.IsApproved != "" {
		if filter.IsApproved == "true" {
			query.WriteString(`t.is_approved = 1 and `)
		} else if filter.IsApproved == "false" {
			query.WriteString(`t.is_approved = 0 and `)
		}
	}

	if filter.Status != "" {
		query.WriteString(`t.status = ? and `)
		param = append(param, filter.Status)
	}

	if filter.Name != "" {
		query.WriteString(`t.name like ? and `)
		p := `%` + filter.Name + `%`
		param = append(param, p)
	}

	query.WriteString(`t.deleted_at is null `)
	return query.String(), param
}
