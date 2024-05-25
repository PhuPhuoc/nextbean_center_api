package controller

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/AccountServices/business"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/AccountServices/model"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/AccountServices/repository"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
)

// @Summary		create new account
// @Description	account creation information
// @Tags			Account
// @Accept			json
// @Produce		json
// @Param			request	body		model.AccountCreationInfo	true	"account creation request"
// @Success		200		{object}	utils.success_response		"Successful create"
// @Failure		400		{object}	utils.error_response		"create failure"
// @Router			/api/v1/account [post]
func HandleCreateAccount(db *sql.DB) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		acc_info := new(model.AccountCreationInfo)

		var body_data bytes.Buffer
		if _, err_read_body := body_data.ReadFrom(req.Body); err_read_body != nil {
			utils.WriteJSON(rw, utils.ErrorResponse_InvalidRequest(err_read_body))
			return
		}
		json.Unmarshal(body_data.Bytes(), acc_info)

		store := repository.NewAccountStore(db)
		biz := business.NewCreateAccountBusiness(store)
		if err := biz.CreateNewAccount(acc_info); err != nil {
			if strings.Contains(err.Error(), "exists") {
				utils.WriteJSON(rw, utils.ErrorResponse_BadRequest(err.Error(), err))
			} else {
				utils.WriteJSON(rw, utils.ErrorResponse_DB(err))
			}
			return
		}
		utils.WriteJSON(rw, utils.SuccessResponse_Message("Account created successfully!"))
	}
}
