package repository

import (
	"fmt"

	query "github.com/PhuPhuoc/hrm_nextbean_api/rawsql/tech_query"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/TechnicalServices/model"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
)

func (store *techStore) CreateTech(info *model.TechnicalCreationInfo) error {
	rawsql := query.QueryCreateTechnical()
	result, err := store.db.Exec(rawsql, info.TechnicalSkill, utils.CreateDateTimeCurrentFormated())
	if err != nil {
		return fmt.Errorf("error when CreateTech in store: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error when CreateTech in store (check affect): %v", err)
	}
	if rowsAffected == 1 {
		return nil // created sucessfully
	} else {
		return fmt.Errorf("error when CreateTech in store (No user created): %v", err)
	}
}
