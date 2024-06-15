package repository

import (
	"fmt"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/AccountServices/model"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
)

func (store *accountStore) AccountLogin(email, password string, account *model.Account) error {
	rawqsl := `select id, user_name, email, password, role from  account where email = ? and deleted_at is null`
	rows, err := store.db.Query(rawqsl, email)
	if err != nil {
		return fmt.Errorf("DB error (query failed): %v", err)
	}
	count := 0
	for rows.Next() {
		count++
		if count > 1 {
			return fmt.Errorf("DB error (there's more than 1 address: %v in the db)", email)
		}
		err_scan := rows.Scan(
			&account.Id,
			&account.UserName,
			&account.Email,
			&account.Password,
			&account.Role)
		if err_scan != nil {
			return fmt.Errorf("DB error (scan failed): %v", err)
		}
	}
	if count == 0 {
		return fmt.Errorf("email: %v not exists", email)
	}

	if !utils.CompareHash(password, account.Password) {
		return fmt.Errorf("wrong pwd")
	}
	return nil
}
