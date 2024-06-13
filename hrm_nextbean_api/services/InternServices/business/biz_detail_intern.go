package business

import (
	"github.com/PhuPhuoc/hrm_nextbean_api/services/InternServices/model"
)

type getDetailInternStore interface {
	GetDetailIntern(int_id string) (*model.InternDetailInfo, error)
}

type getDetailInternBiz struct {
	store getDetailInternStore
}

func NewGetDetailInternBiz(store getDetailInternStore) *getDetailInternBiz {
	return &getDetailInternBiz{
		store: store,
	}
}

func (biz *getDetailInternBiz) GetDetailInternBiz(int_id string) (*model.InternDetailInfo, error) {
	data, err_query := biz.store.GetDetailIntern(int_id)
	if err_query != nil {
		return nil, err_query
	}
	return data, nil
}
