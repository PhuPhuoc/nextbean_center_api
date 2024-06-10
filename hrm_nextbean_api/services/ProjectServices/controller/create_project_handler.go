package controller

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/ProjectServices/model"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
)

//	@Summary		create new project
//	@Description	project creation information
//	@Tags			Project
//	@Accept			json
//	@Produce		json
//	@Param			request	body		model.ProjectCreationInfo	true	"project creation request"
//	@Success		200		{object}	utils.success_response		"Successful create"
//	@Failure		400		{object}	utils.error_response		"create failure"
//	@Router			/api/v1/project [post]
func handleCreateProject(db *sql.DB) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		info := new(model.ProjectCreationInfo)
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

		utils.WriteJSON(rw, utils.SuccessResponse_MessageCreated("Project created successfully!"))
	}
}
