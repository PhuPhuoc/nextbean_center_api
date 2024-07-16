package controller

import (
	"database/sql"
	"net/http"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/DashboardServices/business"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/DashboardServices/repository"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
)

//	@Summary		Get ojt model - Dashboard
//	@Description	Get a ojt info in dashboard
//	@Tags			Dashboards
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	utils.success_response{data=[]model.DashboardOJTInProgress}	"OK"
//	@Failure		400	{object}	utils.error_response										"Bad Request"
//	@Failure		404	{object}	utils.error_response										"Not Found"
//	@Router			/dashboards/inprogress-ojt [get]
//	@Security		ApiKeyAuth
func handleDashboardGetInProgressOJT(db *sql.DB) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {

		store := repository.NewDashboardStore(db)
		biz := business.NewGetInProgressOJTBiz(store)
		data, err := biz.GetInProgressOJTBiz()
		if err != nil {
			utils.WriteJSON(rw, utils.ErrorResponse_CannotGetEntity("dashboard in progress ojt", err))
			return
		}
		utils.WriteJSON(rw, utils.SuccessResponse_Data(data))
	}
}

