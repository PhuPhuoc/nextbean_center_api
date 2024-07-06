package controller

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/PhuPhuoc/hrm_nextbean_api/common"
	"github.com/PhuPhuoc/hrm_nextbean_api/middleware"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/business"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/model"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/repository"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
)

// @Summary		Get projects
// @Description	Get a list of projects with filtering, sorting, and pagination
// @Tags			Projects
// @Accept			json
// @Produce		json
// @Param			page						query		int												false	"Page number"
// @Param			psize						query		int												false	"Number of records per page"
// @Param			name						query		string											false	"Project's Name"
// @Param			status						query		string											false	"Project's Status"
// @Param			est-start-time-from			query		string											false	"get project which have start date from this date"
// @Param			est-start-time-to			query		string											false	"get project which have start date to this date"
// @Param			est-completion-time-from	query		string											false	"get project which have completed date from this date"
// @Param			est-completion-time-to		query		string											false	"get project which have completed date to this date"
// @Success		200							{object}	utils.success_response{data=[]model.Project}	"OK"
// @Failure		400							{object}	utils.error_response							"Bad Request"
// @Failure		404							{object}	utils.error_response							"Not Found"
// @Router			/projects [get]
// @Security		ApiKeyAuth
func handleGetProject(db *sql.DB) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		var role string
		if v := ctx.Value(middleware.AccRoleKey); v != nil {
			role = v.(string)
		}
		var accID string
		if v := ctx.Value(middleware.AccIDKey); v != nil {
			accID = v.(string)
		}

		pagin := new(common.Pagination)
		filter := new(model.ProjectFilter)

		getRequestQuery(req, pagin, filter)
		filter.Role = role
		filter.AccId = accID

		store := repository.NewProjectStore(db)
		biz := business.NewGetProBiz(store)
		data, err := biz.GetProBiz(pagin, filter)
		if err != nil {

			utils.WriteJSON(rw, utils.ErrorResponse_CannotGetEntity("project", err))
			return
		}

		utils.WriteJSON(rw, utils.SuccessResponse_GetObject(pagin, filter, data))
	}
}

func getRequestQuery(req *http.Request, pagin *common.Pagination, filter *model.ProjectFilter) {
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

	filter.Name = req.URL.Query().Get("name")
	filter.Status = req.URL.Query().Get("status")
	filter.EstStartTimeFrom = req.URL.Query().Get("est-start-time-from")
	filter.EstStartTimeTo = req.URL.Query().Get("est-start-time-to")
	filter.EstCompletionTimeFrom = req.URL.Query().Get("est-completion-time-from")
	filter.EstCompletionTimeTo = req.URL.Query().Get("est-completion-time-to")
	filter.OrderBy = req.URL.Query().Get("order-by")
}
