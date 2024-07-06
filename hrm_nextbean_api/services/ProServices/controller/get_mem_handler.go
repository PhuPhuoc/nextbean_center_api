package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PhuPhuoc/hrm_nextbean_api/common"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/business"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/model"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/repository"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
	"github.com/gorilla/mux"
)

// @Summary		Get Member in Project
// @Description	Get a list of Member in Project
// @Tags			Projects
// @Accept			json
// @Produce		json
// @Param			project-id		path		string										true	"enter project-id"
// @Param			page			query		int											false	"Page number"
// @Param			psize			query		int											false	"Number of records per page"
// @Param			user-name		query		string										false	"member'name"
// @Param			student-code	query		string										false	"student-code of member"
// @Param			semester		query		string										false	"member'semester"
// @Param			university		query		string										false	"member's university"
// @Param			status			query		string										false	"status of member in project"
// @Success		200				{object}	utils.success_response{data=[]model.Member}	"OK"
// @Failure		400				{object}	utils.error_response						"Bad Request"
// @Failure		404				{object}	utils.error_response						"Not Found"
// @Router			/projects/{project-id}/member-in-project [get]
// @Security		ApiKeyAuth
func handleGetMember(db *sql.DB) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		proid := mux.Vars(req)["project-id"]
		if proid == "" {
			utils.WriteJSON(rw, utils.ErrorResponse_BadRequest("Missing project ID", fmt.Errorf("missing project ID")))
			return
		}

		pagin := new(common.Pagination)
		filter := new(model.MemberFilter)
		getRequestQueryForMemberInProject(req, pagin, filter)

		store := repository.NewProjectStore(db)
		biz := business.NewGetProMemBiz(store)
		data, err := biz.GetProMemBiz(proid, pagin, filter)
		if err != nil {
			if strings.Contains(err.Error(), "invalid-request") {
				utils.WriteJSON(rw, utils.ErrorResponse_BadRequest(err.Error(), err))
			} else {
				utils.WriteJSON(rw, utils.ErrorResponse_CannotGetEntity("pm", err))
			}
			return
		}

		utils.WriteJSON(rw, utils.SuccessResponse_GetObject(pagin, filter, data))
	}
}

func getRequestQueryForMemberInProject(req *http.Request, pagin *common.Pagination, filter *model.MemberFilter) {
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

	filter.UserName = req.URL.Query().Get("user-name")
	filter.StudentCode = req.URL.Query().Get("student-code")
	filter.Semester = req.URL.Query().Get("semester")
	filter.University = req.URL.Query().Get("university")
	filter.Status = req.URL.Query().Get("status")
}
