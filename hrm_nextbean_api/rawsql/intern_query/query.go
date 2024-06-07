package internquery

import (
	"strconv"
	"strings"

	"github.com/PhuPhuoc/hrm_nextbean_api/common"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/InternSevices/model"
)

var table = `intern`
var table_name = ` ` + table + ` `

func QueryCreateNewAccount() string {
	return `insert into` + table_name + `(id, user_name, email, password, role, created_at) values (?,?,?,?,?,?)`
}

func QueryCreateNewIntern() string {
	return `insert into` + table_name + `(id,account_id,ojt_id,avatar,gender,dateofbirth,phone_number,address) values (?,?,?,?,?,?,?,?)`
}

func QueryCheckDulicateDataInIntern() string {
	start := `select`
	part1 := ` case when exists (select 1 from account where email = ?) then 'email' end as email_exists,`
	part2 := ` case when exists (select 1 from intern where id = ?) then 'id' end as studentcode_exists,`
	part3 := ` case when exists (select 1 from intern where phone_number = ?) then 'phone_number' end as phonenumber_exists`
	end := `from DUAL;`
	return start + ` ` + part1 + ` ` + part2 + ` ` + part3 + ` ` + end
}

func QueryGetIntern(pagin *common.Pagination, filter *model.InternFilter) (string, []interface{}) {
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

func createCTEClause(condition_clause string) string {
	var query strings.Builder
	join := `from intern i join account acc on i.account_id = acc.id join ojt o on i.ojt_id = o.id`
	query.WriteString(`with cte as ( select count(*) as total_record ` + join + ` ` + condition_clause + `)`)
	return query.String()
}

func createSelectClause(condition_clause string) string {
	var query strings.Builder
	join := `from intern i join account acc on i.account_id = acc.id join ojt o on i.ojt_id = o.id join cte`
	query.WriteString(` select acc.id, acc.user_name, acc.email, i.id, o.semester, i.avatar, i.gender, i.date_of_birth, i.phone_number, i.address , cte.total_record ` + join)
	query.WriteString(condition_clause)
	return query.String()
}

func createOrderByClause(order string) string {
	var query strings.Builder
	if order == "" {
		order = "acc.created_at desc"
	}
	query.WriteString(`order by ` + order + ` `)
	return query.String()
}

func createPaginationClause(pagin *common.Pagination) string {
	var query strings.Builder
	query.WriteString(`limit ` + strconv.Itoa(pagin.PSize))
	query.WriteString(` offset ` + strconv.Itoa((pagin.Page-1)*pagin.PSize))
	return query.String()
}

func createConditionClause(filter *model.InternFilter) (string, []interface{}) {
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
		query.WriteString(`i.id like ? and `)
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
