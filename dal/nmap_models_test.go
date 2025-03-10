package dal

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/rostonn/nmap-be/dto"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

// Test nmap models functions using mock db and mock nmap from xml file
func TestMockDB(t *testing.T) {
	nmap := loadNMap()
	db, mock, err := sqlmock.New()
	mock.ExpectBegin()
	mock.ExpectExec("SELECT").WillReturnError(fmt.Errorf("some error"))
	mock.ExpectCommit()

	// Save Scan should return 0, error
	id, err := saveScan(db, nmap)
	if err != nil && id != 0 {
		t.Errorf("Expected error")
	}

	//	Insert Map should return false and error
	res, err := InsertNmapResults(db, nmap)
	if res != false && err != nil {
		t.Errorf("Scan save error should return error")
	}

	db, mock, err = sqlmock.New()
	// Insert returns id of 1 and 1 record that has been updated
	mock.ExpectExec("INSERT INTO scans").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	id, err = saveScan(db, nmap)
	if id != 1 && err != nil {
		t.Errorf("It should return the id 1 and nil")
	}
}

// Load NMap results for test
func loadNMap() dto.Nmap {
	data, err := ioutil.ReadFile("nmap_test_data.xml")
	if err != nil {
		return dto.Nmap{}
	}
	var nmap dto.Nmap
	xml.Unmarshal(data, &nmap)
	return nmap
}
