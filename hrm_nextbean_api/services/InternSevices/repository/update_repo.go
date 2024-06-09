package repository

import (
	"database/sql"
	"fmt"
	"strings"

	intern_query "github.com/PhuPhuoc/hrm_nextbean_api/rawsql/intern_query"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/InternSevices/model"
)

func (store *internStore) UpdateIntern(intern_update_info *model.InternUpdateInfo) error {
	current_student_code, err_check := getStudentCodeByAccountID(store, intern_update_info.AccountId)
	if err_check != nil {
		return err_check
	}
	if err_check_duplicate := checkDuplicateDataWhenUpdateIntern(store, intern_update_info, *current_student_code); err_check_duplicate != nil {
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
	_, err = tx.Exec(rawsql_acc, intern_update_info.UserName, intern_update_info.Email, intern_update_info.AccountId)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error when UpdateIntern in store - transaction - account: %v", err)
	}

	// Execute second query
	_, err = tx.Exec(rawsql_intern, intern_update_info.StudentCode, intern_update_info.OjtId, intern_update_info.Avatar, intern_update_info.Gender, intern_update_info.DateOfBirth, intern_update_info.PhoneNumber, intern_update_info.Address, current_student_code)
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

func getStudentCodeByAccountID(store *internStore, accID string) (*string, error) {
	var stuID *string
	query := intern_query.QueryGetCurrentStudentCodeByAccountID()
	err := store.db.QueryRow(query, accID).Scan(&stuID)
	if err != nil {
		return nil, fmt.Errorf("error when get current student-code: %v", err)
	}
	return stuID, nil
}

func checkDuplicateDataWhenUpdateIntern(store *internStore, intern_update_info *model.InternUpdateInfo, current_student_code string) error {
	rawsql := intern_query.QueryCheckDulicateDataInInternUpdate(intern_update_info.AccountId, current_student_code)
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
