package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/PhuPhuoc/hrm_nextbean_api/docs"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
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

	// for api docs - swagger
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	router.HandleFunc("/hello", helloHandler).Methods("GET")

	log.Printf("|api|\n $ Server is listening at port |%v|\n $ URL swagger: http://localhost%v/swagger/index.html", sv.address, sv.address)
	return http.ListenAndServe(sv.address, router)
}
