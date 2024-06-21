package controller

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/TaskServices/business"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/TaskServices/model"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/TaskServices/repository"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
	"github.com/gorilla/mux"
)

// @Summary		update  task
// @Description	update task with new information
// @Tags			Tasks
// @Accept			json
// @Produce		json
// @Param			project-id	path		string					true	"enter project-id"
// @Param			request		body		model.TaskUpdate		true	"task creation request"
// @Success		201			{object}	utils.success_response	"Successful update"
// @Failure		400			{object}	utils.error_response	"update failure"
// @Router			/projects/{project-id}/tasks [put]
// @Security		ApiKeyAuth
func handleUpdateTask(db *sql.DB) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		proid := mux.Vars(req)["project-id"]
		if proid == "" {
			utils.WriteJSON(rw, utils.ErrorResponse_BadRequest("Missing project's ID", fmt.Errorf("missing project's ID")))
			return
		}

		up_info := new(model.TaskUpdate)
		var req_body_json map[string]interface{}

		var body_data bytes.Buffer
		if _, err_read_body := body_data.ReadFrom(req.Body); err_read_body != nil {
			utils.WriteJSON(rw, utils.ErrorResponse_InvalidRequest(err_read_body))
			return
		}
		json.Unmarshal(body_data.Bytes(), &req_body_json)

		check := utils.CreateValidateRequestBody(req_body_json, up_info)
		if flag, list_err := check.GetValidateStatus(); !flag {
			utils.WriteJSON(rw, utils.ErrorResponse_BadRequest_ListError(list_err, fmt.Errorf("check request-body failed")))
			return
		}
		json.Unmarshal(body_data.Bytes(), up_info)

		store := repository.NewTaskStore(db)
		biz := business.NewUpdateTaskBiz(store)
		if err := biz.UpdateTaskBiz(proid, up_info); err != nil {
			if strings.Contains(err.Error(), "invalid-request") {
				utils.WriteJSON(rw, utils.ErrorResponse_BadRequest(err.Error(), err))
			} else {
				utils.WriteJSON(rw, utils.ErrorResponse_DB(err))
			}
			return
		}
		utils.WriteJSON(rw, utils.SuccessResponse_MessageCreated("task updated successfully!"))
	}
}
