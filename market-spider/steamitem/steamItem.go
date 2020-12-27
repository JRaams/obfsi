package steamitem

import (
	"database/sql"
	"encoding/json"
	"log"
	"time"

	"github.com/jraams/obfsi/market-spider/network"
	"github.com/prprprus/scheduler"
)

// FetchPrices fetches the latest Steam market items for all jobs in priceJobs.json
func FetchPrices(db *sql.DB) {
	// Spread all jobs over 45 minutes to prevent spamming the market and getting timeouts
	s, err := scheduler.NewScheduler(100)
	if err != nil {
		log.Panicf("Error creating a scheduler: %s", err.Error())
	}
	delayPerJob := 45 / len(priceJobs)
	for i := 0; i < len(priceJobs); i++ {
		s.Delay().Second(delayPerJob*i).Do(getSteamMarketPrices, db, &priceJobs[i])
	}
	time.Sleep(50 * time.Minute)
}

func getSteamMarketPrices(db *sql.DB, priceJob *PriceJob) {
	log.Printf("<getSteamMarketPrices> Job: %s", priceJob.Name)
	body := network.Get(priceJob.URL)

	var steamResponse SteamResponse

	err := json.Unmarshal([]byte(body), &steamResponse)
	if err != nil {
		log.Panicf("Error unmarshaling steamResponse: %s", err.Error())
	}

	saveSteamItems(db, steamResponse.Results)
}

func saveSteamItems(db *sql.DB, steamItems []SteamItem) {
	log.Printf("Saving %d prices to db...", len(steamItems))

	for _, item := range steamItems {
		if !AssetExists(db, item.Name) {
			log.Printf("Creating asset %s...", item.Name)
			asset := &Asset{
				Name:      item.Name,
				NameColor: item.Desc.NameColor,
				IconURL:   item.Desc.IconURL,
				Type:      item.Desc.Type,
			}
			CreateAsset(db, asset)
		}

		price := &Price{
			AssetName: item.Name,
			Time:      time.Now().Unix(),
			Listings:  item.Listings,
			Price:     item.Price,
		}
		InsertPrice(db, price)
	}
}
