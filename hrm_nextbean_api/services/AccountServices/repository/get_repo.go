package repository

import (
	"strconv"
	"strings"

	"github.com/PhuPhuoc/hrm_nextbean_api/common"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/AccountServices/model"
)

func (store *accountStore) GetAccount(pagin *common.Pagination, filter *model.AccountFilter) ([]model.Account, error) {
	var total_record int64 = 0
	data := []model.Account{}
	rawsql, param := queryGetAccount(pagin, filter)

	rows, err_query := store.db.Query(rawsql, param...)
	if err_query != nil {
		return data, err_query
	}
	defer rows.Close()

	for rows.Next() {
		acc := new(model.Account)
		if err_scan := rows.Scan(&acc.Id, &acc.UserName, &acc.Email, &acc.Role, &acc.CreatedAt, &total_record); err_scan != nil {
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

func queryGetAccount(pagin *common.Pagination, filter *model.AccountFilter) (string, []interface{}) {
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
		order = "created_at desc"
	}
	query.WriteString(`with cte as (select count(*) as total_record from account` + where + `) `)
	query.WriteString(`select id, user_name, email, role, created_at, cte.total_record from account join cte`)
	query.WriteString(where)
	query.WriteString(` order by ` + order)
	query.WriteString(` limit ` + strconv.Itoa(pagin.PSize))
	query.WriteString(` offset ` + strconv.Itoa((pagin.Page-1)*pagin.PSize))
	return query.String()
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
	} else {
		query.WriteString(`role in ('admin','manager','pm') and `)
	}

	if filter.CreatedAtFrom != "" {
		query.WriteString(`created_at > ? and `)
		param = append(param, filter.CreatedAtFrom)
	}

	if filter.CreatedAtTo != "" {
		query.WriteString(`created_at < ? and `)
		param = append(param, filter.CreatedAtTo)
	}

	query.WriteString(`deleted_at is null`)
	return query.String(), param
}
