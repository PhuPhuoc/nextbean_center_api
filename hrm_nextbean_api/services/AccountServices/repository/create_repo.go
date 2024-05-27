package repository

import (
	"fmt"
	"strings"

	query "github.com/PhuPhuoc/hrm_nextbean_api/rawsql/account_query"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/AccountServices/model"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
	"github.com/google/uuid"
)

func (store *AccountStore) CreateAccount(acc_cre_info *model.AccountCreationInfo) error {
	if err_check_email_exist := store.checkEmailExist(acc_cre_info.Email); err_check_email_exist != nil {
		if strings.Contains(err_check_email_exist.Error(), "email exist") {
			return fmt.Errorf("email: %v already exists", acc_cre_info.Email)
		}
		return fmt.Errorf("error when CreateAccount(checkEmailExist) in store: %v", err_check_email_exist)
	}
	pwdHash := utils.GenerateHash(acc_cre_info.Password)
	rawsql := query.QueryCreateNewAccount()
	newUUID := uuid.New()
	result, err := store.db.Exec(rawsql, newUUID, acc_cre_info.UserName, acc_cre_info.Email, pwdHash, acc_cre_info.Role, utils.CreateDateTimeCurrentFormated())
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

func (store *AccountStore) checkEmailExist(email string) error {
	var flag bool
	rawsql := query.QueryCheckExistEmail()
	if err_query := store.db.QueryRow(rawsql, email).Scan(&flag); err_query != nil {
		return fmt.Errorf("error when CreateAccount in store (check exist email): %v", err_query)
	}
	if flag {
		return fmt.Errorf("email exist")
	}
	return nil // user'email not exist in db => ready to create
}

//0945909397
