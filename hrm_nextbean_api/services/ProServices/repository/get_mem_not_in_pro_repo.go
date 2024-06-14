package repository

import (
	"database/sql"

	"github.com/PhuPhuoc/hrm_nextbean_api/common"
	query "github.com/PhuPhuoc/hrm_nextbean_api/rawsql/project_query"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/model"
)

func (store *projectStore) GetMemNotInPro(pro_id string, pagin *common.Pagination, filter *model.MemberFilter) ([]model.Member, error) {
	if err_pro_exist := checkProjectIDExist(store, pro_id); err_pro_exist != nil {
		return nil, err_pro_exist
	}

	var total_record int64 = 0
	data := []model.Member{}
	rawsql, param := query.QueryGetMemberOutsideProject(pro_id, pagin, filter)

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

		if err_scan := rows.Scan(&mem.Id, &mem.UserName, &mem.StudentCode, &mem.Avatar, &ojt_semester, &ojt_uni, &technicalSkills, &total_record); err_scan != nil {
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
