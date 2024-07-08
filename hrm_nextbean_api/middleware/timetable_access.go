package middleware

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
	"github.com/gorilla/mux"
)

func TimetableAccessMiddleware(db *sql.DB, acceptAdmin, acceptMem, needCheckTimetableBelongToMem bool) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			accRole := ctx.Value(AccRoleKey).(string)
			inid := ""
			if accRole == "user" {
				inid = ctx.Value(InternIDKey).(string)
			}
			switch accRole {
			case "admin":
				if acceptAdmin {
					next.ServeHTTP(w, r.WithContext(ctx))
				} else {
					utils.WriteJSON(w, utils.ErrorResponse_NoPermission("account's role (admin) is not allowed to access this api"))
					return
				}
			case "user":
				if acceptMem {
					if needCheckTimetableBelongToMem {
						tid := mux.Vars(r)["timetable-id"]
						if tid == "" {
							utils.WriteJSON(w, utils.ErrorResponse_BadRequest("Missing project ID", fmt.Errorf("missing project ID")))
							return
						}
						if err_pm := checkInternIDBelongToTimeTable(db, tid, inid); err_pm != nil {
							utils.WriteJSON(w, utils.ErrorResponse_NoPermission(err_pm.Error()))
							return
						}
						next.ServeHTTP(w, r.WithContext(ctx))
					} else {
						next.ServeHTTP(w, r.WithContext(ctx))
					}
				} else {
					utils.WriteJSON(w, utils.ErrorResponse_NoPermission("account's role (user) is not allowed to access this api"))
					return
				}
			default:
				mess := fmt.Sprintf("account's role (%v) is not allowed to access this api", accRole)
				utils.WriteJSON(w, utils.ErrorResponse_NoPermission(mess))
				return
			}
		}
	}
}
