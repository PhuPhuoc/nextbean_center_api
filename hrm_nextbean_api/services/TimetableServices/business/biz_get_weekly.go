package business

import "github.com/PhuPhuoc/hrm_nextbean_api/services/TimetableServices/model"

type getWeeklyTimetableStore interface {
	GetWeeklyTimetable(date string) ([7]model.Daily, error)
}

type getWeeklyTimetableBusiness struct {
	store getWeeklyTimetableStore
}

func NewGetWeeklyTimetableBusiness(store getWeeklyTimetableStore) *getWeeklyTimetableBusiness {
	return &getWeeklyTimetableBusiness{
		store: store,
	}
}

func (biz *getWeeklyTimetableBusiness) GetWeeklyTimetableBiz(date string) ([7]model.Daily, error) {
	data, err_query := biz.store.GetWeeklyTimetable(date)
	if err_query != nil {
		return [7]model.Daily{}, err_query
	}
	return data, nil
}
