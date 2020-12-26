package steamItem

import (
	"database/sql"

	"github.com/jraams/obfsi/market-spider/dbc"
)

func CreateTables(db *sql.DB) {
	steamItemsSQL := `CREATE TABLE "steamItems" (
		"id"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"name"	TEXT NOT NULL,
		"name_color"	TEXT NOT NULL,
		"icon_url"	TEXT NOT NULL,
		"type"	TEXT NOT NULL
	);`
	dbc.CreateTable(db, "steamItems", steamItemsSQL)
}
