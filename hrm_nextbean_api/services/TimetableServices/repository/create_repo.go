package repository

import (
	"fmt"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/TimetableServices/model"
)

func (store *timetableStore) CreateTimetable(inid string, info *model.TimtableCreation) error {
	fmt.Println("info: ", info)
	return nil
}
