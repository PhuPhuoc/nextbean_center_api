package middleware

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
	"github.com/gorilla/mux"
)

func ProjectAccessMiddleware(db *sql.DB, acceptManager, acceptPM bool) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			accID := ctx.Value(accIDKey).(string)
			accRole := ctx.Value(accRoleKey).(string)
			// internID := ctx.Value(internIDKey).(string)
			// fmt.Printf("pro middle ~ accID: %v accRole: %v internID: %v\n", accID, accRole, internID)

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
				if acceptPM {
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

func checkPMInProject(db *sql.DB, proID string, pmID string) error {
	var flag bool = false
	rawsql := `select exists(select 1 from project_manager pm join account acc on pm.account_id=acc.id where pm.project_id = ? and pm.account_id = ? and acc.deleted_at is null)`
	if err_query := db.QueryRow(rawsql, proID, pmID).Scan(&flag); err_query != nil {
		return err_query
	}
	if !flag {
		return fmt.Errorf("pm (id: %v) is not part of the project or the pm's account has been deleted", pmID)
	}
	return nil
}
