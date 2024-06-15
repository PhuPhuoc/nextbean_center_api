package repository

import (
	"strconv"
	"strings"

	"github.com/PhuPhuoc/hrm_nextbean_api/common"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/model"
)

func (store *projectStore) GetProject(pagin *common.Pagination, filter *model.ProjectFilter) ([]model.Project, error) {
	var total_record int64 = 0
	data := []model.Project{}
	rawsql, param := rawSqlGetProject(pagin, filter)

	rows, err_query := store.db.Query(rawsql, param...)
	if err_query != nil {
		return data, err_query
	}
	defer rows.Close()

	for rows.Next() {
		acc := new(model.Project)
		if err_scan := rows.Scan(&acc.Id, &acc.Name, &acc.Status, &acc.Description, &acc.Duration, &acc.StartDate, &total_record); err_scan != nil {
			return data, err_scan
		}
		data = append(data, *acc)
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
		order = "created_at desc"
	}
	var query strings.Builder
	query.WriteString(`with cte as ( select count(*) as total_record from project` + where + `) `)
	query.WriteString(`select id, name, status, description, duration, start_date, cte.total_record from project join cte` + where)
	query.WriteString(`order by ` + order)
	query.WriteString(` limit ` + strconv.Itoa(pagin.PSize))
	query.WriteString(` offset ` + strconv.Itoa((pagin.Page-1)*pagin.PSize))
	return query.String()
}

func whereClause(filter *model.ProjectFilter) (string, []interface{}) {
	param := []interface{}{}
	var query strings.Builder
	query.WriteString(` where `)

	if filter.Name != "" {
		query.WriteString(`name like ? and `)
		p := `%` + filter.Name + `%`
		param = append(param, p)
	}

	if filter.Status != "" {
		query.WriteString(`status = ? and `)
		param = append(param, filter.Status)
	}

	if filter.StartDateFrom != "" {
		query.WriteString(`start_date > ? and `)
		param = append(param, filter.StartDateFrom)
	}

	if filter.StarttDateTo != "" {
		query.WriteString(`start_date < ? and `)
		param = append(param, filter.StarttDateTo)
	}

	query.WriteString(`deleted_at is null `)
	return query.String(), param
}
