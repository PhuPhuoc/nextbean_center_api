package controller

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/OJTServices/business"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/OJTServices/model"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/OJTServices/repository"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
	"github.com/gorilla/mux"
)

//	@Summary		update ojt
//	@Description	update ojt's information
//	@Tags			OJTS
//	@Accept			json
//	@Produce		json
//	@Param			ojt-id	path		int						true	"OJT ID"
//	@Param			request	body		model.UpdateOJTInfo		true	"OJT update request"
//	@Success		200		{object}	utils.success_response	"Successful update"
//	@Failure		400		{object}	utils.error_response	"update failure"
//	@Router			/ojts/{ojt-id} [put]
func handleUpdateOJT(db *sql.DB) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		ojt_id := mux.Vars(req)["ojt-id"]
		if ojt_id == "" {
			utils.WriteJSON(rw, utils.ErrorResponse_BadRequest("Missing account ID", fmt.Errorf("missing account ID")))
			return
		}
		info := new(model.UpdateOJTInfo)
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

		store := repository.NewOjtStore(db)
		biz := business.NewUpdateOJTBiz(store)
		if err_biz := biz.UpdateOJTBiz(ojt_id, info); err_biz != nil {
			utils.WriteJSON(rw, utils.ErrorResponse_DB(err_biz))
			return
		}

		utils.WriteJSON(rw, utils.SuccessResponse_MessageUpdated("OJT created successfully!"))
	}
}
