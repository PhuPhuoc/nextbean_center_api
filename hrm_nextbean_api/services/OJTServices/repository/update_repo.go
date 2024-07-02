package repository

import (
	"fmt"
	"strconv"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/OJTServices/model"
)

func (store *ojtStore) UpdateOJT(ojt_id string, info *model.UpdateOJTInfo) error {
	id, err_exist := checkOjtIDExists(store, ojt_id)
	if err_exist != nil {
		return err_exist
	}

	rawsql := `update ojt set semester=?, university=?, start_at=?, end_at=?, status=? where id = ?`
	result, err := store.db.Exec(rawsql, info.Semester, info.University, info.StartAt, info.EndAt, info.Status, id)
	if err != nil {
		return fmt.Errorf("error db in UpdateOJT: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error db in UpdateOJT (check affect): %v", err)
	}
	if rowsAffected == 0 {
		return nil
	}
	return nil
}

func checkOjtIDExists(store *ojtStore, ojt_id string) (int, error) {
	id, err_int := strconv.Atoi(ojt_id)
	if err_int != nil {
		return 0, fmt.Errorf("invalid-request: id '%v' is invalid", id)
	}
	var flag bool
	rawsql := `select exists(select 1 from ojt where id = ? and deleted_at is null)`
	if err_query := store.db.QueryRow(rawsql, id).Scan(&flag); err_query != nil {
		return 0, fmt.Errorf("error db in checkOjtIDExists: %v", err_query)
	}
	if !flag {
		return 0, fmt.Errorf("invalid-request: id '%v' not exist", id)
	}
	return id, nil
}
