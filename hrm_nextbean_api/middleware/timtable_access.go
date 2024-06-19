package middleware

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
)

func TimetableAccessMiddleware(db *sql.DB, acceptAdmin, acceptMem bool) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			accRole := ctx.Value(AccRoleKey).(string)
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
					next.ServeHTTP(w, r.WithContext(ctx))
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
