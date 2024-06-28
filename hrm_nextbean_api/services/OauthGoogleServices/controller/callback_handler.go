package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/OauthGoogleServices/model"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
)

func HandleCallback(a *model.OauthApp) func(rw http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")

		// Exchanging the code for an access token
		t, err := a.Conf.Exchange(context.Background(), code)
		if err != nil {
			utils.WriteJSON(w, utils.ErrorResponse_BadRequest("cannot login with google account", err))
			return
		}
		fmt.Println("google token: ", t)
		// utils.WriteJSON(w, t)
	}
}
