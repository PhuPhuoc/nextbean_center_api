package business

import "github.com/PhuPhuoc/hrm_nextbean_api/services/InternSevices/model"

type createInternStorage interface {
	CreateIntern(intern_cre_info *model.InternCreation) error
}

type createInternBuisiness struct {
	store createInternStorage
}

func NewCreateAccountBusiness(store createInternStorage) *createInternBuisiness {
	return &createInternBuisiness{store: store}
}

func (biz *createInternBuisiness) CreateNewAccountBiz(intern_cre_info *model.InternCreation) error {
	if err_query := biz.store.CreateIntern(intern_cre_info); err_query != nil {
		return err_query
	}
	return nil
}
