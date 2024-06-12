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
	return `select exists(select 1 from project where id = ?)`
}

func QueryCheckPMIDExist() string {
	return `select exists(select 1 from account where id = ? and role='pm')`
}

func QueryCheckMemIDExist() string {
	return `select exists(select 1 from intern where id = ?)`
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
	return `update project_intern set leave_at=? and status='leave' where project_id=? intern_id=?`
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
	return ``
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
	query.WriteString(` select id, name, status, duration, start_date, cte.total_record from project join cte ` + condition_clause)
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
