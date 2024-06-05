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

// @Summary		login by account
// @Description	Log in using account with email and password
// @Tags			Authentication
// @Accept			json
// @Produce			json
// @Param			request	body		model.LoginForm	 true	"Login request"
// @Success		200		{object}	utils.success_response	"Successful login"
// @Failure		400		{object}	utils.error_response	"login failure"
// @Router			/api/v1/login [post]
func HandleLogin(db *sql.DB) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		info_login := new(model.LoginForm)
		data_response := make(map[string]interface{})

		var body_data bytes.Buffer
		if _, err_read_body := body_data.ReadFrom(req.Body); err_read_body != nil {
			utils.WriteJSON(rw, utils.ErrorResponse_InvalidRequest(err_read_body))
			return
		}
		json.Unmarshal(body_data.Bytes(), info_login)

		store := repository.NewAccountStore(db)
		biz := business.NewLoginBusiness(store)
		if err_login := biz.Login(info_login, data_response); err_login != nil {
			if err_login.Error() == "wrong pwd" {
				utils.WriteJSON(rw, utils.ErrorResponse_BadRequest("wrong password", nil))
			} else if strings.Contains(err_login.Error(), "not exists") {
				utils.WriteJSON(rw, utils.ErrorResponse_BadRequest(err_login.Error(), nil))
			} else {
				utils.WriteJSON(rw, utils.ErrorResponse_DB(err_login))
			}
			return
		}
		utils.WriteJSON(rw, utils.SuccessResponse_Data(data_response))
	}
}
