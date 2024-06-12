package business

import "github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/model"

type getPMNotInProStore interface {
	GetPMNotInPro(pro_id string) ([]model.PM, error)
}

type getPMNotInProBiz struct {
	store getPMNotInProStore
}

func NewGetPMNotInProBiz(store getPMNotInProStore) *getPMNotInProBiz {
	return &getPMNotInProBiz{
		store: store,
	}
}

func (biz *getPMNotInProBiz) GetPMNotInProBiz(pro_id string) ([]model.PM, error) {
	data, err_query := biz.store.GetPMNotInPro(pro_id)
	if err_query != nil {
		return nil, err_query
	}
	return data, nil
}
