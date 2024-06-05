package repository

import (
	"fmt"
	"strings"

	query "github.com/PhuPhuoc/hrm_nextbean_api/rawsql/account_query"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
)

func (store *AccountStore) DeleteAccount(id string) error {

	if err_check_id_exist := store.checkIdExist(id); err_check_id_exist != nil {
		if strings.Contains(err_check_id_exist.Error(), "id not exist") {
			return fmt.Errorf("account'ID not exists")
		}
		return fmt.Errorf("error when DeleteAccount(checkIdExist) in store: %v", err_check_id_exist)
	}

	rawsql := query.QueryDeleteAccount()
	result, err := store.db.Exec(rawsql, utils.CreateDateTimeCurrentFormated(), id)
	if err != nil {
		return fmt.Errorf("error when DeleteAccount in store: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error when DeleteAccount in store (check affect): %v", err)
	}
	if rowsAffected == 1 {
		return nil // created sucessfully
	} else {
		return fmt.Errorf("error when DeleteAccount in store (No user deleted): %v", err)
	}
}
