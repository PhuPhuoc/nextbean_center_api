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

func (biz *createAccountBuisiness) CreateNewAccountBiz(acc_cre_info *model.AccountCreationInfo) error {
	if err_query := biz.store.CreateAccount(acc_cre_info); err_query != nil {
		return err_query
	}
	return nil
}
