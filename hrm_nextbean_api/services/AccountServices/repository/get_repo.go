package repository

import (
	"github.com/PhuPhuoc/hrm_nextbean_api/common"
	query "github.com/PhuPhuoc/hrm_nextbean_api/rawsql/account_query"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/AccountServices/model"
)

func (store *AccountStore) GetAccount(pagin *common.Pagination, filter *model.AccountFilter) ([]model.Account, error) {
	var total_record int64 = 0
	data := []model.Account{}
	rawsql, param := query.QueryGetAccount(pagin, filter)

	rows, err_query := store.db.Query(rawsql, param...)
	if err_query != nil {
		return data, err_query
	}
	defer rows.Close()

	for rows.Next() {
		acc := new(model.Account)
		if err_scan := rows.Scan(&acc.Id, &acc.UserName, &acc.Email, &acc.Role, &acc.CreatedAt, &total_record); err_scan != nil {
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
