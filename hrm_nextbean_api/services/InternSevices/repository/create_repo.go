package repository

import (
	"database/sql"
	"fmt"

	intern_query "github.com/PhuPhuoc/hrm_nextbean_api/rawsql/intern_query"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/InternSevices/model"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
	"github.com/google/uuid"
)

func (store *internStore) checkInfoDuplicateInTableIntern(email, studentcode, phone string) (bool, map[string]bool, error) {
	rawsql := intern_query.QueryCheckDulicateDataInIntern()

	row := store.db.QueryRow(rawsql, email, studentcode, phone)

	var emailExists, studentcodeExists, phoneExists sql.NullString
	err := row.Scan(&emailExists, &studentcodeExists, &phoneExists)
	if err != nil {
		return false, nil, err
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

func (store *internStore) CreateIntern(intern_cre_info *model.InternCreation) error {
	isDuplicate, duplicateFields, err_checkDup := store.checkInfoDuplicateInTableIntern(intern_cre_info.Email, intern_cre_info.StudentCode, intern_cre_info.PhoneNumber)
	if err_checkDup != nil {
		return fmt.Errorf("error when CreateIntern(checkInfoDuplicateInTableIntern) in store: %v", err_checkDup)
	}

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
	pwdHash := utils.GenerateHash(intern_cre_info.Password)
	rawsql_acc := intern_query.QueryCreateNewAccount()
	rawsql_intern := intern_query.QueryCreateNewIntern()
	newUUID := uuid.New()

	// Begin transaction
	tx, err := store.db.Begin()
	if err != nil {
		return fmt.Errorf("error when CreateIntern(start transaction) in store: %v", err)
	}

	// Execute first query
	_, err = tx.Exec(rawsql_acc, newUUID, intern_cre_info.UserName, intern_cre_info.Email, pwdHash, "user", utils.CreateDateTimeCurrentFormated())
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error when CreateIntern in store - transaction - account: %v", err)
	}

	// Execute second query
	_, err = tx.Exec(rawsql_intern, intern_cre_info.StudentCode, newUUID, intern_cre_info.OjtId, intern_cre_info.Avatar, intern_cre_info.Gender, intern_cre_info.DateOfBirth, intern_cre_info.PhoneNumber, intern_cre_info.Address)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error when CreateIntern in store - transaction - intern: %v", err)
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error when committing transaction in store: %v", err)
	}

	return nil
}
