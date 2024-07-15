package business

import (
	"github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/model"
)

type getProDetailStore interface {
	GetDetail(pro_id string) (model.ProjectDetail, error)
}

type getProDetailBiz struct {
	store getProDetailStore
}

func NewGetProDetailBiz(store getProDetailStore) *getProDetailBiz {
	return &getProDetailBiz{
		store: store,
	}
}

func (biz *getProDetailBiz) GetProDetailBiz(pro_id string) (model.ProjectDetail, error) {
	data, err_query := biz.store.GetDetail(pro_id)
	if err_query != nil {
		return model.ProjectDetail{}, err_query
	}
	return data, nil
}
