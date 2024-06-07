package business

import (
	"github.com/PhuPhuoc/hrm_nextbean_api/common"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/AccountServices/model"
)

type getAccountStore interface {
	GetAccount(pagin *common.Pagination, filter *model.AccountFilter) ([]model.Account, error)
}

type getAccountBusiness struct {
	store getAccountStore
}

func NewGetAccountBusiness(store getAccountStore) *getAccountBusiness {
	return &getAccountBusiness{
		store: store,
	}
}

func (biz *getAccountBusiness) GetAccountBiz(pagin *common.Pagination, filter *model.AccountFilter) ([]model.Account, error) {
	data, err_query := biz.store.GetAccount(pagin, filter)
	if err_query != nil {
		return nil, err_query
	}
	return data, nil
}
