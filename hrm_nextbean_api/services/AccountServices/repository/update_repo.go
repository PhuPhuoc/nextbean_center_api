package repository

import (
	"fmt"
	"strings"

	query "github.com/PhuPhuoc/hrm_nextbean_api/rawsql/account_query"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/AccountServices/model"
)

func (store *accountStore) UpdateAccount(accountID string, acc_update_info *model.UpdateAccountInfo) error {

	if err_check_id_exist := store.checkIdExist(accountID); err_check_id_exist != nil {
		if strings.Contains(err_check_id_exist.Error(), "not_exist_id") {
			return fmt.Errorf("account'ID dose not exists")
		}
		return fmt.Errorf("error when UpdateAccount(checkIdExist) in store: %v", err_check_id_exist)
	}

	if err_check_email_exist := store.checkEmailExistWithID(acc_update_info.Email, accountID); err_check_email_exist != nil {
		if strings.Contains(err_check_email_exist.Error(), "duplicate_data_email") {
			return fmt.Errorf("email: %v already exists", acc_update_info.Email)
		}
		return fmt.Errorf("error when UpdateAccount(checkEmailExist) in store: %v", err_check_email_exist)
	}

	rawsql := query.QueryUpdateAccount()
	result, err := store.db.Exec(rawsql, acc_update_info.UserName, acc_update_info.Email, acc_update_info.Role, accountID)
	if err != nil {
		return fmt.Errorf("error when UpdateAccount in store: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error when UpdateAccount in store (check affect): %v", err)
	}
	if rowsAffected == 1 {
		return nil // created sucessfully
	} else {
		return fmt.Errorf("error when UpdateAccount in store (No user updated): %v", err)
	}
}
