package controller

import (
	"net/http"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/OauthGoogleServices/model"
	"golang.org/x/oauth2"
)

func HandleOauth(a *model.OauthApp) func(rw http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		url := a.Conf.AuthCodeURL("nextbean-center", oauth2.AccessTypeOffline)
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}
}
