package business

import "github.com/PhuPhuoc/hrm_nextbean_api/services/OJTServices/model"

type updateOJTStorage interface {
	UpdateOJT(ojt_id string, info *model.UpdateOJTInfo) error
}

type updateOJTBiz struct {
	store updateOJTStorage
}

func NewUpdateOJTBiz(store updateOJTStorage) *updateOJTBiz {
	return &updateOJTBiz{store: store}
}

func (biz *updateOJTBiz) UpdateOJTBiz(ojt_id string, info *model.UpdateOJTInfo) error {
	if err_query := biz.store.UpdateOJT(ojt_id, info); err_query != nil {
		return err_query
	}
	return nil
}
