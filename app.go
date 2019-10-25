package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rostonn/nmap-be/config"
	"github.com/rostonn/nmap-be/dal"
)

type App struct {
	Router         *mux.Router
	DB             *sql.DB
	Config         config.Configuration
	NmapDalService dal.NmapServiceInterface
	NmapRepository dal.NmapRepository
}

func (a *App) Initialize() {
	var err error
	a.DB, err = sql.Open("sqlite3", a.Config.DBPath)
	if err != nil {
		log.Fatal(err)
	}
	// Ping DB to see if we're connected
	err = a.DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
	a.NmapDalService = &dal.NmapService{}
	a.NmapRepository = &dal.NmapRepositoryImpl{}

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	// Upload XML File of NMAP Results
	a.Router.HandleFunc("/up", a.uploadNmapResultsXmlFile).Methods("POST")

	// Return list of nmpa results by ip address
	a.Router.HandleFunc("/nmap-by-ip", a.getNmapResultsByIpAddress).Methods("POST")

	// Serve static uploader file app
	a.Router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./static/"))))
	// Serve react app

}

func (a *App) Run() {
	headersOk := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(a.Config.Port, handlers.CORS(headersOk, originsOk, methodsOk)(a.Router)))
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
