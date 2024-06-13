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
)

// @Summary		create new ojt (on the job training)
// @Description	ojt creation information
// @Tags			OJTS
// @Accept			json
// @Produce		json
// @Param			request	body		model.OJTCreationInfo	true	"ojt creation request"
// @Success		200		{object}	utils.success_response	"Successful create"
// @Failure		400		{object}	utils.error_response	"create failure"
// @Router			/ojts [post]
func handleCreateOJT(db *sql.DB) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		info := new(model.OJTCreationInfo)
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
		biz := business.NewCreateOJTBiz(store)
		if err_biz := biz.CreateOJTBiz(info); err_biz != nil {
			utils.WriteJSON(rw, utils.ErrorResponse_DB(err_biz))
		}

		utils.WriteJSON(rw, utils.SuccessResponse_MessageCreated("OJT created successfully!"))
	}
}
