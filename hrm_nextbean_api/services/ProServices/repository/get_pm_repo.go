package repository

import (
	"github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/model"
)

func (store *projectStore) GetPM(pro_id string) ([]model.PM, error) {
	if err_pro_exist := checkProjectIDExists(store, pro_id); err_pro_exist != nil {
		return nil, err_pro_exist
	}
	rawsql := `select acc.id, acc.user_name, acc.email from project_manager pm join account acc on pm.account_id=acc.id where pm.project_id=? and acc.deleted_at is null`
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
