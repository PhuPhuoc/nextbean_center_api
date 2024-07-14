package business

import "github.com/PhuPhuoc/hrm_nextbean_api/services/TimetableServices/model"

type getTodayTimetableStore interface {
	GetTodayTimetable(inid string) (model.Today, error)
}

type getTodayTimetableBusiness struct {
	store getTodayTimetableStore
}

func NewGetTodayTimetableBusiness(store getTodayTimetableStore) *getTodayTimetableBusiness {
	return &getTodayTimetableBusiness{
		store: store,
	}
}

func (biz *getTodayTimetableBusiness) GetTodayBiz(inid string) (model.Today, error) {
	data, err_query := biz.store.GetTodayTimetable(inid)
	if err_query != nil {
		return model.Today{}, err_query
	}
	return data, nil
}
