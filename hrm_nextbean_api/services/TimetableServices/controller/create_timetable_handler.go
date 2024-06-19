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
)

//	@Summary		create new intern timetable to work offline in office
//	@Description	timetable creation information
//	@Tags			Timetables
//	@Accept			json
//	@Produce		json
//	@Param			request	body		model.TimtableCreation	true	"timetable creation request"
//	@Success		200		{object}	utils.success_response	"Successful create"
//	@Failure		400		{object}	utils.error_response	"create failure"
//	@Router			/timetables [post]
// @Security		ApiKeyAuth
func handleCreateTimeTable(db *sql.DB) func(rw http.ResponseWriter, req *http.Request) {
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

		cre_info := new(model.TimtableCreation)
		var req_body_json map[string]interface{}

		var body_data bytes.Buffer
		if _, err_read_body := body_data.ReadFrom(req.Body); err_read_body != nil {
			utils.WriteJSON(rw, utils.ErrorResponse_InvalidRequest(err_read_body))
			return
		}
		json.Unmarshal(body_data.Bytes(), &req_body_json)

		check := utils.CreateValidateRequestBody(req_body_json, cre_info)
		if flag, list_err := check.GetValidateStatus(); !flag {
			utils.WriteJSON(rw, utils.ErrorResponse_BadRequest_ListError(list_err, fmt.Errorf("check request-body failed")))
			return
		}
		json.Unmarshal(body_data.Bytes(), cre_info)

		store := repository.NewTimeTableStore(db)
		biz := business.NewCreateTimetableBiz(store)
		if err := biz.CreateTimetabletBiz(inid, cre_info); err != nil {
			if strings.Contains(err.Error(), "invalid-request") {
				utils.WriteJSON(rw, utils.ErrorResponse_BadRequest(err.Error(), err))
			} else {
				utils.WriteJSON(rw, utils.ErrorResponse_DB(err))
			}
			return
		}
		utils.WriteJSON(rw, utils.SuccessResponse_MessageCreated("timetable created successfully!"))
	}
}
