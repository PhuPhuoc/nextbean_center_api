package business

import "github.com/PhuPhuoc/hrm_nextbean_api/services/TaskServices/model"

type createTaskStorage interface {
	CreateTask(role, proid, internIDWhoRegisteredTheTask string, info *model.TaskCreation) error
}

type createTaskBiz struct {
	store createTaskStorage
}

func NewCreateTaskBiz(store createTaskStorage) *createTaskBiz {
	return &createTaskBiz{store: store}
}

func (biz *createTaskBiz) CreateTaskBiz(role, proid, internIDWhoRegisteredTheTask string, info *model.TaskCreation) error {
	if err_query := biz.store.CreateTask(role, proid, internIDWhoRegisteredTheTask, info); err_query != nil {
		return err_query
	}
	return nil
}
