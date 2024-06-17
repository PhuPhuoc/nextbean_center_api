package middleware

import (
	"context"
	"database/sql"
	"net/http"
	"strings"

	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
)

func AuthMiddleware(db *sql.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				utils.WriteJSON(w, utils.ErrorResponse_Unauthorized())
				return
			}

			if !strings.HasPrefix(authHeader, "Bearer ") {
				utils.WriteJSON(w, utils.ErrorResponse_Unauthorized())
				return
			}

			token := strings.TrimPrefix(authHeader, "Bearer ")

			payload, err_verify_token := verifyJWTMiddleware(token)
			if err_verify_token != nil {
				if err_verify_token.Error() == "unauthorized" {
					utils.WriteJSON(w, utils.ErrorResponse_Unauthorized())
				} else if err_verify_token.Error() == "token_expired" {
					utils.WriteJSON(w, utils.ErrorResponse_TokenExpired())
				}
				return
			}

			accID, ok_id := payload["id"].(string)
			if ok_id {
				ctx = context.WithValue(ctx, accIDKey, accID)
			}
			accRole, ok_role := payload["role"].(string)
			if ok_role {
				ctx = context.WithValue(ctx, accRoleKey, accRole)
				if accRole == "user" {
					var internID *string
					query := `select id from intern where account_id=?`
					if err_query := db.QueryRow(query, accID).Scan(&internID); err_query != nil {
						utils.WriteJSON(w, utils.ErrorResponse_DB(err_query))
						return
					}
					ctx = context.WithValue(ctx, internIDKey, internID)
				}
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
