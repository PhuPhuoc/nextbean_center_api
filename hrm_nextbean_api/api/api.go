package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	account_services_controller "github.com/PhuPhuoc/hrm_nextbean_api/services/AccountServices/controller"
	intern_services_controller "github.com/PhuPhuoc/hrm_nextbean_api/services/InternServices/controller"
	ojt_services_controller "github.com/PhuPhuoc/hrm_nextbean_api/services/OJTServices/controller"
	oauthgoogleservices "github.com/PhuPhuoc/hrm_nextbean_api/services/OauthGoogleServices"
	oauthgoogleservices_controller "github.com/PhuPhuoc/hrm_nextbean_api/services/OauthGoogleServices/controller"
	pro_services_controller "github.com/PhuPhuoc/hrm_nextbean_api/services/ProServices/controller"
	task_services_controller "github.com/PhuPhuoc/hrm_nextbean_api/services/TaskServices/controller"
	tech_services_controller "github.com/PhuPhuoc/hrm_nextbean_api/services/TechnicalServices/controller"
	timetable_services_controller "github.com/PhuPhuoc/hrm_nextbean_api/services/TimetableServices/controller"

	_ "github.com/PhuPhuoc/hrm_nextbean_api/docs"
	"github.com/PhuPhuoc/hrm_nextbean_api/middleware"
	"github.com/gorilla/handlers"
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

func (sv *server) RunApp() error {
	router := mux.NewRouter()

	// for api docs - swagger
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Welcome to the server: IS-NextBean-v1"})
	}).Methods("GET")

	subrouter := router.PathPrefix("/api/v1").Subrouter()
	subrouter.Use(middleware.LoggingMiddleware)

	/* auth-config */
	conf, err_oauth := oauthgoogleservices.NewOauthAppConfig()
	if err_oauth != nil {
		log.Println("error when config oauth: ", err_oauth)
	}
	a := oauthgoogleservices.InitOauth(*conf)
	oauthgoogleservices_controller.RegisterOauthGGRouter(subrouter, sv.db, a)

	// router: /login vs /accounts
	account_services_controller.RegisterAccountRouter(subrouter, sv.db)
	// router: /interns
	intern_services_controller.RegisterInterntRouter(subrouter, sv.db)
	// router: /ojts
	ojt_services_controller.RegisterOJTRouter(subrouter, sv.db)
	// router: /technicals
	tech_services_controller.RegisterTechnicalRouter(subrouter, sv.db)
	// router: /projects
	pro_services_controller.RegisterProRouter(subrouter, sv.db)
	// router: /timetables
	timetable_services_controller.RegisterTimetableRouter(subrouter, sv.db)
	// router: /tasks
	task_services_controller.RegisterTaskRouter(subrouter, sv.db)

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)
	log.Printf("|api|\n  -+-  Server is listening at port |%v|\n  -+-  URL swagger: http://localhost%v/swagger/index.html\n", sv.address, sv.address)
	return http.ListenAndServe(sv.address, corsHandler(router))
}
