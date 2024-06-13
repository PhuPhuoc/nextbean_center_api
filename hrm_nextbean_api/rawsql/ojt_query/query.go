package ojtquery

import (
	"strconv"
	"strings"

	"github.com/PhuPhuoc/hrm_nextbean_api/common"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/OJTServices/model"
)

func QueryCreateNewOJT() string {
	return `insert into ojt(semester, university, start_at, end_at, created_at) values (?,?,?,?,?)`
}

func QueryCheckExistID() string {
	return `select exists(select 1 from ojt where id = ? and deleted_at is null)`
}

func QueryUpdateOJT() string {
	return `update ojt set semester = ?, university = ?, start_at = ?, end_at = ? where id = ?`
}

func QueryDeleteOJT() string {
	return `update ojt set deleted_at = ? where id = ?`
}

func QueryGetOJT(pagin *common.Pagination, filter *model.FilterOJT) (string, []interface{}) {
	var query strings.Builder
	where, param := createConditionClause(filter)
	cte := createCTEClause(where)
	main := createSelectClause(where)
	ord := createOrderByClause(filter.OrderBy)
	pag := createPaginationClause(pagin)

	// double param because in this code has 2 part of where clause ( 1 in cte - other in main select )
	doubledParams := make([]interface{}, len(param)*2)
	copy(doubledParams, param)
	copy(doubledParams[len(param):], param)

	query.WriteString(cte + main + ord + pag)
	return query.String(), doubledParams
}

func createConditionClause(filter *model.FilterOJT) (string, []interface{}) {
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

func createCTEClause(condition_clause string) string {
	var query strings.Builder
	query.WriteString(`with cte as ( select count(*) as total_record from ojt` + condition_clause + `)`)
	return query.String()
}

func createSelectClause(condition_clause string) string {
	var query strings.Builder
	query.WriteString(` select id, semester, university, start_at, end_at, cte.total_record from ojt`)
	query.WriteString(` join cte `)
	query.WriteString(condition_clause)
	return query.String()
}

func createOrderByClause(order string) string {
	var query strings.Builder
	if order == "" {
		order = "created_at desc"
	}
	query.WriteString(` order by ` + order + ` `)
	return query.String()
}

func createPaginationClause(pagin *common.Pagination) string {
	var query strings.Builder
	query.WriteString(`limit ` + strconv.Itoa(pagin.PSize))
	query.WriteString(` offset ` + strconv.Itoa((pagin.Page-1)*pagin.PSize))
	return query.String()
}
