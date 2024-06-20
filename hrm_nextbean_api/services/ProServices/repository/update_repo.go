package repository

import (
	"fmt"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/model"
)

func (store *projectStore) UpdateProject(proid string, info *model.UpdateProjectInfo) error {
	if err_check_id := checkProjectIDExists(store, proid); err_check_id != nil {
		return err_check_id
	}
	rawsql := `update project set name=?, status=?, description=?, start_date=?, duration=? where id=?`
	result, err := store.db.Exec(rawsql, info.Name, info.Status, info.Description, info.StartDate, info.Duration, proid)
	if err != nil {
		return fmt.Errorf("error when UpdateProject in store: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error when UpdateProject in store (check affect): %v", err)
	}
	if rowsAffected == 0 {
		return nil
	}
	return nil
}
