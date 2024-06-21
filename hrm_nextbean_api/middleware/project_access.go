package middleware

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
	"github.com/gorilla/mux"
)

func ProjectAccessMiddleware(db *sql.DB, acceptManager, acceptPM, skipCheckPMInPro bool) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			accID := ctx.Value(AccIDKey).(string)
			accRole := ctx.Value(AccRoleKey).(string)

			switch accRole {
			case "admin":
				next.ServeHTTP(w, r.WithContext(ctx))
			case "manager":
				if acceptManager {
					next.ServeHTTP(w, r.WithContext(ctx))
				} else {
					utils.WriteJSON(w, utils.ErrorResponse_NoPermission("account's role (manager) is not allowed to access this api"))
					return
				}
			case "pm":
				if skipCheckPMInPro {
					next.ServeHTTP(w, r.WithContext(ctx))
				} else if acceptPM {
					proid := mux.Vars(r)["project-id"]
					if proid == "" {
						utils.WriteJSON(w, utils.ErrorResponse_BadRequest("Missing project ID", fmt.Errorf("missing project ID")))
						return
					}
					if err_pm := checkPMInProject(db, proid, accID); err_pm != nil {
						utils.WriteJSON(w, utils.ErrorResponse_NoPermission(err_pm.Error()))
						return
					}
					next.ServeHTTP(w, r.WithContext(ctx))
				} else {
					utils.WriteJSON(w, utils.ErrorResponse_NoPermission("account's role (pm) is not allowed to access this api"))
					return
				}
			}
		}
	}
}
