package repository

import (
	"fmt"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/model"
)

func (store *projectStore) MapPM(proid string, info *model.MapProPM) error {
	if err_check_pro_id := checkProjectIDExists(store, proid); err_check_pro_id != nil {
		return err_check_pro_id
	}

	for _, pm_id := range info.ListManagerId {
		if err_check_pro_id := checkPMIDExists(store, pm_id); err_check_pro_id != nil {
			return err_check_pro_id
		}
	}

	// todo: start transaction
	tx, err := store.db.Begin()
	if err != nil {
		return fmt.Errorf("error in MapPM transaction: %v", err)
	}

	deleteQuery := `delete from project_manager where project_id = ?`
	if _, err := tx.Exec(deleteQuery, proid); err != nil {
		tx.Rollback()
		return fmt.Errorf("error in MapPM transaction-delete mapping: %v", err)
	}

	if len(info.ListManagerId) == 0 {
		return nil
	}

	values := ""
	for i := range info.ListManagerId {
		if i > 0 {
			values += ","
		}
		values += fmt.Sprintf("('%s','%s')", proid, info.ListManagerId[i])
	}

	updateQuery := fmt.Sprintf("insert into project_manager (project_id, account_id) VALUES %s", values)
	if _, err := tx.Exec(updateQuery); err != nil {
		tx.Rollback()
		return fmt.Errorf("error in MapPM transaction-update mapping: %v", err)
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return fmt.Errorf("error in MapPM transaction-commit: %v", err)
	}

	return nil
}
