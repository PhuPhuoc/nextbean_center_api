package business

import "github.com/PhuPhuoc/hrm_nextbean_api/services/TimetableServices/model"

type approveTimetableStorage interface {
	ApproveTimetable(timetable_id string, info *model.ApproveTimetable) error
}

type approveTimetableBiz struct {
	store approveTimetableStorage
}

func NewApproveTimetableBiz(store approveTimetableStorage) *approveTimetableBiz {
	return &approveTimetableBiz{store: store}
}

func (biz *approveTimetableBiz) ApproveTimetabletBiz(timetable_id string, info *model.ApproveTimetable) error {
	if err_query := biz.store.ApproveTimetable(timetable_id, info); err_query != nil {
		return err_query
	}
	return nil
}
