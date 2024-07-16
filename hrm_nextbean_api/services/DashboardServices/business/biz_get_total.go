package business

import "github.com/PhuPhuoc/hrm_nextbean_api/services/DashboardServices/model"

type getTotalStore interface {
	GetDashboardTotal() (model.DashboardTotalNumber, error)
}

type getTotalBiz struct {
	store getTotalStore
}

func NewGetTotalBiz(store getTotalStore) *getTotalBiz {
	return &getTotalBiz{
		store: store,
	}
}

func (biz *getTotalBiz) GetTotalBiz() (model.DashboardTotalNumber, error) {
	data, err_query := biz.store.GetDashboardTotal()
	if err_query != nil {
		return model.DashboardTotalNumber{}, err_query
	}
	return data, nil
}
