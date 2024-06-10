package business

import "github.com/PhuPhuoc/hrm_nextbean_api/services/TechnicalServices/model"

type createTechStorage interface {
	CreateTech(info *model.TechnicalCreationInfo) error
}

type createTechBiz struct {
	store createTechStorage
}

func NewCreateTechBiz(store createTechStorage) *createTechBiz {
	return &createTechBiz{store: store}
}

func (biz *createTechBiz) CreateTechBiz(info *model.TechnicalCreationInfo) error {
	if err_query := biz.store.CreateTech(info); err_query != nil {
		return err_query
	}
	return nil
}
