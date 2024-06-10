package business

import (
	"github.com/PhuPhuoc/hrm_nextbean_api/common"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/TechnicalServices/model"
)

type getTechStorage interface {
	GetTech(pagin *common.Pagination, filter *model.FilterTechnical) ([]model.Technical, error)
}

type getTechBiz struct {
	store getTechStorage
}

func NewGetTechBiz(store getTechStorage) *getTechBiz {
	return &getTechBiz{store: store}
}

func (biz *getTechBiz) GetTechBiz(pagin *common.Pagination, filter *model.FilterTechnical) ([]model.Technical, error) {
	data, err_query := biz.store.GetTech(pagin, filter)
	if err_query != nil {
		return nil, err_query
	}
	return data, nil
}
