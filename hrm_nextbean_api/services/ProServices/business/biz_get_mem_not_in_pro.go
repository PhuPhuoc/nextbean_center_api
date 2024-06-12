package business

import "github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/model"

type getMemNotInProStore interface {
	GetMemNotInPro(pro_id string) ([]model.Member, error)
}

type getMemNotInProBiz struct {
	store getMemNotInProStore
}

func NewGetMemNotInProBiz(store getMemNotInProStore) *getMemNotInProBiz {
	return &getMemNotInProBiz{
		store: store,
	}
}

func (biz *getMemNotInProBiz) GetMemNotInProBiz(pro_id string) ([]model.Member, error) {
	data, err_query := biz.store.GetMemNotInPro(pro_id)
	if err_query != nil {
		return nil, err_query
	}
	return data, nil
}
