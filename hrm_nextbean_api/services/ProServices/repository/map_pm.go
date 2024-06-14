package repository

import (
	"fmt"

	query "github.com/PhuPhuoc/hrm_nextbean_api/rawsql/project_query"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/model"
)

func (store *projectStore) MapPM(proid string, info *model.MapProPM) error {
	if err_check_pro_id := checkProjectIDExist(store, proid); err_check_pro_id != nil {
		return err_check_pro_id
	}

	for _, pm_id := range info.ListManagerId {
		if err_check_pro_id := checkPMIDExist(store, pm_id); err_check_pro_id != nil {
			return err_check_pro_id
		}
	}

	// todo: start transaction
	tx, err := store.db.Begin()
	if err != nil {
		return fmt.Errorf("error when MapPM (start transaction) in store: %v", err)
	}

	deleteQuery := query.QueryDeleteMapProjectPM()
	if _, err := tx.Exec(deleteQuery, proid); err != nil {
		tx.Rollback()
		return fmt.Errorf("error when MapPM (delete mapping transaction) in store: %v", err)
	}

	values := ""
	for i := range info.ListManagerId {
		if i > 0 {
			values += ","
		}
		values += fmt.Sprintf("('%s','%s')", proid, info.ListManagerId[i])
	}

	updateQuery := query.QueryMapProjectPM(values)
	if _, err := tx.Exec(updateQuery); err != nil {
		tx.Rollback()
		return fmt.Errorf("error when MapPM (update mapping transaction) in store: %v", err)
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return fmt.Errorf("error when MapPM (commit transaction) in store: %v", err)
	}

	return nil
}
