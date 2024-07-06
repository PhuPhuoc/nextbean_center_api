package repository

import (
	"database/sql"
	"strconv"
	"strings"

	"github.com/PhuPhuoc/hrm_nextbean_api/common"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/model"
)

func (store *projectStore) GetMem(pro_id string, pagin *common.Pagination, filter *model.MemberFilter) ([]model.Member, error) {
	if err_pro_exist := checkProjectIDExists(store, pro_id); err_pro_exist != nil {
		return nil, err_pro_exist
	}

	var total_record int64 = 0
	data := []model.Member{}
	rawsql, param := rawSqlGetMemberInProject(pro_id, pagin, filter)

	rows, err_query := store.db.Query(rawsql, param...)
	if err_query != nil {
		return data, err_query
	}
	defer rows.Close()

	for rows.Next() {
		mem := new(model.Member)
		var technicalSkills sql.NullString
		var ojt_semester sql.NullString
		var ojt_uni sql.NullString

		if err_scan := rows.Scan(&mem.Id, &mem.UserName, &mem.StudentCode, &mem.Avatar, &ojt_semester, &ojt_uni, &mem.Status, &technicalSkills, &total_record); err_scan != nil {
			return data, err_scan
		}

		if technicalSkills.Valid {
			mem.TechnicalSkills = technicalSkills.String
		} else {
			mem.TechnicalSkills = ""
		}

		if ojt_semester.Valid || ojt_uni.Valid {
			mem.OjtSemesterUniversity = ojt_semester.String + " - " + ojt_uni.String
		} else {
			mem.OjtSemesterUniversity = ""
		}

		data = append(data, *mem)
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

func rawSqlGetMemberInProject(proid string, pagin *common.Pagination, filter *model.MemberFilter) (string, []interface{}) {
	var query strings.Builder
	where, param := queryWhere_in(proid, filter)
	cte := cte_in(where)
	main := sel_in(where, pagin)
	// double param because in this code has 2 part of where clause ( 1 in cte - other in main select )
	doubledParams := make([]interface{}, len(param)*2)
	copy(doubledParams, param)
	copy(doubledParams[len(param):], param)

	query.WriteString(cte + main)
	return query.String(), doubledParams
}

func cte_in(where string) string {
	var query strings.Builder
	query.WriteString(`with cte as (select count(*) AS total_record from intern i`)
	query.WriteString(` join account acc on i.account_id=acc.id `)
	query.WriteString(` left join project_intern proi on proi.intern_id = i.id `)
	query.WriteString(` left join ojt o on o.id=i.ojt_id `)
	query.WriteString(where)
	query.WriteString(`)`)
	return query.String()
}

func sel_in(where string, pagin *common.Pagination) string {
	var query strings.Builder
	query.WriteString(`select i.id, acc.user_name, i.student_code, i.avatar, o.semester, o.university, proi.status, GROUP_CONCAT(tech.technical_skill SEPARATOR ', ') AS technical_skills, cte.total_record`)
	query.WriteString(` from intern i `)
	query.WriteString(` join account acc on i.account_id=acc.id `)
	query.WriteString(` join project_intern proi on proi.intern_id = i.id `)
	query.WriteString(` join intern_skill ins on ins.intern_id = i.id `)
	query.WriteString(` join technical tech on tech.id=ins.technical_id `)
	query.WriteString(` join ojt o on o.id=i.ojt_id `)
	query.WriteString(` join cte `)
	query.WriteString(where)
	query.WriteString(` group by i.id, acc.user_name, i.student_code, i.avatar, o.semester, o.university, proi.status, cte.total_record `)
	query.WriteString(` order by acc.created_at desc `)
	query.WriteString(`limit ` + strconv.Itoa(pagin.PSize))
	query.WriteString(` offset ` + strconv.Itoa((pagin.Page-1)*pagin.PSize))
	return query.String()
}

func queryWhere_in(proid string, filter *model.MemberFilter) (string, []interface{}) {
	param := []interface{}{}
	var query strings.Builder
	query.WriteString(` where proi.project_id = ? and `)
	param = append(param, proid)

	if filter.UserName != "" {
		query.WriteString(`acc.user_name like ? and `)
		p := `%` + filter.UserName + `%`
		param = append(param, p)
	}

	if filter.StudentCode != "" {
		query.WriteString(`i.student_code like ? and `)
		p := `%` + filter.StudentCode + `%`
		param = append(param, p)
	}

	if filter.Semester != "" {
		query.WriteString(`o.semester like ? and `)
		p := `%` + filter.Semester + `%`
		param = append(param, p)
	}

	if filter.University != "" {
		query.WriteString(`o.university like ? and `)
		p := `%` + filter.University + `%`
		param = append(param, p)
	}

	if filter.Status != "" {
		query.WriteString(`proi.status = ? and `)
		param = append(param, filter.Status)
	}

	query.WriteString(`acc.deleted_at is null `)
	return query.String(), param
}
