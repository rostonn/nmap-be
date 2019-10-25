package dal

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/rostonn/nmap-be/dto"
)

type NmapRepository interface {
	InsertNmapResults(db *sql.DB, nmap dto.Nmap) (bool, error)
}

type NmapRepositoryImpl struct{}

type Host struct {
	HostID    int    `db:"host_id"`
	ScanID    int    `db:"scan_id"`
	Name      string `db:"name"`
	IPAddress string `db:"ip_address"`
	Status    string `db:"status"`
	StartTs   string `db:"start_ts"`
	EndTs     string `db:"end_ts"`
	Ports     []Port
}

type Port struct {
	PortID int    `db:"port_id"`
	HostID int    `db:"host_id"`
	Number int    `db:"number"`
	Status string `db:"status"`
	Reason string `db:"reason"`
}

// Insert Nmap XML Results
func (m *NmapRepositoryImpl) InsertNmapResults(db *sql.DB, nmap dto.Nmap) (bool, error) {
	// Save Scan First to see if this Nmap has already been saved
	scanId, err := saveScan(db, nmap)
	if err != nil {
		fmt.Printf("Scan has already been saved or error: %s %s", nmap.StartTs, err.Error())
		return false, err
	}

	dtoHosts := nmap.Hosts
	for _, dtoHost := range dtoHosts {
		// For each dtoHost, create a host
		host := Host{}
		host.ScanID = scanId
		host.setHostData(dtoHost)
		// Save host to DB and get id
		hostId, err := saveHost(db, host)
		// If we get an error saving to the database, log host info and continue to next host
		if err != nil {
			fmt.Printf("Failed Saving host to DB %s %s", host.Name, err.Error())
			continue
		}

		for _, dtoPort := range dtoHost.Ports.Port {
			portInt, err := strconv.Atoi(dtoPort.PortId)
			if err == nil {
				port := Port{}
				port.HostID = hostId
				port.setPortData(dtoPort, portInt)

				// Save Port to DB
				_, err := savePort(db, port)
				// Log Error and Continue
				if err != nil {
					fmt.Printf("Failed Saving Port HostId: %d Port %d Error %s", hostId, port.Number, err.Error())
				}
			}
		}
	}
	return true, nil
}

func saveScan(db *sql.DB, nmap dto.Nmap) (int, error) {
	statement := `INSERT INTO scans (start_ts) VALUES (?)`
	res, err := db.Exec(statement, nmap.StartTs)
	if err != nil {
		fmt.Println("Saving Scan Error")
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func saveHost(db *sql.DB, h Host) (int, error) {
	statement := `INSERT INTO hosts (host_name, ip_address, status, start_ts, end_ts) VALUES (?,?,?,?,?)`
	res, err := db.Exec(statement, h.Name, h.IPAddress, h.Status, h.StartTs, h.EndTs)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func savePort(db *sql.DB, p Port) (int, error) {

	statement := `INSERT INTO ports (host_id, number, status, reason) VALUES (?,?,?,?)`
	res, err := db.Exec(statement, p.HostID, p.Number, p.Status, p.Reason)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (p *Port) setPortData(dtoPort dto.Port, number int) {
	p.Number = number
	p.Status = dtoPort.State.State
	p.Reason = dtoPort.State.Reason
}

func (h *Host) setHostData(dtoHost dto.Host) {
	h.addHostNameIfItExists(dtoHost.Hostnames)
	h.IPAddress = dtoHost.Address.Ip
	h.Status = dtoHost.Status.Reason
	h.StartTs = dtoHost.StartTime
	h.EndTs = dtoHost.EndTime

}

func (h *Host) addHostNameIfItExists(dtoHost dto.Hostnames) {
	if len(dtoHost.Hostname) > 0 {
		h.Name = dtoHost.Hostname[0].Name
	}
}
