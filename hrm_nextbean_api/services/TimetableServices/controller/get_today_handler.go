package controller

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/PhuPhuoc/hrm_nextbean_api/middleware"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/TimetableServices/business"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/TimetableServices/repository"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
)

//	@Summary		Intern get today timetables
//	@Description	Intern check if there is any work scheduled today
//	@Tags			Timetables
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	utils.success_response{data=[]model.Today}	"OK"
//	@Failure		400	{object}	utils.error_response						"Bad Request"
//	@Failure		404	{object}	utils.error_response						"Not Found"
//	@Router			/timetables/today [get]
//	@Security		ApiKeyAuth
func handleGetTodayTimetable(db *sql.DB) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		var inid string
		if v := ctx.Value(middleware.InternIDKey); v != nil {
			inid = v.(string)
		}

		if inid == "" {
			utils.WriteJSON(rw, utils.ErrorResponse_BadRequest("Missing intern's ID", fmt.Errorf("missing intern's ID")))
			return
		}
		store := repository.NewTimeTableStore(db)
		biz := business.NewGetTodayTimetableBusiness(store)
		data, err := biz.GetTodayBiz(inid)
		if err != nil {
			utils.WriteJSON(rw, utils.ErrorResponse_CannotGetEntity("today timetable", err))
			return
		}
		utils.WriteJSON(rw, utils.SuccessResponse_Data(data))
	}
}
