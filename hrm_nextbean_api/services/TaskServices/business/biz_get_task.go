package business

import (
	"github.com/PhuPhuoc/hrm_nextbean_api/common"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/TaskServices/model"
)

type getTaskStore interface {
	GetTask(pagin *common.Pagination, filter *model.TaskFilter) ([]model.Task, error)
}

type getTaskBiz struct {
	store getTaskStore
}

func NewGetTaskBiz(store getTaskStore) *getTaskBiz {
	return &getTaskBiz{
		store: store,
	}
}

func (biz *getTaskBiz) GetTaskBiz(pagin *common.Pagination, filter *model.TaskFilter) ([]model.Task, error) {
	data, err_query := biz.store.GetTask(pagin, filter)
	if err_query != nil {
		return nil, err_query
	}
	return data, nil
}
