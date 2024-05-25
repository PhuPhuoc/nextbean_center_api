package business

import "github.com/PhuPhuoc/hrm_nextbean_api/services/AccountServices/model"

type createAccountStorage interface {
	CreateAccount(acc_cre_info *model.AccountCreationInfo) error
}

type createAccountBuisiness struct {
	store createAccountStorage
}

func NewCreateAccountBusiness(store createAccountStorage) *createAccountBuisiness {
	return &createAccountBuisiness{store: store}
}

func (biz *createAccountBuisiness) CreateNewAccount(acc_cre_info *model.AccountCreationInfo) error {
	if err_querry := biz.store.CreateAccount(acc_cre_info); err_querry != nil {
		return err_querry
	}
	return nil
}
