package controller

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/PhuPhuoc/hrm_nextbean_api/common"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/TaskServices/business"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/TaskServices/model"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/TaskServices/repository"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
)

// @Summary		Get tasks
// @Description	Get a list of tasks in project with filtering, sorting, and pagination
// @Tags			Tasks
// @Accept			json
// @Produce		json
// @Param			project-id		path		string										true	"enter project-id"
// @Param			page			query		int											false	"Page number"
// @Param			psize			query		int											false	"Number of records per page"
// @Param			name			query		string										false	"Project's Name"
// @Param			status			query		string										false	"Project's Status"
// @Param			assginee-name	query		string										false	"get task that belong to this assignee'name"
// @Param			assginee-code	query		string										false	"get task that belong to this assignee'code"
// @Param			is-approved		query		string										false	"get tasks were approved or not -> enter true or false"
// @Success		200				{object}	utils.success_response{data=[]model.Task}	"OK"
// @Failure		400				{object}	utils.error_response						"Bad Request"
// @Failure		404				{object}	utils.error_response						"Not Found"
// @Router			/projects/{project-id}/tasks [get]
// @Security		ApiKeyAuth
func handleGetTask(db *sql.DB) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		pagin := new(common.Pagination)
		filter := new(model.TaskFilter)

		getRequestQuery(req, pagin, filter)

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

func getRequestQuery(req *http.Request, pagin *common.Pagination, filter *model.TaskFilter) {
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

	filter.AssgineeName = req.URL.Query().Get("assginee-name")
	filter.AssgineeCode = req.URL.Query().Get("assginee-code")
	filter.Status = req.URL.Query().Get("status")
	filter.IsApproved = req.URL.Query().Get("is-approved")
	filter.Name = req.URL.Query().Get("name")
	filter.OrderBy = req.URL.Query().Get("order-by")
}
