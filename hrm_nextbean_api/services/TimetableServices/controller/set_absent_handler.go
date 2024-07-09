package controller

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/TimetableServices/business"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/TimetableServices/model"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/TimetableServices/repository"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
	"github.com/gorilla/mux"
)

//	@Summary		The admin marked the intern absent from timetable.
//	@Description	Admins will mark intern absent when they don't see the student updating their work hours - or they can also remove the absent status.
//	@Tags			Timetables
//	@Accept			json
//	@Produce		json
//	@Param			timetable-id	path		string					true	"enter timetable-id"
//	@Param			request			body		model.StatusAbsent		true	"clockin/clockout validated request"
//	@Success		200				{object}	utils.success_response	"Successful validate"
//	@Failure		400				{object}	utils.error_response	"validate failure"
//	@Router			/timetables/{timetable-id}/absent [patch]
//	@Security		ApiKeyAuth
func handleStatusAbsentInTimetable(db *sql.DB) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		inid := mux.Vars(req)["timetable-id"]
		if inid == "" {
			utils.WriteJSON(rw, utils.ErrorResponse_BadRequest("Missing timetable's ID", fmt.Errorf("missing timetable's ID")))
			return
		}

		info := new(model.StatusAbsent)
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
		biz := business.NewMakeAbsentdBiz(store)
		if err := biz.MakeAbsentBiz(inid, info); err != nil {
			if strings.Contains(err.Error(), "invalid-request") {
				utils.WriteJSON(rw, utils.ErrorResponse_BadRequest(err.Error(), err))
			} else {
				utils.WriteJSON(rw, utils.ErrorResponse_DB(err))
			}
			return
		}
		var mess string
		if info.Status == "absent" {
			mess = "attendance status: absent"
		} else {
			mess = "attendance status: not yet"
		}
		utils.WriteJSON(rw, utils.SuccessResponse_MessageCreated(mess))
	}
}
