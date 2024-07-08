package controller

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/PhuPhuoc/hrm_nextbean_api/middleware"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/TimetableServices/business"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/TimetableServices/model"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/TimetableServices/repository"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
	"github.com/gorilla/mux"
)

//	@Summary		intern declares actual working hours
//	@Description	actual intern's working hours
//	@Tags			Timetables
//	@Accept			json
//	@Produce		json
//	@Param			timetable-id	path		int					true	"enter timetable-id"
//	@Param			request			body		model.Attendance		true	"timetable creation request"
//	@Success		200				{object}	utils.success_response	"Successful create"
//	@Failure		400				{object}	utils.error_response	"create failure"
//	@Router			/timetables/{timetable-id}/attendance [patch]
//	@Security		ApiKeyAuth
func handleClockinClockout(db *sql.DB) func(rw http.ResponseWriter, req *http.Request) {
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

		tid := mux.Vars(req)["timetable-id"]
		if inid == "" {
			utils.WriteJSON(rw, utils.ErrorResponse_BadRequest("Missing timetable's ID", fmt.Errorf("missing timetable's ID")))
			return
		}

		info := new(model.Attendance)
		var req_body_json map[string]interface{}

		var body_data bytes.Buffer
		if _, err_read_body := body_data.ReadFrom(req.Body); err_read_body != nil {
			utils.WriteJSON(rw, utils.ErrorResponse_InvalidRequest(err_read_body))
			return
		}
		json.Unmarshal(body_data.Bytes(), &req_body_json)

		check := utils.CreateValidateRequestBody(req_body_json, info)
		if flag, list_err := check.GetValidateStatus(); !flag {
			utils.WriteJSON(rw, utils.ErrorResponse_BadRequest_ListError(list_err, fmt.Errorf("check request-body failed")))
			return
		}
		json.Unmarshal(body_data.Bytes(), info)

		store := repository.NewTimeTableStore(db)
		biz := business.NewAttendanceTimetableBiz(store)
		if err := biz.AttendanceTimetabletBiz(inid, tid, info); err != nil {
			if strings.Contains(err.Error(), "invalid-request") {
				utils.WriteJSON(rw, utils.ErrorResponse_BadRequest(err.Error(), err))
			} else {
				utils.WriteJSON(rw, utils.ErrorResponse_DB(err))
			}
			return
		}
		utils.WriteJSON(rw, utils.SuccessResponse_MessageCreated("attendance updated successfully"))
	}
}
