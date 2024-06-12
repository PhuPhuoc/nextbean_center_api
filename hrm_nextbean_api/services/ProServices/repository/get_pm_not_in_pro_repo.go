package repository

import (
	query "github.com/PhuPhuoc/hrm_nextbean_api/rawsql/project_query"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/model"
)

func (store *projectStore) GetPMNotInPro(pro_id string) ([]model.PM, error) {
	if err_pro_exist := checkProjectIDExist(store, pro_id); err_pro_exist != nil {
		return nil, err_pro_exist
	}
	rawsql := query.QueryGetAllPMNotInProject()
	data := []model.PM{}

	rows, err_query := store.db.Query(rawsql, pro_id)
	if err_query != nil {
		return data, err_query
	}
	defer rows.Close()

	for rows.Next() {
		pm := new(model.PM)
		if err_scan := rows.Scan(&pm.Id, &pm.UserName, &pm.Email); err_scan != nil {
			return data, err_scan
		}
		data = append(data, *pm)
	}
	return data, nil
}
