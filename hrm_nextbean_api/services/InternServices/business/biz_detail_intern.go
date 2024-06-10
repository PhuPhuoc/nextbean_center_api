package business

import (
	"github.com/PhuPhuoc/hrm_nextbean_api/services/InternServices/model"
)

type getDetailInternStore interface {
	GetDetailIntern(acc_id string) (*model.InternDetailInfo, error)
}

type getDetailInternBiz struct {
	store getDetailInternStore
}

func NewGetDetailInternBiz(store getDetailInternStore) *getDetailInternBiz {
	return &getDetailInternBiz{
		store: store,
	}
}

func (biz *getDetailInternBiz) GetDetailInternBiz(acc_id string) (*model.InternDetailInfo, error) {
	data, err_query := biz.store.GetDetailIntern(acc_id)
	if err_query != nil {
		return nil, err_query
	}
	return data, nil
}
