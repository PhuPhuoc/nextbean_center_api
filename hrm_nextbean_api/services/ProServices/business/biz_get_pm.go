package business

import "github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/model"

type getPMStore interface {
	GetPM(pro_id string) ([]model.PM, error)
}

type getPMBiz struct {
	store getPMStore
}

func NewGetPMBiz(store getPMStore) *getPMBiz {
	return &getPMBiz{
		store: store,
	}
}

func (biz *getPMBiz) GetPMBiz(pro_id string) ([]model.PM, error) {
	data, err_query := biz.store.GetPM(pro_id)
	if err_query != nil {
		return nil, err_query
	}
	return data, nil
}
