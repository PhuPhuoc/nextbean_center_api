package business

import (
	"github.com/PhuPhuoc/hrm_nextbean_api/common"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/InternServices/model"
)

type getInternStore interface {
	GetIntern(pagin *common.Pagination, filter *model.InternFilter) ([]model.Intern, error)
}

type getInternBusiness struct {
	store getInternStore
}

func NewGetInternBusiness(store getInternStore) *getInternBusiness {
	return &getInternBusiness{
		store: store,
	}
}

func (biz *getInternBusiness) GetInternBiz(pagin *common.Pagination, filter *model.InternFilter) ([]model.Intern, error) {
	data, err_query := biz.store.GetIntern(pagin, filter)
	if err_query != nil {
		return nil, err_query
	}
	return data, nil
}
