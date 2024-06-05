package repository

import (
	"fmt"
	"strings"

	query "github.com/PhuPhuoc/hrm_nextbean_api/rawsql/account_query"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/AccountServices/model"
)

func (store *AccountStore) UpdateAccount(acc_update_info *model.UpdateAccountInfo) error {

	if err_check_id_exist := store.checkIdExist(acc_update_info.Id); err_check_id_exist != nil {
		if strings.Contains(err_check_id_exist.Error(), "id not exist") {
			return fmt.Errorf("account'ID not exists")
		}
		return fmt.Errorf("error when UpdateAccount(checkIdExist) in store: %v", err_check_id_exist)
	}

	if err_check_email_exist := store.checkEmailExistWithID(acc_update_info.Email, acc_update_info.Id); err_check_email_exist != nil {
		if strings.Contains(err_check_email_exist.Error(), "email_exist") {
			return fmt.Errorf("email: %v already exists", acc_update_info.Email)
		}
		return fmt.Errorf("error when UpdateAccount(checkEmailExist) in store: %v", err_check_email_exist)
	}

	rawsql := query.QueryUpdateAccount()
	result, err := store.db.Exec(rawsql, acc_update_info.UserName, acc_update_info.Email, acc_update_info.Role, acc_update_info.Id)
	if err != nil {
		return fmt.Errorf("error when UpdateAccount in store: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error when UpdateAccount in store (check affect): %v", err)
	}
	if rowsAffected == 1 {
		return nil // created sucessfully
	} else {
		return fmt.Errorf("error when UpdateAccount in store (No user updated): %v", err)
	}
}

func (store *AccountStore) checkIdExist(id string) error {
	var flag bool = false
	rawsql := query.QueryIdExist()
	if err_query := store.db.QueryRow(rawsql, id).Scan(&flag); err_query != nil {
		return fmt.Errorf("error when UpdateAccount in store (check Id exist): %v", err_query)
	}
	if !flag {
		return fmt.Errorf("id not exist")
	}
	return nil // user'id exist in db => ready to update
}

func (store *AccountStore) checkEmailExistWithID(email, id string) error {
	var flag bool = false
	rawsql := query.QueryCheckExistEmailWithID()
	if err_query := store.db.QueryRow(rawsql, email, id).Scan(&flag); err_query != nil {
		return fmt.Errorf("error when UpdateAccount in store (check exist email): %v", err_query)
	}
	if flag {
		return fmt.Errorf("email_exist")
	}
	return nil // user'email not exist in db => ready to create
}
