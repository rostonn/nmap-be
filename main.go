package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rostonn/nmap-be/config"
)

// type Configuration struct {
// 	DbUsername           string `json:"db_username"`
// 	DbPassword           string `json:"db_password"`
// 	Dbname               string `json:"db_name"`
// 	Port                 string `json:"port"`
// 	AmazonClientId       string `json:"amazonClientId"`
// }

func main() {
	// Load Data From config file
	configuration := config.Configuration{}
	// Get environment variable PROFILE to determine what config file to use
	env := strings.ToLower(os.Getenv(" "))
	var configFilename string
	if env == "" {
		configFilename = "config/config.json"
	} else {
		configFilename = "config/config." + env + ".json"
	}

	file, err := os.Open(configFilename)
	if err != nil {
		log.Fatal("Config Filename doesn't exist ", configFilename)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Fatal(err)
	}

	a := App{}
	a.Config = configuration

	fmt.Println(a.Config)
	a.Initialize()
	a.Run()
}

// 	database, _ := sql.Open("sqlite3", "./FoxIp.db")
// 	statement, _ := database.Prepare("INSERT INTO ip_records (ip_address) VALUES (?)")
// 	statement.Exec("hello world")

// 	fs := http.FileServer(http.Dir("static"))
// 	http.Handle("/", fs)

// 	http.ListenAndServe(":3000", nil)
// }

func formatRequest(r *http.Request) string {
	// Create return string
	var request []string
	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	// If this is a POST, add post data
	if r.Method == "POST" {
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())
	}
	// Return the request as a string
	return strings.Join(request, "\n")
}
