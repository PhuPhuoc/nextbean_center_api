package controller

import (
	"database/sql"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/OauthGoogleServices/model"
	"github.com/gorilla/mux"
)

func RegisterOauthGGRouter(r *mux.Router, db *sql.DB, a *model.OauthApp) {
	auth_router := r.PathPrefix("/auth").Subrouter()
	auth_router.HandleFunc("/get-token", HandleGetGoogleToken(a)).Methods("GET")
	auth_router.HandleFunc("/login-google", HandleGoogleLogin(a, db)).Methods("POST")
}
