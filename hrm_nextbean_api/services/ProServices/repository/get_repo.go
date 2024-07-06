package repository

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PhuPhuoc/hrm_nextbean_api/common"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/model"
)

func (store *projectStore) GetProject(pagin *common.Pagination, filter *model.ProjectFilter) ([]model.Project, error) {
	var total_record int64 = 0
	data := []model.Project{}
	rawsql, param := rawSqlGetProject(pagin, filter)

	fmt.Println("rawsql: ", rawsql)

	rows, err_query := store.db.Query(rawsql, param...)
	if err_query != nil {
		return data, err_query
	}
	defer rows.Close()

	for rows.Next() {
		p := new(model.Project)
		if err_scan := rows.Scan(&p.Id, &p.Name, &p.Status, &p.Description, &p.EstStartTime, &p.EstCompletionTime, &total_record); err_scan != nil {
			return data, err_scan
		}
		fmt.Println("time: ", p.EstStartTime)
		data = append(data, *p)
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

func rawSqlGetProject(pagin *common.Pagination, filter *model.ProjectFilter) (string, []interface{}) {
	where, param := whereClause(filter)
	main := mainClause(where, filter.OrderBy, pagin)

	// double param because in this code has 2 part of where clause ( 1 in cte - other in main select )
	doubledParams := make([]interface{}, len(param)*2)
	copy(doubledParams, param)
	copy(doubledParams[len(param):], param)

	return main, doubledParams
}

func mainClause(where, order string, pagin *common.Pagination) string {
	if order == "" {
		order = "est_start_time desc"
	}
	var query strings.Builder
	query.WriteString(`with cte as ( select count(*) as total_record from project` + where + `) `)
	query.WriteString(`select id, name, status, description, est_start_time, est_completion_time, cte.total_record from project join cte` + where)
	query.WriteString(`order by ` + order)
	query.WriteString(` limit ` + strconv.Itoa(pagin.PSize))
	query.WriteString(` offset ` + strconv.Itoa((pagin.Page-1)*pagin.PSize))
	return query.String()
}

func whereClause(filter *model.ProjectFilter) (string, []interface{}) {
	param := []interface{}{}
	var query strings.Builder
	query.WriteString(` where `)

	if filter.Role == "pm" {
		query.WriteString(`id in (select project_id from project_manager where account_id='` + filter.AccId + `') and `)
	}

	if filter.Role == "user" {
		query.WriteString(`id in (select project_id from project_intern where intern_id=(select id from intern where account_id ='` + filter.AccId + `')) and `)
	}

	if filter.Name != "" {
		query.WriteString(`name like ? and `)
		p := `%` + filter.Name + `%`
		param = append(param, p)
	}

	if filter.Status != "" {
		query.WriteString(`status = ? and `)
		param = append(param, filter.Status)
	}

	if filter.EstStartTimeFrom != "" {
		query.WriteString(`est_start_time >= ? and `)
		param = append(param, filter.EstStartTimeFrom)
	}

	if filter.EstStartTimeTo != "" {
		query.WriteString(`est_start_time <= ? and `)
		param = append(param, filter.EstStartTimeTo)
	}

	if filter.EstCompletionTimeFrom != "" {
		query.WriteString(`est_completion_time >= ? and `)
		param = append(param, filter.EstCompletionTimeFrom)
	}

	if filter.EstCompletionTimeTo != "" {
		query.WriteString(`est_completion_time <= ? and `)
		param = append(param, filter.EstCompletionTimeTo)
	}

	query.WriteString(`deleted_at is null `)
	return query.String(), param
}
