package business

import (
	"github.com/PhuPhuoc/hrm_nextbean_api/services/AccountServices/model"
)

type LoginStorage interface {
	AccountLogin(email, password string, account *model.Account) error
}

type loginBuisiness struct {
	store LoginStorage
}

func NewLoginBusiness(store LoginStorage) *loginBuisiness {
	return &loginBuisiness{store: store}
}

func (biz *loginBuisiness) Login(login_form *model.LoginForm, data_response map[string]interface{}) error {
	account := new(model.Account)
	data_response["account_info"] = "acc"
	data_response["token"] = "token"
	if err_repo := biz.store.AccountLogin(login_form.Email, login_form.Password, account); err_repo != nil {
		return err_repo
	}
	// create jwt
	return nil
}
