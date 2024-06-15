package repository

import (
	"fmt"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/OJTServices/model"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
)

func (store *ojtStore) CreateOJT(info *model.OJTCreationInfo) error {
	rawsql := `insert into ojt(semester, university, start_at, end_at, created_at) values (?,?,?,?,?)`
	result, err := store.db.Exec(rawsql, info.Semester, info.University, info.StartAt, info.EndAt, utils.CreateDateTimeCurrentFormated())
	if err != nil {
		return fmt.Errorf("error when CreateOJT in store: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error when CreateOJT in store (check affect): %v", err)
	}
	if rowsAffected == 1 {
		return nil // created sucessfully
	} else {
		return fmt.Errorf("error when CreateOJT in store (No user created): %v", err)
	}
}
