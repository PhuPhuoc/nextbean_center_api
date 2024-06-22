package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/PhuPhuoc/hrm_nextbean_api/common"
	"github.com/PhuPhuoc/hrm_nextbean_api/middleware"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/TaskServices/business"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/TaskServices/model"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/TaskServices/repository"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
)

// @Summary		Get my tasks (for member)
// @Description	Get a list of my tasks in project with filtering, sorting, and pagination
// @Tags			Tasks
// @Accept			json
// @Produce		json
// @Param			project-id	path		string										true	"enter project-id"
// @Param			page		query		int											false	"Page number"
// @Param			psize		query		int											false	"Number of records per page"
// @Success		200			{object}	utils.success_response{data=[]model.Task}	"OK"
// @Failure		400			{object}	utils.error_response						"Bad Request"
// @Failure		404			{object}	utils.error_response						"Not Found"
// @Router			/projects/{project-id}/tasks/my-task [get]
// @Security		ApiKeyAuth
func handleGetMyTask(db *sql.DB) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		var inid string
		if v := ctx.Value(middleware.InternIDKey); v != nil {
			inid = v.(string)
		}

		if inid == "" {
			utils.WriteJSON(rw, utils.ErrorResponse_BadRequest("Missing intern's ID", fmt.Errorf("missing intern's ID")))
			return
		}
		pagin := new(common.Pagination)
		filter := new(model.TaskFilter)

		getRequestQuery_mytask(req, pagin)
		filter.AssgineeId = inid

		store := repository.NewTaskStore(db)
		biz := business.NewGetTaskBiz(store)
		data, err := biz.GetTaskBiz(pagin, filter)
		if err != nil {
			utils.WriteJSON(rw, utils.ErrorResponse_CannotGetEntity("task", err))
			return
		}
		utils.WriteJSON(rw, utils.SuccessResponse_GetObject(pagin, filter, data))
	}
}

func getRequestQuery_mytask(req *http.Request, pagin *common.Pagination) {
	page, err := strconv.Atoi(req.URL.Query().Get("page"))
	if err != nil {
		pagin.Page = 1
	}
	psize, err := strconv.Atoi(req.URL.Query().Get("psize"))
	if err != nil {
		pagin.PSize = 10
	}
	pagin.Page = page
	pagin.PSize = psize
	pagin.Process()
}
