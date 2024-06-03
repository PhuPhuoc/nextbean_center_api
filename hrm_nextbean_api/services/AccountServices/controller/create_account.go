package controller

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/AccountServices/business"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/AccountServices/model"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/AccountServices/repository"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
)

func arrayToString(arr []string) string {
	return strings.Join(arr, " ~ ")
}

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
		var req_body_json map[string]interface{}

		var body_data bytes.Buffer
		if _, err_read_body := body_data.ReadFrom(req.Body); err_read_body != nil {
			utils.WriteJSON(rw, utils.ErrorResponse_InvalidRequest(err_read_body))
			return
		}
		json.Unmarshal(body_data.Bytes(), &req_body_json)

		check := utils.CreateValidateRequestBody(req_body_json, acc_info)
		if flag, list_err := check.GetValidateStatus(); !flag {
			utils.WriteJSON(rw, utils.ErrorResponse_BadRequest_ListError(list_err, fmt.Errorf("check request-body failed")))
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
