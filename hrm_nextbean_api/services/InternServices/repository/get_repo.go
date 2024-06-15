package repository

import (
	"strconv"
	"strings"

	"github.com/PhuPhuoc/hrm_nextbean_api/common"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/InternServices/model"
)

func (store *internStore) GetIntern(pagin *common.Pagination, filter *model.InternFilter) ([]model.Intern, error) {
	var total_record int64 = 0
	data := []model.Intern{}
	rawsql, param := rawSqlGetIntern(pagin, filter)

	rows, err_query := store.db.Query(rawsql, param...)
	if err_query != nil {
		return data, err_query
	}
	defer rows.Close()

	for rows.Next() {
		in := new(model.Intern)
		if err_scan := rows.Scan(&in.AccountID, &in.InternID, &in.UserName, &in.Email, &in.StudentCode, &in.Ojt_semester, &in.Avatar, &in.Gender, &in.DateOfBirth, &in.PhoneNumber, &in.Address, &total_record); err_scan != nil {
			return data, err_scan
		}
		data = append(data, *in)
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

func rawSqlGetIntern(pagin *common.Pagination, filter *model.InternFilter) (string, []interface{}) {
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
		order = "acc.created_at desc"
	}
	var query strings.Builder
	join := `from intern i join account acc on i.account_id = acc.id join ojt o on i.ojt_id = o.id`
	query.WriteString(`with cte as ( select count(*) as total_record ` + join + ` ` + where + `)`)
	query.WriteString(` select acc.id, i.id, acc.user_name, acc.email, i.student_code, o.semester, i.avatar, i.gender, i.date_of_birth, i.phone_number, i.address , cte.total_record ` + join)
	query.WriteString(` join cte`)
	query.WriteString(where)
	query.WriteString(`order by ` + order)
	query.WriteString(` limit ` + strconv.Itoa(pagin.PSize))
	query.WriteString(` offset ` + strconv.Itoa((pagin.Page-1)*pagin.PSize))
	return query.String()
}

func whereClause(filter *model.InternFilter) (string, []interface{}) {
	param := []interface{}{}
	var query strings.Builder
	query.WriteString(` where `)

	if filter.AccountID != "" {
		query.WriteString(`acc.id = ? and `)
		param = append(param, filter.AccountID)
	}

	if filter.UserName != "" {
		query.WriteString(`acc.user_name like ? and `)
		p := `%` + filter.UserName + `%`
		param = append(param, p)

	}
	if filter.Email != "" {
		query.WriteString(`acc.email like ? and `)
		p := `%` + filter.Email + `%`
		param = append(param, p)
	}

	if filter.OJT_Semester != "" {
		query.WriteString(`o.semester like ? and `)
		p := `%` + filter.OJT_Semester + `%`
		param = append(param, p)
	}

	if filter.StudentCode != "" {
		query.WriteString(`i.student_code like ? and `)
		p := `%` + filter.StudentCode + `%`
		param = append(param, p)
	}

	if filter.Gender != "" {
		query.WriteString(`i.gender = ? and `)
		param = append(param, filter.Gender)
	}

	if filter.PhoneNumber != "" {
		query.WriteString(`i.phone_number like ? and `)
		p := `%` + filter.PhoneNumber + `%`
		param = append(param, p)
	}

	if filter.Address != "" {
		query.WriteString(`i.address like ? and `)
		p := `%` + filter.Address + `%`
		param = append(param, p)
	}

	if filter.Dob_From != "" {
		query.WriteString(`i.date_of_birth > ? and `)
		param = append(param, filter.Dob_From)
	}

	if filter.Dob_To != "" {
		query.WriteString(`i.date_of_birth < ? and `)
		param = append(param, filter.Dob_To)
	}

	query.WriteString(`acc.deleted_at is null `)
	return query.String(), param
}
