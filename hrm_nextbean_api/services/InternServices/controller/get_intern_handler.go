package controller

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/PhuPhuoc/hrm_nextbean_api/common"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/InternServices/business"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/InternServices/model"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/InternServices/repository"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
)

//	@Summary		Get interns
//	@Description	Get a list of interns with filtering, sorting, and pagination
//	@Tags			Interns
//	@Accept			json
//	@Produce		json
//	@Param			page			query		int											false	"Page number"
//	@Param			psize			query		int											false	"Number of records per page"
//	@Param			account-id		query		string										false	"Account ID"
//	@Param			username		query		string										false	"Username"
//	@Param			email			query		string										false	"Email"
//	@Param			student-code	query		string										false	"Student Code"
//	@Param			ojt-semester	query		string										false	"OJT Semester"
//	@Param			gender			query		string										false	"Gender"
//	@Param			phone-number	query		string										false	"Phone Number"
//	@Param			address			query		string										false	"Address"
//	@Param			dob-from		query		string										false	"Date of Birth From"
//	@Param			dob-to			query		string										false	"Date of Birth To"
//	@Param			order-by		query		string										false	"Order by field (created_at or name), prefix with - for descending order ~ Ex: user_name desc"
//	@Success		200				{object}	utils.success_response{data=[]model.Intern}	"OK"
//	@Failure		400				{object}	utils.error_response						"Bad Request"
//	@Failure		404				{object}	utils.error_response						"Not Found"
//	@Router			/interns [get]
func handleGetIntern(db *sql.DB) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		pagin := new(common.Pagination)
		filter := new(model.InternFilter)
		getRequestQuery(req, pagin, filter)

		var body_data bytes.Buffer
		if _, err_read_body := body_data.ReadFrom(req.Body); err_read_body != nil {
			utils.WriteJSON(rw, utils.ErrorResponse_InvalidRequest(err_read_body))
			return
		}
		json.Unmarshal(body_data.Bytes(), filter)

		store := repository.NewInternStore(db)
		biz := business.NewGetInternBusiness(store)
		data, err := biz.GetInternBiz(pagin, filter)
		if err != nil {
			utils.WriteJSON(rw, utils.ErrorResponse_CannotGetEntity("intern", err))
			return
		}

		utils.WriteJSON(rw, utils.SuccessResponse_GetObject(pagin, filter, data))
	}
}

func getRequestQuery(req *http.Request, pagin *common.Pagination, filter *model.InternFilter) {
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

	filter.AccountID = req.URL.Query().Get("account-id")
	filter.UserName = req.URL.Query().Get("username")
	filter.Email = req.URL.Query().Get("email")
	filter.StudentCode = req.URL.Query().Get("student-code")
	filter.OJT_Semester = req.URL.Query().Get("ojt-semester")
	filter.Gender = req.URL.Query().Get("gender")
	filter.PhoneNumber = req.URL.Query().Get("phone-number")
	filter.Address = req.URL.Query().Get("address")
	filter.Dob_From = req.URL.Query().Get("dob-from")
	filter.Dob_To = req.URL.Query().Get("dob-to")
	filter.OrderBy = req.URL.Query().Get("order-by")
}
