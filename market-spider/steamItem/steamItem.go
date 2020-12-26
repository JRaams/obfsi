package steamItem

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/prprprus/scheduler"
)

const userAgent = "obfsi_spider/1.0"

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

type PriceJob struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type SteamResponse struct {
	Results []SteamItem `json:"results"`
}

type SteamItem struct {
	Name     string `json:"name"`
	Listings int    `json:"sell_listings"`
	Price    int    `json:"sell_price"`
}

func FetchPrices() {
	// Read price inputs
	input, err := ioutil.ReadFile("./steamItem/priceJobs.json")
	if err != nil {
		log.Panicf("Error reading req/priceJobs.json: %s", err.Error())
	}

	var priceJobs []PriceJob

	err = json.Unmarshal(input, &priceJobs)
	if err != nil {
		log.Panicf("Error unmarshaling price list input: %s", err.Error())
	}

	// Spread all jobs over 45 minutes to prevent spamming the market and getting timeouts
	s, err := scheduler.NewScheduler(100)
	if err != nil {
		log.Panicf("Error creating a scheduler: %s", err.Error())
	}
	delayPerJob := 45 / len(priceJobs)
	for i := 0; i < len(priceJobs); i++ {
		s.Delay().Minute(delayPerJob*i).Do(GetSteamMarketPrices, &priceJobs[i])
	}
	time.Sleep(50 * time.Minute)
}

func GetSteamMarketPrices(priceJob *PriceJob) []SteamItem {
	// log.Printf("<getSteamMarketPrices> name: %s, url: %s", priceJob.Name, priceJob.Url)
	// body := Get(priceJob.Url)
	// log.Printf("%s: %s", priceJob.Name, body)

	body, err := ioutil.ReadFile("./steamItem/test.json")
	if err != nil {
		log.Panicf("Error reading req/test.json: %s", err.Error())
	}

	var steamResponse SteamResponse

	err = json.Unmarshal(body, &steamResponse)
	if err != nil {
		log.Panicf("Error unmarshaling steamResponse: %s", err.Error())
	}

	return steamResponse.Results
}

func SaveSteamItems(db *sql.DB, steamItems []SteamItem) {
	log.Printf("Saving %d items to db...", len(steamItems))

	// for _, item := range steamItems {
	// 	// log.Printf("%v", item)
	// }
}
