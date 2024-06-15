package repository

import (
	"fmt"

	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
)

func (store *accountStore) DeleteAccount(id string) error {
	// todo 1: check account'id exists
	if err_check_id_exist := store.checkAccountIdExists(id); err_check_id_exist != nil {
		return err_check_id_exist
	}
	// todo 2: delete account by account'id
	rawsql := `update account set deleted_at = ? where id = ?`
	result, err := store.db.Exec(rawsql, utils.CreateDateTimeCurrentFormated(), id)
	if err != nil {
		return fmt.Errorf("error when DeleteAccount in store: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error when DeleteAccount in store (check affect): %v", err)
	}
	if rowsAffected == 1 {
		return nil // deleted sucessfully
	} else {
		return fmt.Errorf("error when DeleteAccount in store (No user deleted): %v", err)
	}
}
