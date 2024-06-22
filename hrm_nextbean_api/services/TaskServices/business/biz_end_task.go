package business

import "github.com/PhuPhuoc/hrm_nextbean_api/services/TaskServices/model"

type endTaskStorage interface {
	EndTask(proid, taskid, assigneeid string, info *model.DoneTask) error
}

type endTaskBiz struct {
	store endTaskStorage
}

func NewEndTaskBiz(store endTaskStorage) *endTaskBiz {
	return &endTaskBiz{store: store}
}

func (biz *endTaskBiz) EndTaskBiz(proid, taskid, assigneeid string, info *model.DoneTask) error {
	if err_query := biz.store.EndTask(proid, taskid, assigneeid, info); err_query != nil {
		return err_query
	}
	return nil
}
