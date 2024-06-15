package repository

import (
	"fmt"

	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
)

func (store *ojtStore) DeleteOJT(ojt_id string) error {
	id, err_exist := checkOjtIDExists(store, ojt_id)
	if err_exist != nil {
		return err_exist
	}

	rawsql := `update ojt set deleted_at = ? where id = ?`
	result, err := store.db.Exec(rawsql, utils.CreateDateTimeCurrentFormated(), id)
	if err != nil {
		return fmt.Errorf("error db in DeleteOJT: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error db in DeleteOJT (check affect): %v", err)
	}
	if rowsAffected == 1 {
		return nil // update sucessfully
	} else {
		return fmt.Errorf("error db in DeleteOJT (No user created): %v", err)
	}
}
