package repository

import (
	"fmt"
	"strconv"
	"strings"

	query "github.com/PhuPhuoc/hrm_nextbean_api/rawsql/ojt_query"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
)

func (store *ojtStore) DeleteOJT(id string) error {
	ojt_id, err_parse := strconv.Atoi(id)
	if err_parse != nil {
		return fmt.Errorf("ojt'id: %v invalid", id)
	}
	if err_exist := checkExistID(store, ojt_id); err_exist != nil {
		if strings.Contains(err_exist.Error(), "id_not_exist") {
			return fmt.Errorf("ojt'id: %v not exist", id)
		}
		return err_exist
	}

	rawsql := query.QueryDeleteOJT()
	result, err := store.db.Exec(rawsql, utils.CreateDateTimeCurrentFormated(), id)
	if err != nil {
		return fmt.Errorf("error when DeleteOJT in store: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error when DeleteOJT in store (check affect): %v", err)
	}
	if rowsAffected == 1 {
		return nil // update sucessfully
	} else {
		return fmt.Errorf("error when DeleteOJT in store (No user created): %v", err)
	}
}
