package business

import "github.com/PhuPhuoc/hrm_nextbean_api/services/OJTServices/model"

type createOJTStorage interface {
	CreateOJT(info *model.OJTCreationInfo) error
}

type createOJTBiz struct {
	store createOJTStorage
}

func NewCreateOJTBiz(store createOJTStorage) *createOJTBiz {
	return &createOJTBiz{store: store}
}

func (biz *createOJTBiz) CreateOJTBiz(info *model.OJTCreationInfo) error {
	if err_query := biz.store.CreateOJT(info); err_query != nil {
		return err_query
	}
	return nil
}
