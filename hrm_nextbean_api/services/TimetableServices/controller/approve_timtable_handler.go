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

// @Summary		approve intern timetable to work offline in office
// @Description	admin approve intern'schedule
// @Tags			Timetables
// @Accept			json
// @Produce		json
// @Param			timetable-id	path		string					true	"enter timetable-id"
// @Param			request			body		model.ApproveTimetable	true	"timetable creation request"
// @Success		200				{object}	utils.success_response	"Successful create"
// @Failure		400				{object}	utils.error_response	"create failure"
// @Router			/timetables/{timetable-id}/approve [post]
// @Security		ApiKeyAuth
func handleApproveInternTimeTable(db *sql.DB) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		inid := mux.Vars(req)["timetable-id"]
		if inid == "" {
			utils.WriteJSON(rw, utils.ErrorResponse_BadRequest("Missing timetable's ID", fmt.Errorf("missing timetable's ID")))
			return
		}

		info := new(model.ApproveTimetable)
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
		biz := business.NewApproveTimetableBiz(store)
		if err := biz.ApproveTimetabletBiz(inid, info); err != nil {
			if strings.Contains(err.Error(), "invalid-request") {
				utils.WriteJSON(rw, utils.ErrorResponse_BadRequest(err.Error(), err))
			} else {
				utils.WriteJSON(rw, utils.ErrorResponse_DB(err))
			}
			return
		}
		mess := fmt.Sprintf("%s intern's request (timetable) successfully", info.Status)
		utils.WriteJSON(rw, utils.SuccessResponse_MessageCreated(mess))
	}
}
