package steamitem

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/jraams/obfsi/market-spider/dbc"
)

// InsertPrice inserts an asset price (at the current time) into the database
func InsertPrice(db *sql.DB, price *Price) {
	SQL := fmt.Sprintf(
		`INSERT INTO "prices" ("assets_name", "time", "listings", "price") VALUES ('%s', %d, %d, %d);`,
		strings.ReplaceAll(price.Assets_Name, "'", "''"),
		price.Time,
		price.Listings,
		price.Price,
	)
	dbc.Exec(db, SQL)
}
