package controller

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/business"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/repository"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
	"github.com/gorilla/mux"
)

//	@Summary		Get PM
//	@Description	Get a list of PM in Project
//	@Tags			Project
//	@Accept			json
//	@Produce		json
//	@Param			project-id	path		string									false	"enter project-id"
//	@Success		200			{object}	utils.success_response{data=[]model.PM}	"OK"
//	@Failure		400			{object}	utils.error_response					"Bad Request"
//	@Failure		404			{object}	utils.error_response					"Not Found"
//	@Router			/api/v1/project/get-pm/{project-id} [get]
func handleGetPM(db *sql.DB) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		pro_id := vars["project-id"]

		store := repository.NewProjectStore(db)
		biz := business.NewGetPMBiz(store)
		data, err := biz.GetPMBiz(pro_id)
		if err != nil {
			if strings.Contains(err.Error(), "not exist") {
				utils.WriteJSON(rw, utils.ErrorResponse_BadRequest(err.Error(), err))
			} else {
				utils.WriteJSON(rw, utils.ErrorResponse_CannotGetEntity("pm", err))
			}
			return
		}

		utils.WriteJSON(rw, utils.SuccessResponse_Data(data))
	}
}
