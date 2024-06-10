package controller

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/InternServices/business"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/InternServices/model"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/InternServices/repository"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
)

// @Summary		update intern
// @Description	update intern's information
// @Tags			Intern
// @Accept			json
// @Produce		json
// @Param			request	body		model.InternUpdateInfo	true	"account creation request"
// @Success		200		{object}	utils.success_response	"Successful update"
// @Failure		400		{object}	utils.error_response	"update failure"
// @Router			/api/v1/intern [put]
func handleUpdateIntern(db *sql.DB) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		intern_info := new(model.InternUpdateInfo)
		var req_body_json map[string]interface{}

		var body_data bytes.Buffer
		if _, err_read_body := body_data.ReadFrom(req.Body); err_read_body != nil {
			utils.WriteJSON(rw, utils.ErrorResponse_InvalidRequest(err_read_body))
			return
		}
		json.Unmarshal(body_data.Bytes(), &req_body_json)

		check := utils.CreateValidateRequestBody(req_body_json, intern_info)
		if flag, list_err := check.GetValidateStatus(); !flag {
			utils.WriteJSON(rw, utils.ErrorResponse_BadRequest_ListError(list_err, fmt.Errorf("check request-body failed")))
			return
		}
		json.Unmarshal(body_data.Bytes(), intern_info)
		store := repository.NewInternStore(db)
		biz := business.NewUpdateInternBusiness(store)
		if err_biz := biz.UpdateInternBiz(intern_info); err_biz != nil {
			if strings.Contains(err_biz.Error(), "duplicate data") {
				utils.WriteJSON(rw, utils.ErrorResponse_BadRequest(err_biz.Error(), err_biz))
			} else {
				utils.WriteJSON(rw, utils.ErrorResponse_DB(err_biz))
			}
			return
		}
		utils.WriteJSON(rw, utils.SuccessResponse_MessageUpdated("intern updated successfully!"))
	}
}
