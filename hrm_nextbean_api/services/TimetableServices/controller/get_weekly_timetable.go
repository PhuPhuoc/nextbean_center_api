package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/TimetableServices/business"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/TimetableServices/repository"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
)

// @Summary		Get weekly timetables
// @Description	Get a list of timetables in week
// @Tags			Timetables
// @Accept			json
// @Produce		json
// @Param			date	query		string											false	"enter current-date (YYYY-MM-DD) ~ ex:2024-05-29"
// @Success		200		{object}	utils.success_response{data=[]model.Timtable}	"OK"
// @Failure		400		{object}	utils.error_response							"Bad Request"
// @Failure		404		{object}	utils.error_response							"Not Found"
// @Router			/timetables/weekly [get]
// @Security		ApiKeyAuth
func handleGetWeeklyTimetable(db *sql.DB) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		date := req.URL.Query().Get("date")
		if date == "" {
			now := time.Now()
			date = now.Format("2006-01-02")
		}

		if is_date := utils.IsValidDate(date); !is_date {
			utils.WriteJSON(rw, utils.ErrorResponse_BadRequest("Invalid date ~ date must follow format YYYY-MM-DD", fmt.Errorf("invalid date")))
			return
		}
		store := repository.NewTimeTableStore(db)
		biz := business.NewGetWeeklyTimetableBusiness(store)
		data, err := biz.GetWeeklyTimetableBiz(date)
		if err != nil {
			utils.WriteJSON(rw, utils.ErrorResponse_CannotGetEntity("weekly timetable", err))
			return
		}

		utils.WriteJSON(rw, utils.SuccessResponse_Data(data))
	}
}
