package business

import "github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/model"

type createProjectStorage interface {
	CreateProject(info *model.ProjectCreationInfo) error
}

type createProjectBiz struct {
	store createProjectStorage
}

func NewCreateProjectBiz(store createProjectStorage) *createProjectBiz {
	return &createProjectBiz{store: store}
}

func (biz *createProjectBiz) CreateNewProjectBiz(prj_cre_info *model.ProjectCreationInfo) error {
	if err_query := biz.store.CreateProject(prj_cre_info); err_query != nil {
		return err_query
	}
	return nil
}
