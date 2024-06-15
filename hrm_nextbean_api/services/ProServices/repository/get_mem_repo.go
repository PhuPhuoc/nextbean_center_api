package repository

import (
	"database/sql"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/model"
)

func (store *projectStore) GetMem(pro_id string) ([]model.Member, error) {
	if err_pro_exist := checkProjectIDExists(store, pro_id); err_pro_exist != nil {
		return nil, err_pro_exist
	}
	rawsql := rawSqlGetAllMemberInProject()
	data := []model.Member{}

	rows, err_query := store.db.Query(rawsql, pro_id)
	if err_query != nil {
		return data, err_query
	}
	defer rows.Close()

	for rows.Next() {
		mem := new(model.Member)
		var technicalSkills sql.NullString
		var ojt_semester sql.NullString
		var ojt_uni sql.NullString

		if err_scan := rows.Scan(&mem.Id, &mem.UserName, &mem.StudentCode, &mem.Avatar, &ojt_semester, &ojt_uni, &technicalSkills); err_scan != nil {
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
	return data, nil
}

func rawSqlGetAllMemberInProject() string {
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
