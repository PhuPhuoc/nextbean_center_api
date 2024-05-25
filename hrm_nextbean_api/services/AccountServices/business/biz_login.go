package business

import (
	"fmt"
	"time"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/AccountServices/model"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
)

type loginStorage interface {
	AccountLogin(email, password string, account *model.Account) error
}

type loginBuisiness struct {
	store loginStorage
}

func NewLoginBusiness(store loginStorage) *loginBuisiness {
	return &loginBuisiness{store: store}
}

func (biz *loginBuisiness) Login(login_form *model.LoginForm, data_response map[string]interface{}) error {
	account := new(model.Account)
	if err_repo := biz.store.AccountLogin(login_form.Email, login_form.Password, account); err_repo != nil {
		return err_repo
	}
	// create jwt
	currentTime := time.Now()
	expirationTime := currentTime.Add(120 * time.Minute)
	expUnix := expirationTime.Unix()

	payload_jwt := map[string]interface{}{
		"id":       account.Id,
		"role":     account.Role,
		"exp_date": expUnix,
	}
	token, err_create_jwt := utils.CreateJWT(payload_jwt)
	if err_create_jwt != nil {
		return fmt.Errorf("cannot create jwt: %v", err_create_jwt)
	}
	data_response["account_info"] = account
	data_response["token"] = token
	return nil
}
