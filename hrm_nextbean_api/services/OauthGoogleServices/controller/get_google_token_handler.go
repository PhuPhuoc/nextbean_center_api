package controller

import (
	"context"
	"net/http"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/OauthGoogleServices/model"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
	"golang.org/x/oauth2"
)

func HandleGetGoogleToken(a *model.OauthApp) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")

		if code == "" {
			url := a.Conf.AuthCodeURL("nextbean-center", oauth2.AccessTypeOffline)
			http.Redirect(w, r, url, http.StatusTemporaryRedirect)
			return
		}

		token, err := a.Conf.Exchange(context.Background(), code)
		if err != nil {
			utils.WriteJSON(w, utils.ErrorResponse_BadRequest("cannot login with google account (s1)", err))
			return
		}

		utils.WriteJSON(w, utils.SuccessResponse_Data(token))
	}
}
