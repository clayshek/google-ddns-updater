package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

func main() {

	ddns_service := "domains.google.com"
	ddns_hostname := os.Getenv("GOOG_DDNS_HOSTNAME")
	ddns_username := os.Getenv("GOOG_DDNS_USERNAME")
	ddns_password := os.Getenv("GOOG_DDNS_PASSWORD")

	if ddns_hostname == "" || ddns_username == "" || ddns_password == "" {
		log.Printf("Please ensure GOOG_DDNS_HOSTNAME, GOOG_DDNS_USERNAME, and GOOG_DDNS_PASSWORD env vars are set")
		return
	}

	externalIP, errExtIP := GetExternalIp()
	if errExtIP != nil {
		log.Printf("Unable to get external IP: %v", errExtIP)
		return
	}

	hostnameIP, errDDNSIP := GetCurrentDDNSIp(ddns_hostname)
	if errDDNSIP != nil {
		log.Printf("Unable to resolve DDNS hostname IP: %v", errDDNSIP)
		return
	}

	if externalIP != hostnameIP {
		log.Printf("DNS UPDATE REQUIRED. External IP=%v. %v=%v.", externalIP, ddns_hostname, hostnameIP)

		reqstr := fmt.Sprintf("https://%v/nic/update?hostname=%v&myip=%v", ddns_service, ddns_hostname, externalIP)
		log.Printf("Submitting update request to %v", reqstr)

		err := UpdateDDNSRecord(reqstr, ddns_username, ddns_password)
		if err != nil {
			log.Printf("Failed to update: %v", err)
			return
		}

	} else {
		log.Printf("NO DNS UPDATE REQUIRED. External IP=%v. %v=%v.", externalIP, ddns_hostname, hostnameIP)
	}

}

func GetExternalIp() (string, error) {

	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		return "Unable to determine external IP: %v ", err
	}
	defer resp.Body.Close()
	ip, err := ioutil.ReadAll(resp.Body)
	return strings.TrimSpace(string(ip)), nil
}

func GetCurrentDDNSIp(fqdn string) (string, error) {

	ips, err := net.LookupIP(fqdn)
	if err != nil {
		return "Unable to perform IP lookup: %v", err
	}
	var curDDNSIP net.IPAddr
	for _, ip := range ips {
		if ip.To4() != nil {
			curDDNSIP.IP = ip
			break
		}
	}
	if curDDNSIP.IP == nil {
		log.Printf("No A record found for: %v", fqdn)
	}

	return curDDNSIP.String(), nil
}

func UpdateDDNSRecord(url, user, pw string) error {

	client := &http.Client{
		CheckRedirect: nil,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Unable to create update request: %v", err)
	}
	//Uncomment block below to debug http request
	//dump, err := httputil.DumpRequestOut(req, true)
	//log.Printf("%s", dump)

	req.SetBasicAuth(user, pw)

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Unable to execute update request: %v", err)
	}
	//Uncomment block below to debug http response
	//dump2, err := httputil.DumpResponse(resp, true)
	//log.Printf("%s", dump2)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Unable to read update response: %v", err)
	}

	log.Printf("Response from service: %v", string(body))
	log.Printf("Request status response: %v", resp.Status)

	return nil

}
