package main

import (
	"github.com/jraams/obfsi/market-spider/dbc"
	"github.com/jraams/obfsi/market-spider/steamitem"
)

func main() {
	db := dbc.GetDBCon()
	defer db.Close()

	steamitem.CreateTables(db)
	steamitem.FetchPrices(db)
}
