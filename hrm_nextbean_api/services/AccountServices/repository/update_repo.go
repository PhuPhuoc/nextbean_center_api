package repository

import (
	"fmt"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/AccountServices/model"
)

func (store *accountStore) UpdateAccount(accountID string, acc_update_info *model.UpdateAccountInfo) error {
	// todo 1: check id and email valid before updating
	if err_check_id_exist := store.checkAccountIdExists(accountID); err_check_id_exist != nil {
		return err_check_id_exist
	}

	if err_check_email_exist := store.checkEmailExistWithID(acc_update_info.Email, accountID); err_check_email_exist != nil {
		return err_check_email_exist
	}
	// todo 2: update account'info by account'id
	rawsql := `update account set user_name = ?, email = ?, role = ? where id = ?`
	result, err := store.db.Exec(rawsql, acc_update_info.UserName, acc_update_info.Email, acc_update_info.Role, accountID)
	if err != nil {
		return fmt.Errorf("error when UpdateAccount in store: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error when UpdateAccount in store (check affect): %v", err)
	}
	if rowsAffected == 0 {
		return nil
	}
	return nil
}
