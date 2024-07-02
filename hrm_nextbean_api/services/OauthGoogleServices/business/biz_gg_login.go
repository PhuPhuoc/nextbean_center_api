package business

import (
	"fmt"
	"time"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/OauthGoogleServices/model"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
)

type googleLoginStorage interface {
	AccountGoogleLogin(email string, account *model.AccountLoginGG) error
}

type googleLoginBuisiness struct {
	store googleLoginStorage
}

func NewLoginBusiness(store googleLoginStorage) *googleLoginBuisiness {
	return &googleLoginBuisiness{store: store}
}

func (biz *googleLoginBuisiness) Login(email string, data_response map[string]interface{}) error {
	account := new(model.AccountLoginGG)
	if err_repo := biz.store.AccountGoogleLogin(email, account); err_repo != nil {
		return err_repo
	}
	// create jwt
	currentTime := time.Now()
	expirationTime := currentTime.Add(120 * time.Minute)
	expUnix := expirationTime.Unix()

	payload_jwt := map[string]interface{}{
		"id":       account.Id,
		"role":     account.Role,
		"username": account.UserName,
		"exp_date": expUnix,
	}
	token, err_create_jwt := utils.CreateJWT(payload_jwt)
	if err_create_jwt != nil {
		return fmt.Errorf("cannot create jwt: %v", err_create_jwt)
	}
	data_response["account_info"] = account
	data_response["token"] = token
	fmt.Printf("login google email: %v\n", account.Email)
	return nil
}
