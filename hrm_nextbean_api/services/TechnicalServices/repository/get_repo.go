package repository

import (
	"github.com/PhuPhuoc/hrm_nextbean_api/common"
	query "github.com/PhuPhuoc/hrm_nextbean_api/rawsql/tech_query"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/TechnicalServices/model"
)

func (store *techStore) GetTech(pagin *common.Pagination, filter *model.FilterTechnical) ([]model.Technical, error) {
	var total_record int64 = 0
	data := []model.Technical{}
	rawsql, param := query.QueryGetTech(pagin, filter)

	rows, err_query := store.db.Query(rawsql, param...)
	if err_query != nil {
		return data, err_query
	}
	defer rows.Close()

	for rows.Next() {
		ojt := new(model.Technical)
		if err_scan := rows.Scan(&ojt.Id, &ojt.TechnicalSkill, &total_record); err_scan != nil {
			return data, err_scan
		}
		data = append(data, *ojt)
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
