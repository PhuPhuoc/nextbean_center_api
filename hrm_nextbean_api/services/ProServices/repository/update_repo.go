package repository

import (
	"fmt"

	query "github.com/PhuPhuoc/hrm_nextbean_api/rawsql/project_query"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/model"
)

func (store *projectStore) UpdateProject(proid string, info *model.UpdateProjectInfo) error {
	if err_check_id := checkProjectIDExist(store, proid); err_check_id != nil {
		return err_check_id
	}
	rawsql := query.QueryUpdateProject()
	result, err := store.db.Exec(rawsql, info.Name, info.Status, info.Description, info.StartDate, info.Duration, proid)
	if err != nil {
		return fmt.Errorf("error when UpdateProject in store: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error when UpdateProject in store (check affect): %v", err)
	}
	if rowsAffected == 1 {
		return nil // created sucessfully
	} else {
		return fmt.Errorf("error when UpdateProject in store (No user created): %v", err)
	}
}
