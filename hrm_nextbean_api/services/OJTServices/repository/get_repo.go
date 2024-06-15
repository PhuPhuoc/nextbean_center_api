package repository

import (
	"strconv"
	"strings"

	"github.com/PhuPhuoc/hrm_nextbean_api/common"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/OJTServices/model"
)

func (store *ojtStore) GetOJT(pagin *common.Pagination, filter *model.FilterOJT) ([]model.OJT, error) {
	var total_record int64 = 0
	data := []model.OJT{}
	rawsql, param := rawSqlGetOJT(pagin, filter)

	rows, err_query := store.db.Query(rawsql, param...)
	if err_query != nil {
		return data, err_query
	}
	defer rows.Close()

	for rows.Next() {
		ojt := new(model.OJT)
		if err_scan := rows.Scan(&ojt.Id, &ojt.Semester, &ojt.University, &ojt.StartAt, &ojt.EndAt, &total_record); err_scan != nil {
			return data, err_scan
		}
		data = append(data, *ojt)
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

func rawSqlGetOJT(pagin *common.Pagination, filter *model.FilterOJT) (string, []interface{}) {
	where, param := whereClause(filter)
	main := mainClause(where, filter.OrderBy, pagin)

	// double param because in this code has 2 part of where clause ( 1 in cte - other in main select )
	doubledParams := make([]interface{}, len(param)*2)
	copy(doubledParams, param)
	copy(doubledParams[len(param):], param)
	return main, doubledParams
}

func whereClause(filter *model.FilterOJT) (string, []interface{}) {
	param := []interface{}{}
	var query strings.Builder
	query.WriteString(` where `)

	if filter.Id != 0 {
		query.WriteString(`id = ? and `)
		param = append(param, filter.Id)
	}
	if filter.Semester != "" {
		query.WriteString(`semester like ? and `)
		p := `%` + filter.Semester + `%`
		param = append(param, p)

	}
	if filter.University != "" {
		query.WriteString(`university like ? and `)
		p := `%` + filter.University + `%`
		param = append(param, p)

	}

	query.WriteString(`deleted_at is null `)
	return query.String(), param
}

func mainClause(where, order string, pagin *common.Pagination) string {
	if order == "" {
		order = "created_at desc"
	}
	var query strings.Builder
	query.WriteString(`with cte as ( select count(*) as total_record from ojt` + where + `) `)
	query.WriteString(`select id, semester, university, start_at, end_at, cte.total_record from ojt join cte`)
	query.WriteString(where)
	query.WriteString(`order by ` + order)
	query.WriteString(` limit ` + strconv.Itoa(pagin.PSize))
	query.WriteString(` offset ` + strconv.Itoa((pagin.Page-1)*pagin.PSize))
	return query.String()
}
