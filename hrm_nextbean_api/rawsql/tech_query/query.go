package techquery

import (
	"strconv"
	"strings"

	"github.com/PhuPhuoc/hrm_nextbean_api/common"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/TechnicalServices/model"
)

func QueryCreateTechnical() string {
	return `insert into technical(technical_skill, created_at) values (?,?)`
}

func QueryGetTech(pagin *common.Pagination, filter *model.FilterTechnical) (string, []interface{}) {
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

func createConditionClause(filter *model.FilterTechnical) (string, []interface{}) {
	param := []interface{}{}
	var query strings.Builder
	query.WriteString(` where `)

	if filter.Id != 0 {
		query.WriteString(`id = ? and `)
		param = append(param, filter.Id)
	}
	if filter.TechnicalSkill != "" {
		query.WriteString(`technical_skill like ? and `)
		p := `%` + filter.TechnicalSkill + `%`
		param = append(param, p)

	}

	query.WriteString(`deleted_at is null `)
	return query.String(), param
}

func createCTEClause(condition_clause string) string {
	var query strings.Builder
	query.WriteString(`with cte as ( select count(*) as total_record from technical` + condition_clause + `)`)
	return query.String()
}

func createSelectClause(condition_clause string) string {
	var query strings.Builder
	query.WriteString(` select id, technical_skill, cte.total_record from technical`)
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
