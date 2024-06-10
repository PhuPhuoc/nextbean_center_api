package business

import (
	"github.com/PhuPhuoc/hrm_nextbean_api/common"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/OJTServices/model"
)

type getOJTStorage interface {
	GetOJT(pagin *common.Pagination, filter *model.FilterOJT) ([]model.OJT, error)
}

type getOJTBiz struct {
	store getOJTStorage
}

func NewGetOJTBiz(store getOJTStorage) *getOJTBiz {
	return &getOJTBiz{store: store}
}

func (biz *getOJTBiz) GetOJTBiz(pagin *common.Pagination, filter *model.FilterOJT) ([]model.OJT, error) {
	data, err_query := biz.store.GetOJT(pagin, filter)
	if err_query != nil {
		return nil, err_query
	}
	return data, nil
}
