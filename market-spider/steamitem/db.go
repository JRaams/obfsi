package steamitem

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jraams/obfsi/market-spider/dbc"
)

var AssetNameIDMap = map[string]int{}

func LoadAssetNameIDs(db *sql.DB) {
	SQL := `SELECT "name", id FROM "assets"`
	rows, err := db.Query(SQL)
	if err != nil {
		log.Panicf("Error loading asset names and ids: %s", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		var id int
		err := rows.Scan(&name, &id)
		if err != nil {
			log.Panicf("Error scanning row: %s", err.Error())
		}
		AssetNameIDMap[name] = id
	}
}

// InsertPrice inserts an asset price (at the current time) into the database
func InsertPrice(db *sql.DB, price *Price) {
	SQL := fmt.Sprintf(
		`INSERT INTO "prices" ("asset_id", "time", "listings", "price") VALUES (%d, %d, %d, %d);`,
		price.Asset_ID,
		price.Time,
		price.Listings,
		price.Price,
	)
	dbc.Exec(db, SQL)
}
