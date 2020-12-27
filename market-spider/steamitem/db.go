package steamitem

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jraams/obfsi/market-spider/dbc"
)

// CreateTables creates the necessary tables to store Steam market assets and prices
func CreateTables(db *sql.DB) {
	assetsSQL := `CREATE TABLE "assets" (
		"name"	TEXT NOT NULL PRIMARY KEY,
		"nameColor"	TEXT NOT NULL,
		"iconUrl"	TEXT NOT NULL,
		"type"	TEXT NOT NULL
	);`
	dbc.CreateTable(db, "assets", assetsSQL)

	pricesSQL := `CREATE TABLE "prices" (
		"id"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"assets_name"	TEXT NOT NULL,
		"time"	INTEGER NOT NULL,
		"listings"	INTEGER NOT NULL,
		"price"	INTEGER NOT NULL,
		FOREIGN KEY("assets_name") REFERENCES "assets"("name")
	);`
	dbc.CreateTable(db, "prices", pricesSQL)
}

// AssetExists checks whether an asset with a name already exists in the database
func AssetExists(db *sql.DB, name string) bool {
	rows, err := db.Query(fmt.Sprintf(`SELECT * FROM assets WHERE name = "%s";`, name))
	if err != nil {
		log.Panicf("Error checking if asset %s exists: %s", name, err.Error())
	}
	defer rows.Close()
	return rows.Next()
}

// CreateAsset inserts a Steam asset into the database
func CreateAsset(db *sql.DB, asset *Asset) {
	SQL := fmt.Sprintf(`INSERT INTO "assets" ("name", "nameColor", "iconUrl", "type") VALUES 
	("%s", "%s", "%s", "%s");`, asset.Name, asset.NameColor, asset.IconURL, asset.Type)
	dbc.Exec(db, SQL)
}

// InsertPrice inserts an asset price (at the current time) into the database
func InsertPrice(db *sql.DB, price *Price) {
	SQL := fmt.Sprintf(`INSERT INTO "prices" ("assets_name", "time", "listings", "price") VALUES
	("%s", %d, %d, %d);`, price.AssetName, price.Time, price.Listings, price.Price)
	dbc.Exec(db, SQL)
}