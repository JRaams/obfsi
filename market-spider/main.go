package main

import (
	"github.com/jraams/obfsi/market-spider/dbc"
	"github.com/jraams/obfsi/market-spider/steamItem"
)

func main() {
	db := dbc.GetDBCon()
	defer db.Close()

	steamItem.CreateTables(db)
	steamItem.FetchPrices(db)
}
