package oauthgoogleservices

import (
	"github.com/PhuPhuoc/hrm_nextbean_api/services/OauthGoogleServices/model"
	"golang.org/x/oauth2"
)

func InitOauth(config oauth2.Config) *model.OauthApp {
	return &model.OauthApp{
		Conf: config,
	}
}
