package business

import "github.com/PhuPhuoc/hrm_nextbean_api/services/InternServices/model"

type updateInternStorage interface {
	UpdateIntern(intern_update_info *model.InternUpdateInfo) error
}

type updateInternBuisiness struct {
	store updateInternStorage
}

func NewUpdateInternBusiness(store updateInternStorage) *updateInternBuisiness {
	return &updateInternBuisiness{store: store}
}

func (biz *updateInternBuisiness) UpdateInternBiz(intern_update_info *model.InternUpdateInfo) error {
	if err_query := biz.store.UpdateIntern(intern_update_info); err_query != nil {
		return err_query
	}
	return nil
}
