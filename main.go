package nso_restconf_lib

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func setUser(u *url.URL) {
	u.Scheme = "http"
	u.Host = "163.162.195.6:8080"
	u.User = url.UserPassword("admin", "admin")
}

func GetSwitchConfiguration(device string) ([]byte, int, string, error) {

	u := new(url.URL)
	setUser(u)

	var netClient = &http.Client{
		Timeout: time.Second * 20,
	}

	req, err := getDeviceConfig(u, device)

	//req, err := interfaceConfig(u, device)
	//req, err := routerConfig(u, device)
	//req, err := syncFrom(u, device)
	//config, err := generateStatic("191.0.0.0/8", "10.87.89.1")

	/*
			config, err := generateStatic("2001:425::/32", "2001:420:2cff:1204::1")
		if err != nil {
			return nil, 500, "", err
		}

			req, err := setRouterConfig(u, device, "static", config)

	*/

	if err != nil {
		return nil, 500, "", err
	}

	resp, err := netClient.Do(req)
	if err != nil {
		return nil, 500, "", err
	}
	defer resp.Body.Close()

	// Read JSON data
	/* 	data := new(Router)
	   	err = decodeJSON(data, resp.Body)
	   if err != nil {
		return nil, 500, "", err
	}
	   	fmt.Printf("%v\n", data)
	*/

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 500, "", err
	}
	return contents, resp.StatusCode, resp.Status, err
}

func SyncFrom(device string) ([]byte, int, string, error) {

	u := new(url.URL)
	setUser(u)

	var netClient = &http.Client{
		Timeout: time.Second * 20,
	}

	req, err := syncFrom(u, device)

	if err != nil {
		return nil, 500, "", err
	}

	resp, err := netClient.Do(req)
	if err != nil {
		return nil, 500, "", err
	}
	defer resp.Body.Close()

	// Read JSON data
	/* 	data := new(Router)
	   	err = decodeJSON(data, resp.Body)
	   if err != nil {
		return nil, 500, "", err
	}
	   	fmt.Printf("%v\n", data)
	*/

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 500, "", err
	}
	return contents, resp.StatusCode, resp.Status, err
}

func CreateVlan(vlanId string, device string) ([]byte, int, string, error) {

	u := new(url.URL)
	setUser(u)

	var netClient = &http.Client{
		Timeout: time.Second * 20,
	}

	req, err := createVlan(u, vlanId, device)

	if err != nil {
		return nil, 500, "", err
	}

	resp, err := netClient.Do(req)
	if err != nil {
		return nil, 500, "", err
	}
	defer resp.Body.Close()

	contents, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 500, "", err
	}
	return contents, resp.StatusCode, resp.Status, err
}

func GetVlan(vlanId string, device string) ([]byte, int, string, error) {

	u := new(url.URL)
	setUser(u)

	var netClient = &http.Client{
		Timeout: time.Second * 20,
	}

	req, err := getVlan(u, vlanId, device)

	if err != nil {
		return nil, 500, "", err
	}

	resp, err := netClient.Do(req)
	defer resp.Body.Close()

	contents, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 500, "", err
	}
	return contents, resp.StatusCode, resp.Status, err
}

func GetVlanConfiguration(vlanId string, device string) ([]byte, int, string, error) {

	u := new(url.URL)
	setUser(u)

	var netClient = &http.Client{
		Timeout: time.Second * 20,
	}

	req, err := getVlanConfiguration(u, vlanId, device)

	if err != nil {
		return nil, 500, "", err
	}

	resp, err := netClient.Do(req)
	if err != nil {
		return nil, 500, "", err
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 500, "", err
	}
	return contents, resp.StatusCode, resp.Status, err
}

func CreateVlanConfiguration(vlanId string, device string, ipaddress string, interfaces []string) ([]byte, int, string, error) {

	u := new(url.URL)
	setUser(u)

	var netClient = &http.Client{
		Timeout: time.Second * 20,
	}

	req, err := createVlanConfiguration(u, vlanId, device, ipaddress, interfaces)

	if err != nil {
		return nil, 500, "", err
	}

	resp, err := netClient.Do(req)
	if err != nil {
		return nil, 500, "", err
	}
	defer resp.Body.Close()

	// Read JSON data
	/* 	data := new(Router)
	   	err = decodeJSON(data, resp.Body)
	   if err != nil {
		return nil, 500, "", err
	}
	   	fmt.Printf("%v\n", data)
	*/

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 500, "", err
	}
	return contents, resp.StatusCode, resp.Status, err
}

func ModifyVlanConfiguration(vlanId string, device string, ipaddress string, interfaces []string) ([]byte, int, string, error) {

	u := new(url.URL)
	setUser(u)

	var netClient = &http.Client{
		Timeout: time.Second * 20,
	}

	req, err := modifyVlanConfiguration(u, vlanId, device, ipaddress, interfaces)

	if err != nil {
		return nil, 500, "", err
	}

	resp, err := netClient.Do(req)
	if err != nil {
		return nil, 500, "", err
	}
	defer resp.Body.Close()

	contents, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 500, "", err
	}
	return contents, resp.StatusCode, resp.Status, err
}

func DeleteVlan(vlanId string, device string) ([]byte, int, string, error) {

	u := new(url.URL)
	setUser(u)

	var netClient = &http.Client{
		Timeout: time.Second * 20,
	}

	req, err := deleteVlan(u, vlanId, device)

	if err != nil {
		return nil, 500, "", err
	}

	resp, err := netClient.Do(req)
	if err != nil {
		return nil, 500, "", err
	}
	defer resp.Body.Close()

	// Read JSON data
	/* 	data := new(Router)
	   	err = decodeJSON(data, resp.Body)
	   if err != nil {
		return nil, 500, "", err
	}
	   	fmt.Printf("%v\n", data)
	*/

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 500, "", err
	}
	return contents, resp.StatusCode, resp.Status, err
}

func DeleteVlanConfiguration(vlanId string, device string) ([]byte, int, string, error) {

	u := new(url.URL)
	setUser(u)

	var netClient = &http.Client{
		Timeout: time.Second * 20,
	}

	req, err := deleteVlanConfiguration(u, vlanId, device)

	if err != nil {
		return nil, 500, "", err
	}

	resp, err := netClient.Do(req)
	if err != nil {
		return nil, 500, "", err
	}
	defer resp.Body.Close()

	// Read JSON data
	/* 	data := new(Router)
	   	err = decodeJSON(data, resp.Body)
	   if err != nil {
		return nil, 500, "", err
	}
	   	fmt.Printf("%v\n", data)
	*/

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 500, "", err
	}
	return contents, resp.StatusCode, resp.Status, err
}
