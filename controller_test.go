package main

import (
	"database/sql"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/rostonn/nmap-be/dal"
	"github.com/stretchr/testify/mock"
)

var app = App{}

type NmapServiceMock struct {
	mock.Mock
}

func (nmock *NmapServiceMock) GetNmapResultsByIp(db *sql.DB, ipAddress string) ([]dal.NmapResultsForFrontend, error) {
	args := nmock.Called()
	return args.Get(0).([]dal.NmapResultsForFrontend), args.Error(1)
}

func TestGetNmapResultsByIpAddress(t *testing.T) {
	nmapServiceMock := &NmapServiceMock{}
	// Mock service
	app.NmapDalService = nmapServiceMock

	// Create mock results for GetNmapResultsByIp function
	nmapResults := []dal.NmapResultsForFrontend{}
	result := dal.NmapResultsForFrontend{IPAddress: "1.1.1.1"}
	nmapResults = append(nmapResults, result)
	nmapServiceMock.On("GetNmapResultsByIp").Return(nmapResults, nil)

	// Test Body for request
	body := strings.NewReader(`{"ipAddress": "1.2.3.4"}`)

	// Test Request
	request := httptest.NewRequest("POST", "/nmap-by-ip", body)
	// Test response recorder
	w := httptest.NewRecorder()
	// Call function
	app.getNmapResultsByIpAddress(w, request)

	resp := w.Result()
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Status should be 200
	if resp.StatusCode != 200 {
		t.Errorf("Response should be 200 got %d", resp.StatusCode)
	}

	// Response body should contain results
	if string(respBody) != `[{"ipAddress":"1.1.1.1","hostName":"","startTs":"","endTs":"","port":0,"status":""}]` {
		t.Errorf("Response body shoud contain results")
	}
}
