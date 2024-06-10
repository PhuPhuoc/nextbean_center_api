package repository

import (
	"github.com/PhuPhuoc/hrm_nextbean_api/common"
	query "github.com/PhuPhuoc/hrm_nextbean_api/rawsql/ojt_query"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/OJTServices/model"
)

func (store *ojtStore) GetOJT(pagin *common.Pagination, filter *model.FilterOJT) ([]model.OJT, error) {
	var total_record int64 = 0
	data := []model.OJT{}
	rawsql, param := query.QueryGetOJT(pagin, filter)

	rows, err_query := store.db.Query(rawsql, param...)
	if err_query != nil {
		return data, err_query
	}
	defer rows.Close()

	for rows.Next() {
		ojt := new(model.OJT)
		if err_scan := rows.Scan(&ojt.Id, &ojt.Semester, &ojt.University, &ojt.StartAt, &ojt.EndAt, &total_record); err_scan != nil {
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
