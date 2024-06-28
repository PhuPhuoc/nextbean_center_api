package oauthgoogleservices

import (
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func NewOauthAppConfig() (*oauth2.Config, error) {
	cl_id := utils.GetClientID()
	cl_secret := utils.GetClientSecret()
	url_callback := utils.GetURLCallback()

	conf := &oauth2.Config{
		ClientID:     cl_id,
		ClientSecret: cl_secret,
		RedirectURL:  url_callback,
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}
	return conf, nil
}
