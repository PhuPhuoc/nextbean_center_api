package repository

import (
	"fmt"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/AccountServices/model"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
	"github.com/google/uuid"
)

func (store *accountStore) CreateAccount(acc_cre_info *model.AccountCreationInfo) error {
	// todo 1: check email of new account are unique in db
	var flag bool
	rawsql_checkEmailExist := `select exists(select 1 from account where email = ?)`
	if err_query := store.db.QueryRow(rawsql_checkEmailExist, acc_cre_info.Email).Scan(&flag); err_query != nil {
		return fmt.Errorf("error when CreateAccount in store (check exists email): %v", err_query)
	}
	if flag {
		return fmt.Errorf("email: %v already exists", acc_cre_info.Email)
	}

	// todo 1.1: reject if new account'role is user
	if acc_cre_info.Role == "user" {
		return fmt.Errorf("please create an account with the user role in api 'intern'")
	}

	// todo 2: create new account
	rawsql_insertNewAccount := `insert into account (id, user_name, email, password, role, created_at) values (?,?,?,?,?,?)`
	newUUID := uuid.New()
	pwdHash := utils.GenerateHash(acc_cre_info.Password)
	result, err := store.db.Exec(rawsql_insertNewAccount, newUUID, acc_cre_info.UserName, acc_cre_info.Email, pwdHash, acc_cre_info.Role, utils.CreateDateTimeCurrentFormated())
	if err != nil {
		return fmt.Errorf("error when CreateAccount in store: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error when CreateAccount in store (check affect): %v", err)
	}
	if rowsAffected == 1 {
		return nil // created sucessfully
	} else {
		return fmt.Errorf("error when CreateAccount in store (No user created): %v", err)
	}
}
