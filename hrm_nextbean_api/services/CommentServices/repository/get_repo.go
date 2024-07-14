package repository

import (
	"strconv"
	"strings"

	"github.com/PhuPhuoc/hrm_nextbean_api/common"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/CommentServices/model"
)

func (store *commentStore) GetComment(ownerid string, pagin *common.Pagination, filter *model.CommentFilter) ([]model.Comment, error) {
	var total_record int64 = 0
	data := []model.Comment{}
	rawsql, param := rawSqlGetComment(pagin, filter)

	rows, err_query := store.db.Query(rawsql, param...)
	if err_query != nil {
		return data, err_query
	}
	defer rows.Close()

	for rows.Next() {
		com := new(model.Comment)
		var accid *string
		if err_scan := rows.Scan(&com.Id, &accid, &com.UserName, &com.Avatar, &com.Type, &com.Content, &com.CreatedAt, &total_record); err_scan != nil {
			return data, err_scan
		}
		if *accid == ownerid {
			com.IsOwner = true
		} else {
			com.IsOwner = false
		}
		data = append(data, *com)
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

func rawSqlGetComment(pagin *common.Pagination, filter *model.CommentFilter) (string, []interface{}) {
	where, param := whereClause(filter)
	main := mainClause(where, pagin)

	// double param because in this code has 2 part of where clause ( 1 in cte - other in main select )
	doubledParams := make([]interface{}, len(param)*2)
	copy(doubledParams, param)
	copy(doubledParams[len(param):], param)

	return main, doubledParams
}

func mainClause(where string, pagin *common.Pagination) string {
	order := `r.created_at DESC`
	join := ` from report r join account a on r.account_id=a.id left join intern i on i.account_id=a.id `
	var query strings.Builder
	query.WriteString(`with cte as ( select count(*) as total_record` + join + where + `) `)
	query.WriteString(`select r.id, a.id, a.user_name, i.avatar, r.type, r.content, r.created_at, cte.total_record`)
	query.WriteString(join + ` join cte` + where)
	query.WriteString(` order by ` + order)
	query.WriteString(` limit ` + strconv.Itoa(pagin.PSize))
	query.WriteString(` offset ` + strconv.Itoa((pagin.Page-1)*pagin.PSize))
	return query.String()
}

func whereClause(filter *model.CommentFilter) (string, []interface{}) {
	param := []interface{}{}
	var query strings.Builder
	query.WriteString(` where `)

	if filter.Type != "" {
		query.WriteString(`r.type = ? and `)
		param = append(param, filter.Type)
	}

	query.WriteString(`r.deleted_at is null `)
	return query.String(), param
}
