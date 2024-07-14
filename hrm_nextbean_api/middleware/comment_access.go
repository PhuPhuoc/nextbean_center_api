package middleware

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
	"github.com/gorilla/mux"
)

func CommentAccessMiddleware(db *sql.DB, acceptManager, acceptPM, acceptMem bool) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			accID := ctx.Value(AccIDKey).(string)
			accRole := ctx.Value(AccRoleKey).(string)
			var internID string
			if accRole == "user" {
				if v := ctx.Value(InternIDKey); v != nil {
					internID = ctx.Value(InternIDKey).(string)
				}
			}

			switch accRole {
			case "admin":
				utils.WriteJSON(w, utils.ErrorResponse_NoPermission("account's role (admin) is not allowed to access this api"))
				return
			case "manager":
				if acceptManager {
					next.ServeHTTP(w, r.WithContext(ctx))
				} else {
					utils.WriteJSON(w, utils.ErrorResponse_NoPermission("account's role (manager) is not allowed to access this api"))
					return
				}
			case "pm":
				if acceptPM {
					taskid := mux.Vars(r)["task-id"]
					if taskid == "" {
						utils.WriteJSON(w, utils.ErrorResponse_BadRequest("Missing task ID", fmt.Errorf("missing task ID")))
						return
					}
					if err_pm := checkPMInTask(db, taskid, accID); err_pm != nil {
						utils.WriteJSON(w, utils.ErrorResponse_NoPermission(err_pm.Error()))
						return
					}
					next.ServeHTTP(w, r.WithContext(ctx))
				} else {
					utils.WriteJSON(w, utils.ErrorResponse_NoPermission("account's role (pm) is not allowed to access this api"))
					return
				}
			case "user":
				if acceptMem {
					taskid := mux.Vars(r)["task-id"]
					if taskid == "" {
						utils.WriteJSON(w, utils.ErrorResponse_BadRequest("Missing task ID", fmt.Errorf("missing project ID")))
						return
					}
					if err_pm := checkMemInTask(db, taskid, internID); err_pm != nil {
						utils.WriteJSON(w, utils.ErrorResponse_NoPermission(err_pm.Error()))
						return
					}
					next.ServeHTTP(w, r.WithContext(ctx))
				} else {
					utils.WriteJSON(w, utils.ErrorResponse_NoPermission("account's role (user) is not allowed to access this api"))
					return
				}
			}
		}
	}
}
