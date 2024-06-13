package repository

import (
	"fmt"
	"strconv"
	"strings"

	query "github.com/PhuPhuoc/hrm_nextbean_api/rawsql/ojt_query"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/OJTServices/model"
)

func (store *ojtStore) UpdateOJT(ojt_id string, info *model.UpdateOJTInfo) error {
	id, err_int := strconv.Atoi(ojt_id)
	if err_int != nil {
		return err_int
	}
	if err_exist := checkExistID(store, id); err_exist != nil {
		if strings.Contains(err_exist.Error(), "id_not_exist") {
			return fmt.Errorf("ojt'id: %v not exist", ojt_id)
		}
		return err_exist
	}

	rawsql := query.QueryUpdateOJT()
	result, err := store.db.Exec(rawsql, info.Semester, info.University, info.StartAt, info.EndAt, id)
	if err != nil {
		return fmt.Errorf("error when UpdateOJT in store: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error when UpdateOJT in store (check affect): %v", err)
	}
	if rowsAffected == 1 {
		return nil // update sucessfully
	} else {
		return fmt.Errorf("error when UpdateOJT in store (No user created): %v", err)
	}
}

func checkExistID(store *ojtStore, id int) error {
	var flag bool
	rawsql := query.QueryCheckExistID()
	if err_query := store.db.QueryRow(rawsql, id).Scan(&flag); err_query != nil {
		return fmt.Errorf("error when UpdateOJT in store (check exists id): %v", err_query)
	}
	if !flag {
		return fmt.Errorf("id_not_exist")
	}
	return nil
}
