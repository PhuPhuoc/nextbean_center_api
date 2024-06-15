package repository

import (
	"database/sql"
	"fmt"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/InternServices/model"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
	"github.com/google/uuid"
)

func (store *internStore) CreateIntern(intern_cre_info *model.InternCreation) error {
	isDuplicate, duplicateFields, err_checkDup := checkInfoDuplicateInTableIntern(store, intern_cre_info.Email, intern_cre_info.StudentCode, intern_cre_info.PhoneNumber)
	if err_checkDup != nil {
		return err_checkDup
	}

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
	rawsql_acc := `insert into account(id, user_name, email, password, role, created_at) values (?,?,?,?,?,?)`
	rawsql_intern := `insert into intern(id,student_code,account_id,ojt_id,avatar,gender,date_of_birth,phone_number,address) values (?,?,?,?,?,?,?,?,?)`
	newUUID_account := uuid.New()
	newUUID_intern := uuid.New()
	pwdHash := utils.GenerateHash(intern_cre_info.Password)

	// Begin transaction
	tx, err := store.db.Begin()
	if err != nil {
		return fmt.Errorf("error db in CreateIntern transaction: %v", err)
	}

	// Execute first query
	_, err = tx.Exec(rawsql_acc, newUUID_account, intern_cre_info.UserName, intern_cre_info.Email, pwdHash, "user", utils.CreateDateTimeCurrentFormated())
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error db in CreateIntern transaction - account: %v", err)
	}

	// Execute second query
	_, err = tx.Exec(rawsql_intern, newUUID_intern, intern_cre_info.StudentCode, newUUID_account, intern_cre_info.OjtId, intern_cre_info.Avatar, intern_cre_info.Gender, intern_cre_info.DateOfBirth, intern_cre_info.PhoneNumber, intern_cre_info.Address)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error db in CreateIntern - transaction - intern: %v", err)
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error when committing transaction in CreateIntern: %v", err)
	}

	return nil
}

func queryCheckDulicateDataInIntern() string {
	start := `select`
	part1 := ` case when exists (select 1 from account where email = ?) then 'email' end as email_exists,`
	part2 := ` case when exists (select 1 from intern where id = ?) then 'id' end as studentcode_exists,`
	part3 := ` case when exists (select 1 from intern where phone_number = ?) then 'phone_number' end as phonenumber_exists`
	end := `from DUAL;`
	return start + ` ` + part1 + ` ` + part2 + ` ` + part3 + ` ` + end
}

func checkInfoDuplicateInTableIntern(store *internStore, email, studentcode, phone string) (bool, map[string]bool, error) {
	rawsql := queryCheckDulicateDataInIntern()
	row := store.db.QueryRow(rawsql, email, studentcode, phone)

	var emailExists, studentcodeExists, phoneExists sql.NullString
	err := row.Scan(&emailExists, &studentcodeExists, &phoneExists)
	if err != nil {
		return false, nil, fmt.Errorf("error db in checkInfoDuplicateInTableIntern - : %v", err)
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
	return isDuplicate, duplicateFields, nil
}
