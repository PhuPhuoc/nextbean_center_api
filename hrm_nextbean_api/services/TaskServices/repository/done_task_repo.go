package repository

import "github.com/PhuPhuoc/hrm_nextbean_api/services/TaskServices/model"

func (store *taskStore) EndTask(proid, taskid, assigneeid string, info *model.DoneTask) error {
	return nil
}
