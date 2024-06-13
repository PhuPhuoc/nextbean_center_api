package controller

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func RegisterInterntRouter(r *mux.Router, db *sql.DB) {
	intern_router := r.PathPrefix("/interns").Subrouter()
	intern_router.HandleFunc("/{intern-id}", handleGetDetailIntern(db)).Methods("GET")
	intern_router.HandleFunc("", handleCreateIntern(db)).Methods("POST") //done
	intern_router.HandleFunc("", handleGetIntern(db)).Methods("GET") //done
	intern_router.HandleFunc("/{intern-id}", handleUpdateIntern(db)).Methods("PUT")
	intern_router.HandleFunc("/{intern-id}/skill", handleMapInternSkill(db)).Methods("POST") //done
}
