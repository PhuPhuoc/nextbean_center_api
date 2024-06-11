package business

import (
	"github.com/PhuPhuoc/hrm_nextbean_api/common"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/model"
)

type getProjectStore interface {
	GetProject(pagin *common.Pagination, filter *model.ProjectFilter) ([]model.Project, error)
}

type getProBiz struct {
	store getProjectStore
}

func NewGetProBiz(store getProjectStore) *getProBiz {
	return &getProBiz{
		store: store,
	}
}

func (biz *getProBiz) GetProBiz(pagin *common.Pagination, filter *model.ProjectFilter) ([]model.Project, error) {
	data, err_query := biz.store.GetProject(pagin, filter)
	if err_query != nil {
		return nil, err_query
	}
	return data, nil
}
