package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/PhuPhuoc/hrm_nextbean_api/middleware"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/TaskServices/business"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/TaskServices/repository"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
	"github.com/gorilla/mux"
)

// @Summary		start my task (for member)
// @Description	start task with new information
// @Tags			Tasks
// @Accept			json
// @Produce		json
// @Param			project-id	path		string					true	"enter project-id"
// @Param			task-id		path		string					true	"enter task-id"
// @Success		201			{object}	utils.success_response	"Successful start my task"
// @Failure		400			{object}	utils.error_response	"start my task failure"
// @Router			/projects/{project-id}/tasks/{task-id}/start-task [put]
// @Security		ApiKeyAuth
func handleStartMyTask(db *sql.DB) func(rw http.ResponseWriter, req *http.Request) {
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

		store := repository.NewTaskStore(db)
		biz := business.NewStartTaskBiz(store)
		if err := biz.StartTaskBiz(proid, taskid, internIDWhoAssignedTheTask); err != nil {
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
