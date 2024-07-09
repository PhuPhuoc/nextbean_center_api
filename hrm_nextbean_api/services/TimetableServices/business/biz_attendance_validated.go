package business

import "github.com/PhuPhuoc/hrm_nextbean_api/services/TimetableServices/model"

type attendanceValidatedStorage interface {
	AttendanceValidated(timetable_id string, info *model.AttendanceValidated) error
}

type attendanceValidatedBiz struct {
	store attendanceValidatedStorage
}

func NewAttendanceValidatedBiz(store attendanceValidatedStorage) *attendanceValidatedBiz {
	return &attendanceValidatedBiz{store: store}
}

func (biz *attendanceValidatedBiz) AttendanceValidatedBiz(timetable_id string, info *model.AttendanceValidated) error {
	if err_query := biz.store.AttendanceValidated(timetable_id, info); err_query != nil {
		return err_query
	}
	return nil
}
