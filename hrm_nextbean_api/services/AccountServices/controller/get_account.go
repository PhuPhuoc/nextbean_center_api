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

func handleGetRequestParamQuery(req *http.Request, pagin *common.Pagination, order *string, filter *model.AccountFilter) {
	p, _ := strconv.Atoi(req.URL.Query().Get("page"))
	ps, _ := strconv.Atoi(req.URL.Query().Get("psize"))
	o := req.URL.Query().Get("order")

	if o == "" {
		o = "created_at desc"
	}

	pagin.Page = p
	pagin.PSize = ps
	pagin.Process()
	*order = o

	filter.Id = req.URL.Query().Get("id")
	filter.UserName = req.URL.Query().Get("username")
	filter.Email = req.URL.Query().Get("email")
	filter.Role = req.URL.Query().Get("role")
	filter.CreatedAtFrom = req.URL.Query().Get("created-at-from")
	filter.CreatedAtTo = req.URL.Query().Get("created-at-to")
}

// @Summary		Get accounts
// @Description	Get a list of accounts with filtering, sorting, and pagination
// @Tags			Account
// @Accept			json
// @Produce		json
// @Param			page			query		int												false	"Page number"
// @Param			psize			query		int												false	"Number of records per page"
// @Param			id				query		int												false	"Filter by account ID"
// @Param			username		query		string											false	"Filter by username"
// @Param			email			query		string											false	"Filter by email"
// @Param			role			query		string											false	"Filter by role"
// @Param			created-at-from	query		string											false	"Filter by creation date from (YYYY-MM-DD)"
// @Param			created-at-to	query		string											false	"Filter by creation date to (YYYY-MM-DD)"
// @Param			order			query		string											false	"Order by field (created_at or name), prefix with - for descending order"
// @Success		200				{object}	utils.success_response{data=[]model.Account}	"OK"
// @Failure		400				{object}	utils.error_response							"Bad Request"
// @Failure		404				{object}	utils.error_response							"Not Found"
// @Router			/api/v1/account [get]
func HandleGetAccount(db *sql.DB) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		var order string
		pagin := new(common.Pagination)
		filter := new(model.AccountFilter)
		handleGetRequestParamQuery(req, pagin, &order, filter)

		store := repository.NewAccountStore(db)
		biz := business.NewGetAccountBusiness(store)
		data, err := biz.GetAccountBiz(pagin, order, filter)
		if err != nil {
			utils.WriteJSON(rw, utils.ErrorResponse_CannotGetEntity("account", err))
			return
		}

		utils.WriteJSON(rw, utils.SuccessResponse_GetObject(order, pagin, filter, data))
	}
}
