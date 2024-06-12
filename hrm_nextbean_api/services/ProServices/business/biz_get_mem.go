package business

import "github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/model"

type getProMemStore interface {
	GetMem(pro_id string) ([]model.Member, error)
}

type getProMemBiz struct {
	store getProMemStore
}

func NewGetProMemBiz(store getProMemStore) *getProMemBiz {
	return &getProMemBiz{
		store: store,
	}
}

func (biz *getProMemBiz) GetProMemBiz(pro_id string) ([]model.Member, error) {
	data, err_query := biz.store.GetMem(pro_id)
	if err_query != nil {
		return nil, err_query
	}
	return data, nil
}
