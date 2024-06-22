package controller

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/PhuPhuoc/hrm_nextbean_api/middleware"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/TaskServices/business"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/TaskServices/model"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/TaskServices/repository"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
	"github.com/gorilla/mux"
)

// @Summary		done my task (for member)
// @Description	done task with new information
// @Tags			Tasks
// @Accept			json
// @Produce		json
// @Param			project-id	path		string					true	"enter project-id"
// @Param			task-id		path		string					true	"enter task-id"
// @Param			request		body		model.DoneTask			true	"info to done the task"
// @Success		201			{object}	utils.success_response	"Successful done my task"
// @Failure		400			{object}	utils.error_response	"done my task failure"
// @Router			/projects/{project-id}/tasks/{task-id}/done-task [put]
// @Security		ApiKeyAuth
func handleDoneMyTask(db *sql.DB) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		var role string
		var internIDWhoAssignedTheTask string
		if v := ctx.Value(middleware.AccRoleKey); v != nil {
			role = v.(string)
		}
		if role == "user" {
			if v := ctx.Value(middleware.InternIDKey); v != nil {
				internIDWhoAssignedTheTask = v.(string)
			}
		}
		proid := mux.Vars(req)["project-id"]
		if proid == "" {
			utils.WriteJSON(rw, utils.ErrorResponse_BadRequest("Missing project's ID", fmt.Errorf("missing project's ID")))
			return
		}

		taskid := mux.Vars(req)["task-id"]
		if taskid == "" {
			utils.WriteJSON(rw, utils.ErrorResponse_BadRequest("Missing task's ID", fmt.Errorf("missing task's ID")))
			return
		}

		info := new(model.DoneTask)
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

		store := repository.NewTaskStore(db)
		biz := business.NewEndTaskBiz(store)
		if err := biz.EndTaskBiz(proid, taskid, internIDWhoAssignedTheTask, info); err != nil {
			if strings.Contains(err.Error(), "invalid-request") {
				utils.WriteJSON(rw, utils.ErrorResponse_BadRequest(err.Error(), err))
			} else {
				utils.WriteJSON(rw, utils.ErrorResponse_DB(err))
			}
			return
		}
		utils.WriteJSON(rw, utils.SuccessResponse_MessageCreated("start my task successfully!"))
	}
}
