package dto

import "encoding/xml"

// Nmap is used to load Nmap xml results file
type Nmap struct {
	XMLName xml.Name `xml:"nmaprun"`
	Scanner string   `xml:"scanner,attr"`
	StartTs string   `xml:"start,attr"`
	Args    string   `xml:"args,attr"`
	Hosts   []Host   `xml:"host"`
}

type Host struct {
	Address   Address   `xml:"address"`
	StartTime string    `xml:"starttime,attr"`
	EndTime   string    `xml:"endtime,attr"`
	Status    Status    `xml:"status" json:"status"`
	Hostnames Hostnames `xml:"hostnames" json:"hostnames"`
	Ports     Ports     `xml:"ports" json:"ports"`
}

type Address struct {
	Ip   string `xml:"addr,attr"`
	Type string `xml:"addrtype,attr"`
}

//       <status state="up" reason="user-set" reason_ttl="0" />
type Status struct {
	State  string `xml:"state,attr" json:"state"`
	Reason string `xml:"reason,attr" json:"reason"`
}

type Hostnames struct {
	Hostname []Hostname `xml:"hostname" json:"hostname"`
}

//          <hostname name="cpc123026-glen5-2-0-cust970.2-1.cable.virginm.net" type="PTR" />
type Hostname struct {
	Name string `xml:"name,attr" json:"name"`
	Type string `xml:"PTR,attr" json:"type"`
}

type Ports struct {
	Port []Port `xml:"port" json:"port"`
}

//          <port protocol="tcp" portid="443">
type Port struct {
	Protocol string  `xml:"protocol,attr" json:"protocol"`
	PortId   string  `xml:"portid,attr" json:"portid"`
	State    State   `xml:"state" json:"state"`
	Service  Service `xml:"service" json:"service"`
}

//             <state state="open" reason="syn-ack" reason_ttl="0" />
type State struct {
	State  string `xml:"state,attr" json:"state"`
	Reason string `xml:"reason,attr" json:"reason"`
}

//             <service name="https" method="table" conf="3" />
type Service struct {
	Name   string `xml:"name" json:"name"`
	Method string `xml:"method" json:"method"`
}
