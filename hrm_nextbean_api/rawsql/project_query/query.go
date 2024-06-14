package projectquery

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PhuPhuoc/hrm_nextbean_api/common"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/model"
)

func QueryCreateProject() string {
	return `insert into project(id, name, status, description, start_date, duration, created_at) values (?,?,?,?,?,?,?)`
}

func QueryCheckProjectIDExist() string {
	return `select exists(select 1 from project where id = ? and deleted_at is null)`
}

func QueryCheckPMIDExist() string {
	return `select exists(select 1 from account where id = ? and role='pm' and deleted_at is null)`
}

func QueryCheckMemIDExist() string {
	return `select exists(select 1 from intern i join account acc on i.account_id=acc.id where i.id = ? and acc.deleted_at is null)`
}

func QueryUpdateProject() string {
	return `update project set name=?, status=?, description=?, start_date=?, duration=? where id=?`
}

// todo: map PM
func QueryDeleteMapProjectPM() string {
	return `DELETE FROM project_manager WHERE project_id = ?`
}

func QueryMapProjectPM(values string) string {
	return fmt.Sprintf("INSERT INTO project_manager (project_id, account_id) VALUES %s", values)
}

// todo: map project-member
func QueryCheckMemberInProjectSatusLeave() string { // 2
	return `select exists(select 1 from project_intern where project_id = ? and intern_id = ? and status = 'leave')`
}

func QueryCheckMemberInProjectNotExist() string { // 1
	return `select exists(select 1 from project_intern where project_id = ? and intern_id = ?)`
}

// todo: for delete
func QueryCheckMemTaskBeforeDelete() string {
	return `select exists(select 1 from task where project_id=? and assigned_to=? and status!='done')`
}

func QueryDeleteMemberInProject() string {
	return `update project_intern set leave_at=?, status='leave' where project_id=? and intern_id=?`
}

// todo: for add
func QueryAddMemberToProject() string {
	return `insert into project_intern(project_id, intern_id, join_at, status) values (?,?,?,?)`
}

func QueryReJoinProject() string {
	return `update project_intern set leave_at=null, status='inprogress' where project_id=? and intern_id=?`
}

// todo: get all member in project
func QueryGetAllMemberInProject() string {
	fields := `i.id, acc.user_name, i.student_code, i.avatar, o.semester, o.university`
	sel := `select ` + fields + ` , GROUP_CONCAT(tech.technical_skill SEPARATOR ', ') AS technical_skills `
	from := `from intern i`
	join1 := ` join account acc on i.account_id=acc.id `
	join2 := ` join project_intern proi on proi.intern_id = i.id `
	join3 := ` join intern_skill ins on ins.intern_id = i.id `
	join4 := ` join technical tech on tech.id=ins.technical_id `
	join5 := ` join ojt o on o.id=i.ojt_id `
	where := `where proi.project_id = ? and acc.deleted_at is null`
	groupby := ` group by ` + fields
	return sel + from + join1 + join2 + join3 + join4 + join5 + where + groupby
}

func QueryGetAllMemberNotInProject() string {
	fields := `i.id, acc.user_name, i.student_code, i.avatar`
	sel := `select ` + fields + ` , GROUP_CONCAT(tech.technical_skill SEPARATOR ', ') AS technical_skills `
	from := `from intern i`
	join1 := ` join account acc on i.account_id=acc.id `
	join2 := ` left join project_intern proi on proi.intern_id = i.id `
	join3 := ` left join intern_skill ins on ins.intern_id = i.id `
	join4 := ` left join technical tech on tech.id=ins.technical_id `
	where := `where i.id not in (select intern_id from project_intern where project_id=?) and acc.deleted_at is null`
	groupby := ` group by ` + fields
	return sel + from + join1 + join2 + join3 + join4 + where + groupby
}

// todo: get all pm in project
func QueryGetAllPMInProject() string {
	return `select acc.id, acc.user_name, acc.email from project_manager pm join account acc on pm.account_id=acc.id where pm.project_id=? and acc.deleted_at is null`
}

func QueryGetAllPMNotInProject() string {
	return `select acc.id, acc.user_name, acc.email from account acc left join project_manager pm on pm.account_id=acc.id where acc.id not in (select account_id from project_manager where project_id=?) and acc.role = 'pm' and acc.deleted_at is null`
}

// todo: get
func QueryGetProject(pagin *common.Pagination, filter *model.ProjectFilter) (string, []interface{}) {
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

func createConditionClause(filter *model.ProjectFilter) (string, []interface{}) {
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

func createCTEClause(condition_clause string) string {
	var query strings.Builder
	query.WriteString(`with cte as ( select count(*) as total_record from project` + condition_clause + `)`)
	return query.String()
}

func createSelectClause(condition_clause string) string {
	var query strings.Builder
	query.WriteString(` select id, name, status, description, duration, start_date, cte.total_record from project join cte ` + condition_clause)
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

func QueryGetMemberOutsideProject(proid string, pagin *common.Pagination, filter *model.MemberFilter) (string, []interface{}) {
	var query strings.Builder
	where, param := queryWhere(proid, filter)
	cte := cte(where)
	main := sel(where)
	pag := createPaginationClause(pagin)

	// double param because in this code has 2 part of where clause ( 1 in cte - other in main select )
	doubledParams := make([]interface{}, len(param)*2)
	copy(doubledParams, param)
	copy(doubledParams[len(param):], param)

	query.WriteString(cte + main + pag)
	return query.String(), doubledParams
}

func queryWhere(proid string, filter *model.MemberFilter) (string, []interface{}) {
	param := []interface{}{}
	var query strings.Builder
	query.WriteString(` where i.id not in (select intern_id from project_intern where project_id=?) and `)
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
		query.WriteString(`o.semester = ? and `)
		param = append(param, filter.Semester)
	}

	if filter.University != "" {
		query.WriteString(`o.university = ? and `)
		param = append(param, filter.University)
	}

	query.WriteString(`acc.deleted_at is null `)
	return query.String(), param
}

func cte(where string) string {
	var query strings.Builder
	query.WriteString(`with cte as (select count(*) AS total_record from intern i`)
	query.WriteString(` join account acc on i.account_id=acc.id `)
	query.WriteString(` left join project_intern proi on proi.intern_id = i.id `)
	query.WriteString(` left join ojt o on o.id=i.ojt_id `)
	query.WriteString(where)
	query.WriteString(`)`)
	return query.String()
}

func sel(where string) string {
	var query strings.Builder
	query.WriteString(`select i.id, acc.user_name, i.student_code, i.avatar, o.semester, o.university, GROUP_CONCAT(tech.technical_skill SEPARATOR ', ') AS technical_skills, cte.total_record`)
	query.WriteString(` from intern i `)
	query.WriteString(` join account acc on i.account_id=acc.id `)
	query.WriteString(` left join project_intern proi on proi.intern_id = i.id `)
	query.WriteString(` left join intern_skill ins on ins.intern_id = i.id `)
	query.WriteString(` left join technical tech on tech.id=ins.technical_id `)
	query.WriteString(` left join ojt o on o.id=i.ojt_id `)
	query.WriteString(` join cte `)
	query.WriteString(where)
	query.WriteString(` group by i.id, acc.user_name, i.student_code, i.avatar, o.semester, o.university, cte.total_record `)
	query.WriteString(` order by acc.created_at desc `)
	return query.String()
}
