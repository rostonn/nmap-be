package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/rostonn/nmap-be/dal"
)

func (a *App) uploadNmapResultsXmlFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println(formatRequest(r))
	var Buf bytes.Buffer

	file, header, err := r.FormFile("file")
	fmt.Println("File Received...")
	fmt.Println(header)
	fmt.Println(err)
	defer file.Close()

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	io.Copy(&Buf, file)

	contents := Buf.String()
	fmt.Println("File Uploaded...")
	fmt.Println(contents)

}

func (a *App) getNmapResultsByIpAddress(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil || body == nil {
		fmt.Println("Body is nil")
		respondWithError(w, 400, "Bad Request")
	}

	m := make(map[string]string)
	err = json.Unmarshal(body, &m)

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Request is bad json")
		respondWithError(w, 400, "Bad Request")
	}

	email, ok := m["ipAddress"]
	// email not in the request body
	if !ok {
		fmt.Println("ipAddress is not in bad json")
		respondWithError(w, 400, "Bad Request")
	}

	result, err := a.NmapDalService.GetNmapResultsByIp(a.DB, email)
	// Error while selecting results
	if err != nil {
		respondWithError(w, 500, "Server Error")
	}

	respondWithJSON(w, 200, result)
}

func (a *App) TestMock(nmapService dal.NmapServiceInterface) {
	result, err := nmapService.GetNmapResultsByIp(a.DB, "")
	fmt.Println(result)
	fmt.Println(err)
}
