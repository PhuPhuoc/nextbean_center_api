package repository

import (
	"github.com/PhuPhuoc/hrm_nextbean_api/common"
	query "github.com/PhuPhuoc/hrm_nextbean_api/rawsql/intern_query"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/InternSevices/model"
)

func (store *InternStore) GetIntern(pagin *common.Pagination, filter *model.InternFilter) ([]model.Intern, error) {
	var total_record int64 = 0
	data := []model.Intern{}
	rawsql, param := query.QueryGetIntern(pagin, filter)

	rows, err_query := store.db.Query(rawsql, param...)
	if err_query != nil {
		return data, err_query
	}
	defer rows.Close()

	for rows.Next() {
		in := new(model.Intern)
		if err_scan := rows.Scan(&in.AccountID, &in.UserName, &in.Email, &in.StudentCode, &in.Ojt_semester, &in.Avatar, &in.Gender, &in.DateOfBirth, &in.PhoneNumber, &in.Address, &total_record); err_scan != nil {
			return data, err_scan
		}
		data = append(data, *in)
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
