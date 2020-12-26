package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const dbName = "market.db"

func GetCon() *sql.DB {
	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		log.Printf("Creating %s...", dbName)
		file, err2 := os.Create(dbName)
		if err2 != nil {
			log.Fatal(err.Error())
		}
		file.Close()
		log.Printf("%s created!", dbName)
	} else {
		log.Println("Existing db found!")
	}

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Succesfully connected to db!")
	return db
}
