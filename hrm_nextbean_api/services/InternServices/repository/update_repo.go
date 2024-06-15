package repository

import (
	"database/sql"
	"fmt"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/InternServices/model"
)

func (store *internStore) UpdateIntern(int_id string, intern_update_info *model.InternUpdateInfo) error {
	if err_check_accID_exist := checkInternIDExists(store, int_id); err_check_accID_exist != nil {
		return err_check_accID_exist
	}

	accID, err_get_acc := getAccIDByInternID(store, int_id)
	if err_get_acc != nil {
		return err_get_acc
	}

	if err_check_duplicate := checkDuplicateDataWhenUpdateIntern(store, intern_update_info, *accID, int_id); err_check_duplicate != nil {
		return err_check_duplicate
	}
	rawsql_acc := `update account set user_name = ?, email = ? where id = ?`
	rawsql_intern := `update intern set student_code = ?, ojt_id = ?, avatar = ?, gender = ?, date_of_birth = ?, phone_number = ?, address = ? where id = ?`

	// Begin transaction
	tx, err := store.db.Begin()
	if err != nil {
		return fmt.Errorf("error in UpdateIntern transaction: %v", err)
	}

	// Execute first query
	_, err = tx.Exec(rawsql_acc, intern_update_info.UserName, intern_update_info.Email, accID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error in UpdateIntern - transaction - account: %v", err)
	}

	// Execute second query
	_, err = tx.Exec(rawsql_intern, intern_update_info.StudentCode, intern_update_info.OjtId, intern_update_info.Avatar, intern_update_info.Gender, intern_update_info.DateOfBirth, intern_update_info.PhoneNumber, intern_update_info.Address, int_id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error in UpdateIntern - transaction - intern: %v", err)
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error in transaction - commit: %v", err)
	}
	return nil
}

func checkDuplicateDataWhenUpdateIntern(store *internStore, intern_update_info *model.InternUpdateInfo, acc_id, intern_id string) error {
	rawsql := queryCheckDulicateDataInInternUpdate(acc_id, intern_id)
	row := store.db.QueryRow(rawsql, intern_update_info.Email, intern_update_info.StudentCode, intern_update_info.PhoneNumber)

	var emailExists, studentcodeExists, phoneExists sql.NullString
	err := row.Scan(&emailExists, &studentcodeExists, &phoneExists)
	if err != nil {
		return fmt.Errorf("error in checkDuplicateDataWhenUpdateIntern: %v", err.Error())
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
		return fmt.Errorf("invalid-request: duplicate data [%v]", f)
	}
	return nil
}

func queryCheckDulicateDataInInternUpdate(acc_id, in_id string) string {
	start := `select`
	part1 := fmt.Sprintf(` case when exists (select 1 from account where email = ? and id != '%s') then 'email' end as email_exists,`, acc_id)
	part2 := fmt.Sprintf(` case when exists (select 1 from intern where student_code = ? and id != '%s') then 'id' end as studentcode_exists,`, in_id)
	part3 := fmt.Sprintf(` case when exists (select 1 from intern where phone_number = ? and id != '%s') then 'phone_number' end as phonenumber_exists`, in_id)
	end := `from DUAL;`
	return start + ` ` + part1 + ` ` + part2 + ` ` + part3 + ` ` + end
}

func getAccIDByInternID(store *internStore, int_id string) (*string, error) {
	var accID *string
	query := `select acc.id from account acc join intern i on i.account_id = acc.id where i.id = ?`
	err := store.db.QueryRow(query, int_id).Scan(&accID)
	if err != nil {
		return nil, fmt.Errorf("error in getAccIDByInternID: %v", err)
	}
	return accID, nil
}
