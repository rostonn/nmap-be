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
	fmt.Println(string(data))
	var nmap Nmap
	xml.Unmarshal(data, &nmap)
	fmt.Println(nmap.Scanner)
	fmt.Println(nmap.Args)
	for i := 0; i < len(nmap.Hosts); i++ {
		// fmt.Println(nmap.Hosts[i])
		// fmt.Println(nmap.Hosts[i].Address)
		// fmt.Println(nmap.Hosts[i].Status)

		// if len(nmap.Hosts[i].Hostnames.Hostname) > 0 {
		// 	fmt.Println(nmap.Hosts[i].Hostnames.Hostname[0])
		// }
		fmt.Println(nmap.Hosts[i].Ports)
	}
}
