package controller

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/business"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/repository"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
	"github.com/gorilla/mux"
)

//	@Summary		Get Project details
//	@Description	Get project info and details (total member, total task, total task have been done, total task in progress v.v)
//	@Tags			Projects
//	@Accept			json
//	@Produce		json
//	@Param			project-id	path		string												true	"enter project-id"
//	@Success		200			{object}	utils.success_response{data=[]model.ProjectDetail}	"OK"
//	@Failure		400			{object}	utils.error_response								"Bad Request"
//	@Failure		404			{object}	utils.error_response								"Not Found"
//	@Router			/projects/{project-id}/detail [get]
//	@Security		ApiKeyAuth
func handleGetProjectDetail(db *sql.DB) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		proid := mux.Vars(req)["project-id"]
		if proid == "" {
			utils.WriteJSON(rw, utils.ErrorResponse_BadRequest("Missing project ID", fmt.Errorf("missing project ID")))
			return
		}

		store := repository.NewProjectStore(db)
		biz := business.NewGetProDetailBiz(store)
		data, err := biz.GetProDetailBiz(proid)
		if err != nil {

			utils.WriteJSON(rw, utils.ErrorResponse_CannotGetEntity("project's detail", err))

			return
		}

		utils.WriteJSON(rw, utils.SuccessResponse_Data(data))
	}
}
