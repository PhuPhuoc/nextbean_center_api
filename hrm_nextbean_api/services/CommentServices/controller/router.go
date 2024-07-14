package controller

import (
	"database/sql"

	"github.com/PhuPhuoc/hrm_nextbean_api/middleware"
	"github.com/gorilla/mux"
)

func RegisterCommentRouter(r *mux.Router, db *sql.DB) {
	comment_router := r.PathPrefix("/tasks/{task-id}/comments").Subrouter()
	comment_router.Use(middleware.AuthMiddleware(db))
	comment_router.HandleFunc("", middleware.CommentAccessMiddleware(db, true, true, true)(handleCreateComment(db))).Methods("POST")
}
