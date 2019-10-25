package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rostonn/nmap-be/dal"
	"github.com/rostonn/nmap-be/dto"
)

// Route to Upload NMAP XML Results and save to DB
func (a *App) uploadNmapResultsXmlFile(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil || body == nil {
		fmt.Println("Body is nil")
		respondWithError(w, 400, "Bad Request - request should contain xml file")
		return
	}
	// Unmarshall Request Body into nmap
	var nmap dto.Nmap
	err = xml.Unmarshal(body, &nmap)
	if err != nil {
		fmt.Printf("Error unmarshalling xml request %s", err.Error())
		respondWithError(w, 500, "XML Parsing Error")
		return
	}

	result, err := a.NmapRepository.InsertNmapResults(a.DB, nmap)
	if err != nil || result == false {
		respondWithError(w, 500, "INTERNAL SERVER ERROR")
		return
	}
	fmt.Println("NMAP Results saved succesfully")
	w.WriteHeader(200)
}

// Route to get NMAP Results By IP Address
func (a *App) getNmapResultsByIpAddress(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil || body == nil {
		fmt.Println("Body is nil")
		respondWithError(w, 400, "Bad Request")
		return
	}

	m := make(map[string]string)
	err = json.Unmarshal(body, &m)

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Request is bad json")
		respondWithError(w, 400, "Bad Request")
		return
	}

	email, ok := m["ipAddress"]
	// email not in the request body
	if !ok {
		fmt.Println("ipAddress is not in bad json")
		respondWithError(w, 400, "Bad Request")
		return
	}

	result, err := a.NmapDalService.GetNmapResultsByIp(a.DB, email)
	// Error while selecting results
	if err != nil {
		respondWithError(w, 500, "Server Error")
		return
	}
	respondWithJSON(w, 200, result)
}

func (a *App) TestMock(nmapService dal.NmapServiceInterface) {
	result, err := nmapService.GetNmapResultsByIp(a.DB, "")
	fmt.Println(result)
	fmt.Println(err)
}
