package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/PhuPhuoc/hrm_nextbean_api/common"
	"github.com/PhuPhuoc/hrm_nextbean_api/middleware"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/CommentServices/business"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/CommentServices/model"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/CommentServices/repository"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
	"github.com/gorilla/mux"
)

// @Summary		Get report/comment in a task
// @Description	Get a list of report/comment in a task
// @Tags			Comments
// @Accept			json
// @Produce		json
// @Param			task-id	path		string											true	"enter task-id"
// @Param			page	query		int												false	"Page number"
// @Param			psize	query		int												false	"Number of records per page"
// @Param			type	query		string											false	"enter type for get report or comment"
// @Success		200		{object}	utils.success_response{data=[]model.Comment}	"OK"
// @Failure		400		{object}	utils.error_response							"Bad Request"
// @Failure		404		{object}	utils.error_response							"Not Found"
// @Router			/tasks/{task-id}/comments [get]
// @Security		ApiKeyAuth
func handleGetCommentInTask(db *sql.DB) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		var accID string
		if v := ctx.Value(middleware.AccIDKey); v != nil {
			accID = v.(string)
		}
		taskid := mux.Vars(req)["task-id"]
		if taskid == "" {
			utils.WriteJSON(rw, utils.ErrorResponse_BadRequest("Missing task's ID", fmt.Errorf("missing task's ID")))
			return
		}
		pagin := new(common.Pagination)
		filter := new(model.CommentFilter)
		getRequestQuery_comment(req, pagin, filter)
		filter.AccId = accID
		filter.TaskId = taskid

		store := repository.NewCommentStore(db)
		biz := business.NewGetCommentBiz(store)
		data, err := biz.GetCommentBiz(accID, pagin, filter)
		if err != nil {
			utils.WriteJSON(rw, utils.ErrorResponse_CannotGetEntity("comment", err))
			return
		}
		utils.WriteJSON(rw, utils.SuccessResponse_GetObject(pagin, filter, data))
	}
}

func getRequestQuery_comment(req *http.Request, pagin *common.Pagination, filter *model.CommentFilter) {
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
	filter.Type = req.URL.Query().Get("type")
}
