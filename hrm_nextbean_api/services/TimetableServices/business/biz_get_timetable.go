package business

import (
	"github.com/PhuPhuoc/hrm_nextbean_api/common"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/TimetableServices/model"
)

type getTimetableStore interface {
	GetTimetable(pagin *common.Pagination, filter *model.TimeTableFilter) ([]model.Timtable, error)
}

type getTimetableBusiness struct {
	store getTimetableStore
}

func NewGetTimetableBusiness(store getTimetableStore) *getTimetableBusiness {
	return &getTimetableBusiness{
		store: store,
	}
}

func (biz *getTimetableBusiness) GetTimetableBiz(pagin *common.Pagination, filter *model.TimeTableFilter) ([]model.Timtable, error) {
	data, err_query := biz.store.GetTimetable(pagin, filter)
	if err_query != nil {
		return nil, err_query
	}
	return data, nil
}
