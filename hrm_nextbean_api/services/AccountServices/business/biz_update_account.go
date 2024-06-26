package business

import "github.com/PhuPhuoc/hrm_nextbean_api/services/AccountServices/model"

type updateAccountStorage interface {
	UpdateAccount(accountID string, acc_update_info *model.UpdateAccountInfo) error
}

type updateAccountBuisiness struct {
	store updateAccountStorage
}

func NewUpdateAccountBusiness(store updateAccountStorage) *updateAccountBuisiness {
	return &updateAccountBuisiness{store: store}
}

func (biz *updateAccountBuisiness) UpdateAccountBiz(accountID string, acc_update_info *model.UpdateAccountInfo) error {
	if err_query := biz.store.UpdateAccount(accountID, acc_update_info); err_query != nil {
		return err_query
	}
	return nil
}
