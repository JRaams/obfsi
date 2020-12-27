package network

import (
	"io/ioutil"
	"log"
	"net/http"
)

const userAgent = "obfsi_spider/1.0"

// Get performs a http GET request to an 'url'
func Get(url string) string {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Panicf("Error creating new HttpGetReq: %s", err.Error())
	}
	req.Header.Set("User-Agent", userAgent)

	res, err := client.Do(req)
	if err != nil {
		log.Panicf("Error making GET req: %s", err.Error())
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Panicf("Error reading response body: %s", err.Error())
	}
	return string(body)
}
