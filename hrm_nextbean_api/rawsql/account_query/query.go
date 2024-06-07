package accountquery

import (
	"strconv"
	"strings"

	"github.com/PhuPhuoc/hrm_nextbean_api/common"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/AccountServices/model"
)

var table = `account`
var table_name = ` ` + table + ` `
var table_field = ` id, user_name, email, password, role `

func QueryCreateNewAccount() string {
	return `insert into` + table_name + `(id, user_name, email, password, role, created_at) values (?,?,?,?,?,?)`
}

func QueryDeleteAccount() string {
	return `update` + table_name + `set deleted_at = ? where id = ?`
}

func QueryUpdateAccount() string {
	return `update` + table_name + `set user_name = ?, email = ?, role = ? where id = ?`
}

func QueryCheckExistEmail() string {
	return `select exists(select 1 from` + table_name + `where email = ?)`
}

func QueryCheckExistEmailWithID() string {
	return `select exists(select 1 from` + table_name + `where email = ? and id != ?)`
}

func QueryIdExist() string {
	return `select exists(select 1 from` + table_name + `where id = ? and deleted_at is null)`
}

func QueryGetAccountByEmailForLogin() string {
	return `select` + table_field + `from` + table_name + `where email = ? and deleted_at is null`
}

func createConditionClause(filter *model.AccountFilter) (string, []interface{}) {
	param := []interface{}{}
	var query strings.Builder
	query.WriteString(` where `)

	if filter.Id != "" {
		query.WriteString(`id = ? and `)
		param = append(param, filter.Id)
	}
	if filter.UserName != "" {
		query.WriteString(`user_name like ? and `)
		p := `%` + filter.UserName + `%`
		param = append(param, p)

	}
	if filter.Email != "" {
		query.WriteString(`email like ? and `)
		p := `%` + filter.Email + `%`
		param = append(param, p)

	}
	if filter.Role != "" {
		query.WriteString(`role in (`)
		parts := strings.Split(filter.Role, "-")
		for i, part := range parts {
			if i > 0 {
				query.WriteString(`,`)
			}
			query.WriteString(`?`)
			param = append(param, part)
		}
		query.WriteString(`) and `)
	}

	if filter.CreatedAtFrom != "" {
		query.WriteString(`created_at > ? and `)
		param = append(param, filter.CreatedAtFrom)
	}

	if filter.CreatedAtTo != "" {
		query.WriteString(`created_at < ? and `)
		param = append(param, filter.CreatedAtTo)
	}

	query.WriteString(`deleted_at is null `)
	return query.String(), param
}

func createCTEClause(condition_clause string) string {
	var query strings.Builder
	query.WriteString(`with cte as ( select count(*) as total_record from ` + table + ` ` + condition_clause + `)`)
	return query.String()
}

func createSelectClause(condition_clause string) string {
	var query strings.Builder
	query.WriteString(` select id, user_name, email, role, created_at, cte.total_record` + ` from ` + table)
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

func QueryGetAccount(pagin *common.Pagination, filter *model.AccountFilter) (string, []interface{}) {
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
