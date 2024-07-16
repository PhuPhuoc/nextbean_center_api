package business

import "github.com/PhuPhuoc/hrm_nextbean_api/services/DashboardServices/model"

type getInProgressOJTStore interface {
	GetDashboardInpgrogressOJT() ([]model.DashboardOJTInProgress, error)
}

type getInProgressOJTBiz struct {
	store getInProgressOJTStore
}

func NewGetInProgressOJTBiz(store getInProgressOJTStore) *getInProgressOJTBiz {
	return &getInProgressOJTBiz{
		store: store,
	}
}

func (biz *getInProgressOJTBiz) GetInProgressOJTBiz() ([]model.DashboardOJTInProgress, error) {
	data, err_query := biz.store.GetDashboardInpgrogressOJT()
	if err_query != nil {
		return nil, err_query
	}
	return data, nil
}
