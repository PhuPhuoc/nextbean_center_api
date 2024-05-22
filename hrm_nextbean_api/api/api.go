package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	address string
	db      *sql.DB
}

func InitServer(addr string, db *sql.DB) *server {
	return &server{
		address: addr,
		db:      db,
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func (sv *server) RunApp() error {
	router := mux.NewRouter()
	router.HandleFunc("/hello", helloHandler).Methods("GET")
	log.Printf("#HRM-nextbean-api#|api| + Server is listening at port ... |%v|", sv.address)
	return http.ListenAndServe(sv.address, router)
}
