package middleware

import (
	"context"
	"log"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

type contextKey string

const (
	accRoleKey  contextKey = "role"
	accIDKey    contextKey = "accID"
	internIDKey contextKey = "internID"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrappedWriter := &responseWriter{ResponseWriter: w}
		// * request info
		method := r.Method
		req_uri := r.RequestURI
		start := time.Now()

		ctx := context.WithValue(r.Context(), accRoleKey, "unknown_role")
		ctx = context.WithValue(ctx, accIDKey, "unknown_acc_id")
		ctx = context.WithValue(ctx, internIDKey, "unknown_intern_id")

		// * Call the next handler
		next.ServeHTTP(wrappedWriter, r.WithContext(ctx))

		accRole, ok := r.Context().Value(accRoleKey).(string)
		if !ok {
			accRole = "still_unknown"
		}
		// * response info
		statusCode := wrappedWriter.statusCode
		execution_time := time.Since(start)
		log.Printf("{role: %s} -%s-   %s  ~  [stt/%v]  (exec_time: %v) \n", accRole, method, req_uri, statusCode, execution_time)
	})
}
