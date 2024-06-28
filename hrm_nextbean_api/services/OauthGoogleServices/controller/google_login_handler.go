package controller

import (
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

func HandleGoogleLogin(a *model.OauthApp, db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data_response := make(map[string]interface{})
		code := r.URL.Query().Get("code")

		if code == "" {
			// Nếu không có code, chuyển hướng đến trang đăng nhập Google
			url := a.Conf.AuthCodeURL("nextbean-center", oauth2.AccessTypeOffline)
			http.Redirect(w, r, url, http.StatusTemporaryRedirect)
			return
		}

		// Nếu có code, xử lý callback
		token, err := a.Conf.Exchange(context.Background(), code)
		if err != nil {
			utils.WriteJSON(w, utils.ErrorResponse_BadRequest("cannot login with google account (s1)", err))
			return
		}

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
