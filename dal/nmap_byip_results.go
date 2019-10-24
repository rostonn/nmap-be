package dal

import "database/sql"

type NmapResultsForFrontend struct {
	IPAddress string `json:"ipAddress"`
	HostName  string `json:"hostName"`
	StartTs   string `json:"startTs"`
	EndTs     string `json:"endTs"`
	Port      int    `json:"port"`
	Status    string `json:"status"`
}

type NmapTableRow struct {
}

func getNmapResultsByIp(db *sql.DB, ipAddress string) ([]NmapResultsForFrontend, error) {
	query := `SELECT h.ip_address, h.host_name, h.start_ts, 
	h.end_ts, p.number, p.status FROM hosts h JOIN ports p ON h.host_id = p.host_id 
	WHERE h.ip_address = ?`

	nmapResults := []NmapResultsForFrontend{}

	rows, err := db.Query(query, ipAddress)
	if err != nil {
		// Log error

		return nmapResults, err
	}

	for rows.Next() {
		nmapRow := NmapResultsForFrontend{}
		err = rows.Scan()
	}

}
