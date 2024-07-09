package business

import "github.com/PhuPhuoc/hrm_nextbean_api/services/TimetableServices/model"

type makeAbsentStorage interface {
	SetStatusAbsent(timetable_id string, info *model.StatusAbsent) error
}

type makeAbsentBiz struct {
	store makeAbsentStorage
}

func NewMakeAbsentdBiz(store makeAbsentStorage) *makeAbsentBiz {
	return &makeAbsentBiz{store: store}
}

func (biz *makeAbsentBiz) MakeAbsentBiz(timetable_id string, info *model.StatusAbsent) error {
	if err_query := biz.store.SetStatusAbsent(timetable_id, info); err_query != nil {
		return err_query
	}
	return nil
}
