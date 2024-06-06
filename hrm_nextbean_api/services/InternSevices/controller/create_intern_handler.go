package controller

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/InternSevices/business"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/InternSevices/model"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/InternSevices/repository"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
)

// @Summary		create new intern-account
// @Description	intern creation information
// @Tags			Intern
// @Accept			json
// @Produce		json
// @Param			request	body		model.InternCreation	true	"Required: user-name, email, password, student-code"
// @Success		200		{object}	utils.success_response	"Successful create"
// @Failure		400		{object}	utils.error_response	"create failure"
// @Router			/api/v1/intern [post]
func HandleCreateIntern(db *sql.DB) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		intern_info := new(model.InternCreation)
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
		biz := business.NewCreateAccountBusiness(store)
		if err_create := biz.CreateNewAccountBiz(intern_info); err_create != nil {
			if strings.Contains(err_create.Error(), "duplicate data") {
				utils.WriteJSON(rw, utils.ErrorResponse_BadRequest("duplicate data, cannot create new intern", err_create))
				return
			}
			utils.WriteJSON(rw, utils.ErrorResponse_DB(err_create))
			return
		}

		utils.WriteJSON(rw, utils.SuccessResponse_MessageCreated("Intern created successfully!"))
	}
}
