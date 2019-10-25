package dto

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"testing"
)

// Test parse xml file - nmap.results.xml
func TestParseXmlFile(t *testing.T) {
	fmt.Println("Test works")
	data, err := ioutil.ReadFile("nmap.results.xml")
	if err != nil {
		t.Errorf("File Error")
	}
	// Unmarshall xml into nmap
	var nmap Nmap
	xml.Unmarshal(data, &nmap)

	if nmap.Scanner != "nmap" {
		t.Errorf("Scanner should be nmap!")
	}

	// Check length of nmap scan
	if len(nmap.Hosts) != 40 {
		t.Errorf("Host length should be 40 got %d", len(nmap.Hosts))
	}
	// Verify hostname is parsing correctly
	if nmap.Hosts[0].Hostnames.Hostname[0].Name != "cpc123026-glen5-2-0-cust970.2-1.cable.virginm.net" {
		t.Errorf("Incorrect XML Parsing Hostname")
	}

	if nmap.Hosts[0].Address.Ip != "81.107.115.203" {
		t.Errorf("Incorrect XML Parsing IP address %s", nmap.Hosts[0].Address.Ip)
	}
}
