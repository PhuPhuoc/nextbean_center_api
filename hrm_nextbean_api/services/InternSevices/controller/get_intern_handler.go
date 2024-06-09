package controller

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/PhuPhuoc/hrm_nextbean_api/common"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/InternSevices/business"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/InternSevices/model"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/InternSevices/repository"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
)

// @Summary		Get interns
// @Description	Get a list of interns with filtering, sorting, and pagination
// @Tags			Intern
// @Accept			json
// @Produce		json
// @Param			page	query		int											false	"Page number"
// @Param			psize	query		int											false	"Number of records per page"
// @Param			request	body		model.InternFilter							false	"intern'filter option"
// @Success		200		{object}	utils.success_response{data=[]model.Intern}	"OK"
// @Failure		400		{object}	utils.error_response						"Bad Request"
// @Failure		404		{object}	utils.error_response						"Not Found"
// @Router			/api/v1/intern/get [post]
func HandleGetIntern(db *sql.DB) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		pagin := new(common.Pagination)
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
		filter := new(model.InternFilter)

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
