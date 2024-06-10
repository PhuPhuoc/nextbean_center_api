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

// @Summary		map intern-intern
// @Description	Add intern skills information
// @Tags			Intern
// @Accept			json
// @Produce		json
// @Param			request	body		model.MapInternSkill	true	"Required: Fill in the id (number) of the skills into this array"
// @Success		200		{object}	utils.success_response	"Successful mapping"
// @Failure		400		{object}	utils.error_response	"mapping failure"
// @Router			/api/v1/intern/skill [post]
func handleMapInternSkill(db *sql.DB) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		map_info := new(model.MapInternSkill)
		var req_body_json map[string]interface{}

		var body_data bytes.Buffer
		if _, err_read_body := body_data.ReadFrom(req.Body); err_read_body != nil {
			utils.WriteJSON(rw, utils.ErrorResponse_InvalidRequest(err_read_body))
			return
		}
		json.Unmarshal(body_data.Bytes(), &req_body_json)

		check := utils.CreateValidateRequestBody(req_body_json, map_info)
		if flag, list_err := check.GetValidateStatus(); !flag {
			utils.WriteJSON(rw, utils.ErrorResponse_BadRequest_ListError(list_err, fmt.Errorf("check request-body failed")))
			return
		}
		json.Unmarshal(body_data.Bytes(), map_info)

		store := repository.NewInternStore(db)
		biz := business.NewMapInternSkillBiz(store)
		if err_biz := biz.MapInternSkillBiz(map_info); err_biz != nil {
			if strings.Contains(err_biz.Error(), "not exist") || strings.Contains(err_biz.Error(), "same length") {
				utils.WriteJSON(rw, utils.ErrorResponse_BadRequest(err_biz.Error(), err_biz))
			} else {
				utils.WriteJSON(rw, utils.ErrorResponse_DB(err_biz))
			}
			return
		}
		utils.WriteJSON(rw, utils.SuccessResponse_MessageCreated("Add skill information to your internship resume successfully!"))
	}
}
