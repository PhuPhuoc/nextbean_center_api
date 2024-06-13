package business

import "github.com/PhuPhuoc/hrm_nextbean_api/services/InternServices/model"

type updateInternStorage interface {
	UpdateIntern(int_id string, intern_update_info *model.InternUpdateInfo) error
}

type updateInternBuisiness struct {
	store updateInternStorage
}

func NewUpdateInternBusiness(store updateInternStorage) *updateInternBuisiness {
	return &updateInternBuisiness{store: store}
}

func (biz *updateInternBuisiness) UpdateInternBiz(int_id string, intern_update_info *model.InternUpdateInfo) error {
	if err_query := biz.store.UpdateIntern(int_id, intern_update_info); err_query != nil {
		return err_query
	}
	return nil
}
