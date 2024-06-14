package repository

import (
	query "github.com/PhuPhuoc/hrm_nextbean_api/rawsql/project_query"

	"github.com/PhuPhuoc/hrm_nextbean_api/common"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/model"
)

func (store *projectStore) GetProject(pagin *common.Pagination, filter *model.ProjectFilter) ([]model.Project, error) {
	var total_record int64 = 0
	data := []model.Project{}
	rawsql, param := query.QueryGetProject(pagin, filter)

	rows, err_query := store.db.Query(rawsql, param...)
	if err_query != nil {
		return data, err_query
	}
	defer rows.Close()

	for rows.Next() {
		acc := new(model.Project)
		if err_scan := rows.Scan(&acc.Id, &acc.Name, &acc.Status, &acc.Description, &acc.Duration, &acc.StartDate, &total_record); err_scan != nil {
			return data, err_scan
		}
		data = append(data, *acc)
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
