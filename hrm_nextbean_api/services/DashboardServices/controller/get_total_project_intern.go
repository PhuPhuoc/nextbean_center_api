package controller

import (
	"database/sql"
	"net/http"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/DashboardServices/business"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/DashboardServices/repository"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
)

//	@Summary		Get Total Project vs Total intern still in progress - Dashboard
//	@Description	Get a total number of project vs intern current in progress
//	@Tags			Dashboards
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	utils.success_response{data=model.DashboardTotalNumber}	"OK"
//	@Failure		400	{object}	utils.error_response									"Bad Request"
//	@Failure		404	{object}	utils.error_response									"Not Found"
//	@Router			/dashboards/total-number [get]
//	@Security		ApiKeyAuth
func handleDashboardGetTotalNumber(db *sql.DB) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {

		store := repository.NewDashboardStore(db)
		biz := business.NewGetTotalBiz(store)
		data, err := biz.GetTotalBiz()
		if err != nil {
			utils.WriteJSON(rw, utils.ErrorResponse_CannotGetEntity("dashboard total number", err))
			return
		}
		utils.WriteJSON(rw, utils.SuccessResponse_Data(data))
	}
}
