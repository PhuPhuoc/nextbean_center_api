package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/business"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/repository"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
	"github.com/gorilla/mux"
)

// @Summary		remove map project-member
// @Description	remove member to project information
// @Tags			Projects
// @Accept			json
// @Produce		json
// @Param			project-id	path		string					true	"Project ID"
// @Param			member-id	path		string					true	"Member ID"
// @Success		200			{object}	utils.success_response	"Successful mapping"
// @Failure		400			{object}	utils.error_response	"mapping failure"
// @Router			/projects/{project-id}/{member-id} [delete]
// @Security		ApiKeyAuth
func handleRemoveMemberInProject(db *sql.DB) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		proid := mux.Vars(req)["project-id"]
		if proid == "" {
			utils.WriteJSON(rw, utils.ErrorResponse_BadRequest("Missing project ID", fmt.Errorf("missing project ID")))
			return
		}
		memid := mux.Vars(req)["member-id"]
		if proid == "" {
			utils.WriteJSON(rw, utils.ErrorResponse_BadRequest("Missing member ID", fmt.Errorf("missing member ID")))
			return
		}

		store := repository.NewProjectStore(db)
		biz := business.NewRemoveProMemBiz(store)
		if err_biz := biz.RemoveProMemBiz(proid, memid); err_biz != nil {
			if strings.Contains(err_biz.Error(), "invalid-request") {
				utils.WriteJSON(rw, utils.ErrorResponse_BadRequest(err_biz.Error(), err_biz))
			} else {
				utils.WriteJSON(rw, utils.ErrorResponse_DB(err_biz))
			}
			return
		}
		utils.WriteJSON(rw, utils.SuccessResponse_MessageCreated("removed member to project successfully!"))
	}
}
