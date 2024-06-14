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

// @Summary		Get PM
// @Description	Get a list of PM in Project
// @Tags			Projects
// @Accept			json
// @Produce		json
// @Param			project-id	path		string									true	"enter project-id"
// @Success		200			{object}	utils.success_response{data=[]model.PM}	"OK"
// @Failure		400			{object}	utils.error_response					"Bad Request"
// @Failure		404			{object}	utils.error_response					"Not Found"
// @Router			/projects/{project-id}/pm-in-project [get]
func handleGetPM(db *sql.DB) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		proid := mux.Vars(req)["project-id"]
		if proid == "" {
			utils.WriteJSON(rw, utils.ErrorResponse_BadRequest("Missing project ID", fmt.Errorf("missing project ID")))
			return
		}

		store := repository.NewProjectStore(db)
		biz := business.NewGetPMBiz(store)
		data, err := biz.GetPMBiz(proid)
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
