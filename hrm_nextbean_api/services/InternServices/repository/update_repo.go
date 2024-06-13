package repository

import (
	"database/sql"
	"fmt"
	"strings"

	intern_query "github.com/PhuPhuoc/hrm_nextbean_api/rawsql/intern_query"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/InternServices/model"
)

func (store *internStore) UpdateIntern(int_id string, intern_update_info *model.InternUpdateInfo) error {
	if err_check_accID_exist := checkInternIDExist(store, int_id); err_check_accID_exist != nil {
		return err_check_accID_exist
	}

	accID, err_get_acc := getAccIDByInternID(store, int_id)
	if err_get_acc != nil {
		return err_get_acc
	}

	if err_check_duplicate := checkDuplicateDataWhenUpdateIntern(store, intern_update_info, *accID, int_id); err_check_duplicate != nil {
		if strings.Contains(err_check_duplicate.Error(), "duplicate data") {
			return err_check_duplicate
		}
		return fmt.Errorf("error when UpdateIntern(checkduplicate) in store: %v", err_check_duplicate)
	}
	rawsql_acc := intern_query.QueryUpdateAccount()
	rawsql_intern := intern_query.QueryUpdateIntern()

	// Begin transaction
	tx, err := store.db.Begin()
	if err != nil {
		return fmt.Errorf("error when UpdateIntern(start transaction) in store: %v", err)
	}

	// Execute first query
	_, err = tx.Exec(rawsql_acc, intern_update_info.UserName, intern_update_info.Email, accID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error when UpdateIntern in store - transaction - account: %v", err)
	}

	// Execute second query
	_, err = tx.Exec(rawsql_intern, intern_update_info.StudentCode, intern_update_info.OjtId, intern_update_info.Avatar, intern_update_info.Gender, intern_update_info.DateOfBirth, intern_update_info.PhoneNumber, intern_update_info.Address, int_id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error when UpdateIntern in store - transaction - intern: %v", err)
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error when committing transaction in store: %v", err)
	}
	return nil
}

func checkDuplicateDataWhenUpdateIntern(store *internStore, intern_update_info *model.InternUpdateInfo, acc_id, intern_id string) error {
	rawsql := intern_query.QueryCheckDulicateDataInInternUpdate(acc_id, intern_id)
	row := store.db.QueryRow(rawsql, intern_update_info.Email, intern_update_info.StudentCode, intern_update_info.PhoneNumber)

	var emailExists, studentcodeExists, phoneExists sql.NullString
	err := row.Scan(&emailExists, &studentcodeExists, &phoneExists)
	if err != nil {
		return err
	}

	duplicateFields := make(map[string]bool)
	if emailExists.Valid {
		duplicateFields["email"] = true
	}
	if studentcodeExists.Valid {
		duplicateFields["student-code"] = true
	}
	if phoneExists.Valid {
		duplicateFields["phone-number"] = true
	}
	isDuplicate := len(duplicateFields) > 0
	if isDuplicate {
		f := ""
		for field := range duplicateFields {
			if f != "" {
				f += ", "
			}
			f += field
		}
		return fmt.Errorf("duplicate data: %v", f)
	}
	return nil
}
