package controller

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/InternServices/business"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/InternServices/repository"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
	"github.com/gorilla/mux"
)

// @Summary		Get intern'details by account-id
// @Description	Get details of intern (base infomation, skills, projects)
// @Tags			Interns
// @Accept			json
// @Produce		json
// @Param			intern-id	path		string												true	"enter intern-id"
// @Success		200			{object}	utils.success_response{data=model.InternDetailInfo}	"OK"
// @Failure		400			{object}	utils.error_response								"Bad Request"
// @Failure		404			{object}	utils.error_response								"Not Found"
// @Router			/interns/{intern-id} [get]
func handleGetDetailIntern(db *sql.DB) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		internID := vars["intern-id"]

		store := repository.NewInternStore(db)
		biz := business.NewGetDetailInternBiz(store)
		data, err_biz := biz.GetDetailInternBiz(internID)
		if err_biz != nil {
			if strings.Contains(err_biz.Error(), "not exists") {
				utils.WriteJSON(rw, utils.ErrorResponse_BadRequest(err_biz.Error(), err_biz))
			} else {
				utils.WriteJSON(rw, utils.ErrorResponse_CannotGetEntity("intern", err_biz))
			}
			return
		}
		utils.WriteJSON(rw, utils.SuccessResponse_Data(data))
	}
}
