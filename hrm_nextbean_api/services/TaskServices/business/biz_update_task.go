package business

import "github.com/PhuPhuoc/hrm_nextbean_api/services/TaskServices/model"

type updateTaskStorage interface {
	UpdateTask(proid, taskid string, info *model.TaskUpdate) error
}

type updateTaskBiz struct {
	store updateTaskStorage
}

func NewUpdateTaskBiz(store updateTaskStorage) *updateTaskBiz {
	return &updateTaskBiz{store: store}
}

func (biz *updateTaskBiz) UpdateTaskBiz(proid, taskid string, info *model.TaskUpdate) error {
	if err_query := biz.store.UpdateTask(proid, taskid, info); err_query != nil {
		return err_query
	}
	return nil
}
