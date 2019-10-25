package dal

import (
	"database/sql"
	"fmt"
)

type NmapServiceInterface interface {
	GetNmapResultsByIp(db *sql.DB, ipAddress string) ([]NmapResultsForFrontend, error)
}

type NmapService struct{}

type NmapResultsForFrontend struct {
	IPAddress string `json:"ipAddress"`
	HostName  string `json:"hostName"`
	StartTs   string `json:"startTs"`
	EndTs     string `json:"endTs"`
	Port      int    `json:"port"`
	Status    string `json:"status"`
}

// getNmapResultsByIp - Query database based on ip address and return list of nmap scan data
func (n *NmapService) GetNmapResultsByIp(db *sql.DB, ipAddress string) ([]NmapResultsForFrontend, error) {

	query := `SELECT h.ip_address, h.host_name, h.start_ts, 
	h.end_ts, p.number, p.status FROM hosts h JOIN ports p ON h.host_id = p.host_id 
	WHERE h.ip_address = ?`

	nmapResults := []NmapResultsForFrontend{}

	rows, err := db.Query(query, ipAddress)
	if err != nil {
		// Log error if query failed
		fmt.Printf("Error querying getNmapResultsByIp %s", err.Error())
		return nmapResults, err
	}

	for rows.Next() {
		nmapRow := NmapResultsForFrontend{}
		err = rows.Scan(&nmapRow.IPAddress, &nmapRow.HostName,
			&nmapRow.StartTs, &nmapRow.EndTs, &nmapRow.Port, &nmapRow.Status)
		// Log error if scan failed
		if err != nil {
			fmt.Printf("Error parsing db row %s", err.Error())
		}
		nmapResults = append(nmapResults, nmapRow)
	}
	return nmapResults, nil
}
