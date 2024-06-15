package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/PhuPhuoc/hrm_nextbean_api/common"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/OJTServices/business"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/OJTServices/model"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/OJTServices/repository"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
)

// @Summary		Get OJT
// @Description	Get a list of ojt with filtering, sorting, and pagination
// @Tags			OJTS
// @Accept			json
// @Produce		json
// @Param			page		query		int											false	"Page number"
// @Param			psize		query		int											false	"Number of records per page"
// @Param			id			query		int											false	"Filter by ojt'ID"
// @Param			semester	query		string										false	"Filter by semester"
// @Param			university	query		string										false	"Filter by university"
// @Param			order-by	query		string										false	"Order by field (created_at or name), prefix with - for descending order ~ Ex: university desc"
// @Success		200			{object}	utils.success_response{data=[]model.OJT}	"OK"
// @Failure		400			{object}	utils.error_response						"Bad Request"
// @Failure		404			{object}	utils.error_response						"Not Found"
// @Router			/ojts [get]
func handleGetOJT(db *sql.DB) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		pagin := new(common.Pagination)
		filter := new(model.FilterOJT)
		err := getRequestQuery(req, pagin, filter)
		if err != nil {
			utils.WriteJSON(rw, utils.ErrorResponse_BadRequest(err.Error(), err))
			return
		}

		store := repository.NewOjtStore(db)
		biz := business.NewGetOJTBiz(store)
		data, err := biz.GetOJTBiz(pagin, filter)
		if err != nil {
			utils.WriteJSON(rw, utils.ErrorResponse_CannotGetEntity("ojt", err))
			return
		}

		utils.WriteJSON(rw, utils.SuccessResponse_GetObject(pagin, filter, data))
	}
}

func getRequestQuery(req *http.Request, pagin *common.Pagination, filter *model.FilterOJT) error {
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

	ojt_id := req.URL.Query().Get("id")
	if ojt_id != "" {
		filter.Id, err = strconv.Atoi(ojt_id)
		if err != nil {
			return fmt.Errorf("ojt-id must be a number")
		}
	}
	filter.Semester = req.URL.Query().Get("semester")
	filter.University = req.URL.Query().Get("university")
	filter.OrderBy = req.URL.Query().Get("order-by")
	return nil
}
