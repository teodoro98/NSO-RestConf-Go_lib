package nso_restconf_lib

import (
	"bytes"
	"net/http"
	"net/url"
	"strings"
)

func getDeviceConfig(u *url.URL, device string) (req *http.Request, err error) {
	// "http://163.162.195.6:8080/restconf/data/tailf-ncs:devices/device=Clarinet4-leaf2/config"

	u.Path = "restconf/data/tailf-ncs:devices/device=" + device + "/config"
	// All the details
	//q := u.Query()
	//q.Set("deep", "true")
	//u.RawQuery = q.Encode()
	// Request
	req, err = http.NewRequest("GET", u.String(), nil)
	req.Header.Add("Accept", "application/yang-data+json")
	return req, err
}

func getVlan(u *url.URL, vlanId string, device string) (req *http.Request, err error) {

	//http://163.162.195.6:8080/restconf//data/VLAN-creation:VLAN-creation=14,Clarinet4-leaf2

	u.Path = "restconf//data/VLAN-creation:VLAN-creation=" + vlanId + "," + device
	// All the details
	//q := u.Query()
	//q.Set("deep", "true")
	//u.RawQuery = q.Encode()
	// Request
	req, err = http.NewRequest("GET", u.String(), nil)
	req.Header.Add("Accept", "application/yang-data+json")
	return req, err
}

func getVlanConfiguration(u *url.URL, vlanId string, device string) (req *http.Request, err error) {

	//http://163.162.195.6:8080/restconf//data/VLAN-configuration:VLAN-configuration=14,Clarinet4-leaf2

	u.Path = "restconf//data/VLAN-configuration:VLAN-configuration=" + vlanId + "," + device
	// All the details
	//q := u.Query()
	//q.Set("deep", "true")
	//u.RawQuery = q.Encode()
	// Request
	req, err = http.NewRequest("GET", u.String(), nil)
	req.Header.Add("Accept", "application/yang-data+json")
	return req, err
}

func createVlan(u *url.URL, vlanId string, device string) (req *http.Request, err error) {

	//http://163.162.195.6:8080/restconf/data

	u.Path = "restconf/data"
	// Request

	vlancreation := []byte(`{"VLAN-creation:VLAN-creation": [
        {
            "vlan-id": "` + vlanId + `",
            "device": "` + device + `"
        }
    ]}`)

	//fmt.Printf("%s\n", vlancreation)

	req, err = http.NewRequest("POST", u.String(), bytes.NewBuffer(vlancreation))
	req.Header.Add("Content-Type", "application/yang-data+json")

	//req.Header.Add("Accept", "application/yang-data+json")
	return req, err
}

func createVlanConfiguration(u *url.URL, vlanId string, device string, ipaddress string, interfaces []string) (req *http.Request, err error) {

	//http://163.162.195.6:8080/restconf/data

	u.Path = "restconf/data"
	// Request

	vlanconfigurationstring := `{"VLAN-configuration:VLAN-configuration": [
        {
            "vlan-id": "` + vlanId + `",
            "device": "` + device + `",
			"ip-address": "` + ipaddress + `"
			`

	if len(interfaces) != 0 {
		vlanconfigurationstring += `,"ifaces" : [`

		vlanconfigurationstring += `"` + interfaces[0] + `"`

		for i := 1; i < len(interfaces); i++ {
			vlanconfigurationstring += `,"` + interfaces[1] + `"`
		}

		vlanconfigurationstring += `]`
	}

	vlanconfigurationstring += `
        }
    ]}`

	vlanconfiguration := []byte(vlanconfigurationstring)

	//fmt.Printf("%s\n", vlanconfiguration)

	req, err = http.NewRequest("POST", u.String(), bytes.NewBuffer(vlanconfiguration))
	req.Header.Add("Content-Type", "application/yang-data+json")

	//req.Header.Add("Accept", "application/yang-data+json")
	return req, err
}

func modifyVlanConfiguration(u *url.URL, vlanId string, device string, ipaddress string, interfaces []string) (req *http.Request, err error) {

	//http://163.162.195.6:8080/restconf/data/VLAN-configuration:VLAN-configuration=101,Clarinet4-leaf2

	u.Path = "restconf/data/VLAN-configuration:VLAN-configuration=" + vlanId + "," + device
	// Request

	vlanconfigurationstring := `{"VLAN-configuration:VLAN-configuration": [
        {
            "vlan-id": "` + vlanId + `",
            "device": "` + device + `",
			"ip-address": "` + ipaddress + `"
			`

	if len(interfaces) != 0 {
		vlanconfigurationstring += `,"ifaces" : [`

		vlanconfigurationstring += `"` + interfaces[0] + `"`

		for i := 1; i < len(interfaces); i++ {
			vlanconfigurationstring += `,"` + interfaces[i] + `"`
		}

		vlanconfigurationstring += `]`
	}

	vlanconfigurationstring += `
        }
    ]}`

	vlanconfiguration := []byte(vlanconfigurationstring)

	//fmt.Printf("%s\n", vlanconfiguration)

	req, err = http.NewRequest("PATCH", u.String(), bytes.NewBuffer(vlanconfiguration))
	req.Header.Add("Content-Type", "application/yang-data+json")

	//req.Header.Add("Accept", "application/yang-data+json")
	return req, err
}

func deleteVlan(u *url.URL, vlanId string, device string) (req *http.Request, err error) {

	//http://163.162.195.6:8080/restconf//data/VLAN-creation:VLAN-creation=14,Clarinet4-leaf2

	u.Path = "restconf//data/VLAN-creation:VLAN-creation=" + vlanId + "," + device
	// All the details
	//q := u.Query()
	//q.Set("deep", "true")
	//u.RawQuery = q.Encode()
	// Request
	req, err = http.NewRequest("DELETE", u.String(), nil)
	req.Header.Add("Accept", "application/yang-data+json")
	return req, err
}

func deleteVlanConfiguration(u *url.URL, vlanId string, device string) (req *http.Request, err error) {

	//http://163.162.195.6:8080/restconf//data/VLAN-configuration:VLAN-configuration=14,Clarinet4-leaf2

	u.Path = "restconf//data/VLAN-configuration:VLAN-configuration=" + vlanId + "," + device
	// All the details
	//q := u.Query()
	//q.Set("deep", "true")
	//u.RawQuery = q.Encode()
	// Request
	req, err = http.NewRequest("DELETE", u.String(), nil)
	req.Header.Add("Accept", "application/yang-data+json")
	return req, err
}

func interfaceConfig(u *url.URL, d string) (req *http.Request, err error) {
	u.Path = "restconf/data/tailf-ncs:devices/device=" + d + "/config/interface/TenGigE"
	// All the details
	q := u.Query()
	q.Set("deep", "true")
	u.RawQuery = q.Encode()
	// Request
	req, err = http.NewRequest("GET", u.String(), nil)
	req.Header.Add("Accept", "application/vnd.yang.collection+json")
	return req, err
}

func routerConfig(u *url.URL, d string) (req *http.Request, err error) {
	// static, isis, bgp, etc...
	p := "static"
	u.Path = "restconf/data/tailf-ncs:devices/device=" + d + "/config/router/" + p
	// All the details
	q := u.Query()
	q.Set("deep", "true")
	u.RawQuery = q.Encode()
	// Request
	req, err = http.NewRequest("GET", u.String(), nil)
	req.Header.Add("Accept", "application/vnd.yang.data+json")
	return req, err
}

func setRouterConfig(u *url.URL, d string, p string, c string) (req *http.Request, err error) {
	// PUT, POST -> REPLACE. PATH -> MERGE
	// p = static, isis, bgp, etc...
	u.Path = "restconf/data/tailf-ncs:devices/device=" + d + "/config/router/" + p
	// Request
	req, err = http.NewRequest("PATCH", u.String(), strings.NewReader(c))
	req.Header.Add("Content-Type", "application/vnd.yang.data+json")
	req.Header.Add("Accept", "application/vnd.yang.data+json")
	return req, err
}

func syncFrom(u *url.URL, device string) (req *http.Request, err error) {

	//  http://163.162.195.6:8080/restconf/data/tailf-ncs:devices/device=Clarinet4-leaf2/sync-from

	u.Path = "restconf/data/tailf-ncs:devices/device=" + device + "/sync-from"
	// Request
	req, err = http.NewRequest("POST", u.String(), nil)
	req.Header.Add("Content-Type", "application/vnd.yang.operation+json")
	req.Header.Add("Accept", "application/vnd.yang.operation+json")
	return req, err
}
