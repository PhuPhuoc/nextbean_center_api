package business

import "github.com/PhuPhuoc/hrm_nextbean_api/services/OJTServices/model"

type updateOJTStorage interface {
	UpdateOJT(info *model.UpdateOJTInfo) error
}

type updateOJTBiz struct {
	store updateOJTStorage
}

func NewUpdateOJTBiz(store updateOJTStorage) *updateOJTBiz {
	return &updateOJTBiz{store: store}
}

func (biz *updateOJTBiz) UpdateOJTBiz(info *model.UpdateOJTInfo) error {
	if err_query := biz.store.UpdateOJT(info); err_query != nil {
		return err_query
	}
	return nil
}
