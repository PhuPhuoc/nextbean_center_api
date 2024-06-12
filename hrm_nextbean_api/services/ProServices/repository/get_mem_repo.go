package repository

import (
	query "github.com/PhuPhuoc/hrm_nextbean_api/rawsql/project_query"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/model"
)

func (store *projectStore) GetMem(pro_id string) ([]model.Member, error) {
	if err_pro_exist := checkProjectIDExist(store, pro_id); err_pro_exist != nil {
		return nil, err_pro_exist
	}
	rawsql := query.QueryGetAllMemberInProject()
	data := []model.Member{}

	rows, err_query := store.db.Query(rawsql, pro_id)
	if err_query != nil {
		return data, err_query
	}
	defer rows.Close()

	for rows.Next() {
		mem := new(model.Member)
		if err_scan := rows.Scan(&mem.Id, &mem.UserName, &mem.StudentCode, &mem.Avatar, &mem.TechnicalSkills); err_scan != nil {
			return data, err_scan
		}
		data = append(data, *mem)
	}
	return data, nil
}
