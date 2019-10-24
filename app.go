package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rostonn/nmap-be/config"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
	Config config.Configuration
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

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	// Upload XML File of NMAP Results
	a.Router.HandleFunc("/up", a.uploadNmapResultsXmlFile).Methods("POST")

	// Return list of nmpa results by ip address
	a.Router.HandleFunc("/nmap-by-ip", a.getNmapResultsByIpAddress).Methods("GET")

	a.Router.PathPrefix("/app/").Handler(http.StripPrefix("/app/", http.FileServer(http.Dir("./static/"))))
}

func (a *App) Run() {
	log.Fatal(http.ListenAndServe(a.Config.Port, a.Router))
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
