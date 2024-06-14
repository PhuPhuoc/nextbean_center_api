package business

import (
	"github.com/PhuPhuoc/hrm_nextbean_api/common"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/model"
)

type getMemNotInProStore interface {
	GetMemNotInPro(pro_id string, pagin *common.Pagination, filter *model.MemberFilter) ([]model.Member, error)
}

type getMemNotInProBiz struct {
	store getMemNotInProStore
}

func NewGetMemNotInProBiz(store getMemNotInProStore) *getMemNotInProBiz {
	return &getMemNotInProBiz{
		store: store,
	}
}

func (biz *getMemNotInProBiz) GetMemNotInProBiz(pro_id string, pagin *common.Pagination, filter *model.MemberFilter) ([]model.Member, error) {
	data, err_query := biz.store.GetMemNotInPro(pro_id, pagin, filter)
	if err_query != nil {
		return nil, err_query
	}
	return data, nil
}
