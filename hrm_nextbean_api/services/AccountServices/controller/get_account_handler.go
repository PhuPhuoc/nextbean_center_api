package controller

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/PhuPhuoc/hrm_nextbean_api/common"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/AccountServices/business"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/AccountServices/model"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/AccountServices/repository"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
)

// @Summary		Get accounts
// @Description	Get a list of accounts with filtering, sorting, and pagination
// @Tags			Accounts
// @Accept			json
// @Produce		json
// @Param			page			query		int												false	"Page number"
// @Param			psize			query		int												false	"Number of records per page"
// @Param			id				query		int												false	"Filter by account ID"
// @Param			username		query		string											false	"Filter by username"
// @Param			email			query		string											false	"Filter by email"
// @Param			role			query		string											false	"Filter by role ~ ex: admin-pm | admin-manager-pm"
// @Param			created-at-from	query		string											false	"Filter by creation date from (YYYY-MM-DD) ~ ex:2024-05-29"
// @Param			created-at-to	query		string											false	"Filter by creation date to (YYYY-MM-DD)"
// @Param			order-by		query		string											false	"Order by field (created_at or name), prefix with - for descending order ~ Ex: user_name desc"
// @Success		200				{object}	utils.success_response{data=[]model.Account}	"OK"
// @Failure		400				{object}	utils.error_response							"Bad Request"
// @Failure		404				{object}	utils.error_response							"Not Found"
// @Router			/accounts [get]
func handleGetAccount(db *sql.DB) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		pagin := new(common.Pagination)
		filter := new(model.AccountFilter)
		getRequestQuery(req, pagin, filter)

		store := repository.NewAccountStore(db)
		biz := business.NewGetAccountBusiness(store)
		data, err := biz.GetAccountBiz(pagin, filter)
		if err != nil {
			utils.WriteJSON(rw, utils.ErrorResponse_CannotGetEntity("account", err))
			return
		}

		utils.WriteJSON(rw, utils.SuccessResponse_GetObject(pagin, filter, data))
	}
}

func getRequestQuery(req *http.Request, pagin *common.Pagination, filter *model.AccountFilter) {
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

	filter.Id = req.URL.Query().Get("id")
	filter.UserName = req.URL.Query().Get("username")
	filter.Email = req.URL.Query().Get("email")
	filter.Role = req.URL.Query().Get("role")
	filter.CreatedAtFrom = req.URL.Query().Get("created-at-from")
	filter.CreatedAtTo = req.URL.Query().Get("created-at-to")
	filter.OrderBy = req.URL.Query().Get("order-by")
}
