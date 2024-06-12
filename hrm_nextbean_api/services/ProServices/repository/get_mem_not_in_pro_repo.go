package repository

import (
	"database/sql"

	query "github.com/PhuPhuoc/hrm_nextbean_api/rawsql/project_query"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/model"
)

func (store *projectStore) GetMemNotInPro(pro_id string) ([]model.Member, error) {
	if err_pro_exist := checkProjectIDExist(store, pro_id); err_pro_exist != nil {
		return nil, err_pro_exist
	}
	rawsql := query.QueryGetAllMemberNotInProject()
	data := []model.Member{}

	rows, err_query := store.db.Query(rawsql, pro_id)
	if err_query != nil {
		return data, err_query
	}
	defer rows.Close()

	for rows.Next() {
		mem := new(model.Member)
		var technicalSkills sql.NullString

		if err_scan := rows.Scan(&mem.Id, &mem.UserName, &mem.StudentCode, &mem.Avatar, &technicalSkills); err_scan != nil {
			return data, err_scan
		}

		if technicalSkills.Valid {
			mem.TechnicalSkills = technicalSkills.String
		} else {
			mem.TechnicalSkills = ""
		}

		data = append(data, *mem)
	}
	return data, nil
}
