package business

import "github.com/PhuPhuoc/hrm_nextbean_api/services/TimetableServices/model"

type createTimetableStorage interface {
	CreateTimetable(inid string, info *model.TimtableCreation) error
}

type createTimetableBiz struct {
	store createTimetableStorage
}

func NewCreateTimetableBiz(store createTimetableStorage) *createTimetableBiz {
	return &createTimetableBiz{store: store}
}

func (biz *createTimetableBiz) CreateTimetabletBiz(inid string, info *model.TimtableCreation) error {
	if err_query := biz.store.CreateTimetable(inid, info); err_query != nil {
		return err_query
	}
	return nil
}
