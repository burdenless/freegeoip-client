/////////////
//
// GeoIP
// Freegeoip.net API tool written in Go
//
// Developed by: Byt3smith
//
/////////////

package main

import (
	"fmt"
	"encoding/json"
	"github.com/jmcvetta/napping"
)

var BASE_URL = "http://freegeoip.net/json/"

type ResponseUserAgent struct {
	Useragent string `json:"user-agent"`
}

// A Params is a map containing URL parameters.
type Params map[string]string


func main() {

	// Start Session
	s := napping.Session{}
	var userdomain string

	fmt.Printf("Domain:> ")
	fmt.Scan(&userdomain)
  url := BASE_URL + userdomain

	fmt.Println("--------------------------------------------------------------------------------")
	println("")

	res := ResponseUserAgent{}
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
    fmt.Println("Error in request")
	}

	// Prepare the JSON response to be parsed
	var r interface{}
	err = json.Unmarshal([]byte(resp.RawText()), &r)

	//
	// Process response
	//
	println("")
	fmt.Println("Response Code:", resp.Status())
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println("Header")
	fmt.Println(resp.HttpResponse().Header)
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println("Data")
	fmt.Println(r)
	println("")
}
