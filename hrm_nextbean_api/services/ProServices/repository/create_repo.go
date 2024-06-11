package repository

import (
	"fmt"

	query "github.com/PhuPhuoc/hrm_nextbean_api/rawsql/project_query"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/model"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
	"github.com/google/uuid"
)

func (store *projectStore) CreateProject(info *model.ProjectCreationInfo) error {
	newUUID := uuid.New()
	rawsql := query.QueryCreateProject()
	result, err := store.db.Exec(rawsql, newUUID, info.Name, "not_start", info.Description, info.StartAt, info.Duration, utils.CreateDateTimeCurrentFormated())
	if err != nil {
		return fmt.Errorf("error when CreateProject in store: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error when CreateProject in store (check affect): %v", err)
	}
	if rowsAffected == 1 {
		return nil // created sucessfully
	} else {
		return fmt.Errorf("error when CreateProject in store (No user created): %v", err)
	}
}
