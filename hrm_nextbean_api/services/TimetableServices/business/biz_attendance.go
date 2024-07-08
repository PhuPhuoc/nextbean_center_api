package business

import "github.com/PhuPhuoc/hrm_nextbean_api/services/TimetableServices/model"

type attendanceTimetableStorage interface {
	AttendanceTimetable(inid, tid string, info *model.Attendance) error
}

type attendanceTimetableBiz struct {
	store attendanceTimetableStorage
}

func NewAttendanceTimetableBiz(store attendanceTimetableStorage) *attendanceTimetableBiz {
	return &attendanceTimetableBiz{store: store}
}

func (biz *attendanceTimetableBiz) AttendanceTimetabletBiz(inid, tid string, info *model.Attendance) error {
	if err_query := biz.store.AttendanceTimetable(inid, tid, info); err_query != nil {
		return err_query
	}
	return nil
}
