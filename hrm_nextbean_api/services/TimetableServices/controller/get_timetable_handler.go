package controller

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/PhuPhuoc/hrm_nextbean_api/common"
	"github.com/PhuPhuoc/hrm_nextbean_api/middleware"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/TimetableServices/business"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/TimetableServices/model"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/TimetableServices/repository"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
)

// @Summary		Get timetables
// @Description	Get a list of timetables with filtering, sorting, and pagination
// @Tags			Timetables
// @Accept			json
// @Produce		json
// @Param			page				query		int												false	"Page number"
// @Param			psize				query		int												false	"Number of records per page"
// @Param			id					query		int												false	"Filter by timetable ID"
// @Param			intern-name			query		string											false	"Filter by intern name"
// @Param			student-code		query		string											false	"Filter by student-code"
// @Param			verified			query		string											false	"Filter by verified ~ ex: denied-approved | processing-denied-approved"
// @Param			status-attendance	query		string											false	"Filter by status-attendance ~ ex: denied-approved | processing-denied-approved"
// @Param			office-time-from	query		string											false	"Filter by office-time from (YYYY-MM-DD) ~ ex:2024-05-29"
// @Param			office-time-to		query		string											false	"Filter by office-time to (YYYY-MM-DD)"
// @Param			order-by			query		string											false	"Order by field (created_at or name), prefix with - for descending order ~ Ex: user_name desc"
// @Success		200					{object}	utils.success_response{data=[]model.Timtable}	"OK"
// @Failure		400					{object}	utils.error_response							"Bad Request"
// @Failure		404					{object}	utils.error_response							"Not Found"
// @Router			/timetables [get]
// @Security		ApiKeyAuth
func handleGetTimetable(db *sql.DB) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		var role string
		if v := ctx.Value(middleware.AccRoleKey); v != nil {
			role = v.(string)
		}
		var accID string
		if v := ctx.Value(middleware.AccIDKey); v != nil {
			accID = v.(string)
		}
		pagin := new(common.Pagination)
		filter := new(model.TimeTableFilter)
		filter.Role = role
		filter.Account_id = accID
		getRequestQuery(req, pagin, filter)

		store := repository.NewTimeTableStore(db)
		biz := business.NewGetTimetableBusiness(store)
		data, err := biz.GetTimetableBiz(pagin, filter)
		if err != nil {
			utils.WriteJSON(rw, utils.ErrorResponse_CannotGetEntity("timetable", err))
			return
		}

		utils.WriteJSON(rw, utils.SuccessResponse_GetObject(pagin, filter, data))
	}
}

func getRequestQuery(req *http.Request, pagin *common.Pagination, filter *model.TimeTableFilter) {
	page, err := strconv.Atoi(req.URL.Query().Get("page"))
	if err != nil {
		pagin.Page = 1
	}
	psize, err := strconv.Atoi(req.URL.Query().Get("psize"))
	if err != nil {
		pagin.PSize = 10
	}
	pagin.Page = page
	pagin.PSize = psize
	pagin.Process()

	filter.Id = req.URL.Query().Get("id")
	if filter.Role != "user" {
		filter.InternName = req.URL.Query().Get("intern-name")
		filter.StudentCode = req.URL.Query().Get("student-code")
	}
	filter.Verified = req.URL.Query().Get("verified")
	filter.StatusAttendance = req.URL.Query().Get("status-attendance")
	filter.OfficeTimeFrom = req.URL.Query().Get("office-time-from")
	filter.OfficeTimeTo = req.URL.Query().Get("office-time-to")
	filter.OrderBy = req.URL.Query().Get("order-by")
}
