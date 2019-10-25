package dal

import (
	"fmt"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

// Mock rows returned from DB, result should be array of nmap results
func TestGetNmapResultsByIpReturnsListOfNmapResults(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	// Mock rows returned from database
	// h.ip_address, h.host_name, h.start_ts, h.end_ts, p.number, p.status
	rows := sqlmock.NewRows([]string{"h.ip_address", "h.host_name", "h.start_ts", "h.end_ts", "p.number", "p.status"}).
		AddRow("1.2.3.4", "test.com", "123456", "543213", 443, "open").
		AddRow("4.5.6.7", "end.com", "7", "8", 8080, "closed")

	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	mock.ExpectBegin()

	result, err := GetNmapResultsByIp(db, "1")
	if len(result) != 2 {
		t.Errorf("Size should be 2")
	}

	if result[0].IPAddress != "1.2.3.4" {
		t.Errorf("Ip address should be 1.2.3.4 got %s ", result[0].IPAddress)
	}

	if err != nil {
		t.Errorf("Error should be nil on this result")
	}
}

func TestGetNmapResultsByIPReturnsError(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT").WillReturnError(fmt.Errorf("some error"))

	result, err := GetNmapResultsByIp(db, "1")
	if err != nil && len(result) != 0 {
		t.Errorf("Error should result in error and empty array")
	}

}
