package business

import "github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/model"

type updateProjectStorage interface {
	UpdateProject(proid string, info *model.UpdateProjectInfo) error
}

type updateProjectBiz struct {
	store updateProjectStorage
}

func NewUpdateProjectBiz(store updateProjectStorage) *updateProjectBiz {
	return &updateProjectBiz{store: store}
}

func (biz *updateProjectBiz) UpdateProjectBiz(proid string, info *model.UpdateProjectInfo) error {
	if err_query := biz.store.UpdateProject(proid, info); err_query != nil {
		return err_query
	}
	return nil
}
