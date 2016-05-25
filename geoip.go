/////////////
//
// GeoIP
// Freegeoip.net API client written in Go
//
// @Byt3smith
//
/////////////

package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

var BASE_URL = "http://freegeoip.net/json/"
type ResponseUserAgent struct {
	Useragent string `json:"user-agent"`
}

func geoQuery(url string) {

	// Struct to hold response data
	//
	type Response struct {
		IP		string `json:"ip"`
		CC		string `json:"country_code"`
		CN 		string `json:"country_name"`
		RN		string `json:"region_name"`
		CITY	string `json:"city"`
		LAT		float32 `json:"latitude"`
		LONG	float32 `json:"longitude"`
	}

	var client http.Client

	resp, err := client.Get(url)
	defer resp.Body.Close()
	if err != nil {
		 fmt.Println("[ ERROR ] ", err)
	}

	//
	// Process response
	//
	fmt.Println("\n[ Response Status: ", resp.StatusCode, " ]")

	if resp.StatusCode == 200 { // OK
		 bodyBytes, err2 := ioutil.ReadAll(resp.Body)
		 if err2 != nil {
			 fmt.Println(err2)
		 }
		 bodyString := string(bodyBytes)
		 res := Response{}
		 json.Unmarshal([]byte(bodyString), &res)
		 fmt.Println("\n[ Data ]")
		 fmt.Println("IP: ", res.IP)
		 fmt.Println("Country Code: ", res.CC)
		 fmt.Println("Country Name: ", res.CN)
		 fmt.Println("Region: ", res.RN)
		 fmt.Println("City: ", res.CITY)
		 fmt.Println("Latitude: ", res.LAT)
		 fmt.Println("Longitude: ", res.LONG)

	} else { // Not OK
		fmt.Println("[ ERROR ] ", err)
	}
}

func main() {
	var userdomain string

	fmt.Printf("Query [google.com, 8.8.4.4]: ")
	fmt.Scan(&userdomain)
  url := BASE_URL + userdomain
	geoQuery(url)

}
