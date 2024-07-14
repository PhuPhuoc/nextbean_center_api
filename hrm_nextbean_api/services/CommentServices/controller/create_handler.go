package controller

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PhuPhuoc/hrm_nextbean_api/middleware"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/CommentServices/business"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/CommentServices/model"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/CommentServices/repository"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
	"github.com/gorilla/mux"
)

// @Summary		create new comment/report
// @Description	member vs manager/pm comment something in task or member can send daily report to manager/pm of project
// @Tags			Comments
// @Accept			json
// @Produce		json
// @Param			task-id	path		string					true	"enter task-id"
// @Param			request	body		model.CommentCreation	true	"comment creation request"
// @Success		200		{object}	utils.success_response	"Successful create"
// @Failure		400		{object}	utils.error_response	"create failure"
// @Router			/tasks/{task-id}/comments [post]
// @Security		ApiKeyAuth
func handleCreateComment(db *sql.DB) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		var accID string
		var role string
		var internIDInTask string
		if v := ctx.Value(middleware.AccRoleKey); v != nil {
			role = v.(string)
		}
		if v := ctx.Value(middleware.AccIDKey); v != nil {
			accID = v.(string)
		}
		if role == "user" {
			if v := ctx.Value(middleware.InternIDKey); v != nil {
				internIDInTask = v.(string)
			}
		}

		taskid := mux.Vars(req)["task-id"]
		if taskid == "" {
			utils.WriteJSON(rw, utils.ErrorResponse_BadRequest("Missing task's ID", fmt.Errorf("missing task's ID")))
			return
		}

		info := new(model.CommentCreation)
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

		store := repository.NewCommentStore(db)
		biz := business.NewCreateCommentBiz(store)
		if err := biz.CreateCommentBiz(taskid, role, accID, internIDInTask, info); err != nil {
			utils.WriteJSON(rw, utils.ErrorResponse_DB(err))
			return
		}
		mess := info.Type + " successfully!"
		utils.WriteJSON(rw, utils.SuccessResponse_MessageCreated(mess))
	}
}
