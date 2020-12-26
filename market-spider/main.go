package main

import (
	"log"

	"github.com/jraams/obfsi/market-spider/db"
)

func main() {
	db := db.GetCon()
	defer db.Close()
	log.Println(db.Driver())
}
