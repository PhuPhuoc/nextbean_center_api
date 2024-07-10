package controller

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/ReportServices/business"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/ReportServices/repository"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
	"github.com/gorilla/mux"
)

//	@Summary		Download Excel file about intern'task report in an ojt
//	@Description	Endpoint return file Excel
//	@Tags			Reports
//	@Accept			json
//	@Produce		application/vnd.openxmlformats-officedocument.spreadsheetml.sheet
//	@Param			ojt-id	path		int	true	"OJT ID"
//	@Success		200		{file}		"Excel file"
//	@Failure		400		{object}	utils.error_response	"Download failure"
//	@Router			/reports/{ojt-id}/project-intern [get]
func handleGetProjectReport(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ojt_id := mux.Vars(r)["ojt-id"]
		if ojt_id == "" {
			utils.WriteJSON(w, utils.ErrorResponse_BadRequest("missing ojt'ID", fmt.Errorf("missing ojt'ID")))
			return
		}
		store := repository.NewReportStore(db)
		biz := business.NewGetProjectReportBiz(store)
		data, err := biz.GetProjectReportBiz(ojt_id)
		if err != nil {
			utils.WriteJSON(w, utils.ErrorResponse_CannotGetEntity("ojt", err))
			return
		}

		utils.WriteFileExcel(w, data)
	}
}
