package controller

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/OauthGoogleServices/business"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/OauthGoogleServices/model"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/OauthGoogleServices/repository"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
	"golang.org/x/oauth2"
)

// @Summary		login by google email account
// @Description	login in using google email account
// @Tags			Authentication
// @Accept			json
// @Produce			json
// @Param			request	body		model.GoogleToken	 true	"Login request"
// @Success		200		{object}	utils.success_response	"Successful login"
// @Failure		400		{object}	utils.error_response	"login failure"
// @Router			/auth/login-google [post]
func HandleGoogleLogin(a *model.OauthApp, db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data_response := make(map[string]interface{})
		token := &oauth2.Token{}

		var body_data bytes.Buffer
		if _, err_read_body := body_data.ReadFrom(r.Body); err_read_body != nil {
			utils.WriteJSON(w, utils.ErrorResponse_InvalidRequest(err_read_body))
			return
		}
		json.Unmarshal(body_data.Bytes(), token)

		client := a.Conf.Client(context.Background(), token)
		resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
		if err != nil {
			utils.WriteJSON(w, utils.ErrorResponse_BadRequest("failed to get user info", err))
			return
		}
		defer resp.Body.Close()

		var userInfo map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
			utils.WriteJSON(w, utils.ErrorResponse_BadRequest("failed to decode user info", err))
			return
		}

		email, ok := userInfo["email"]
		if !ok {
			utils.WriteJSON(w, utils.ErrorResponse_BadRequest("failed to get user email", fmt.Errorf("failed to get user email")))
			return
		}

		store := repository.NewLoginGGStore(db)
		biz := business.NewLoginBusiness(store)
		if err_login := biz.Login(email.(string), data_response); err_login != nil {
			utils.WriteJSON(w, utils.ErrorResponse_BadRequest("cannot login", err_login))
			return
		}
		utils.WriteJSON(w, utils.SuccessResponse_Data(data_response))
	}
}
