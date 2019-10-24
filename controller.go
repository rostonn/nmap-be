package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
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
	fmt.Println("Getting to Controller")
	respondWithJSON(w, 200, nil)
	fmt.Println("It works")
}
