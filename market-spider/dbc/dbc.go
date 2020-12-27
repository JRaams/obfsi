package dbc

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const dbName = "market.db"

func GetDBCon() *sql.DB {
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

func TableExists(db *sql.DB, name string) bool {
	rows, err := db.Query(fmt.Sprintf("PRAGMA table_info(%s);", name))
	if err != nil {
		log.Panicf("Error checking if %s table exists: %s", name, err.Error())
	}
	defer rows.Close()
	return rows.Next()
}

func CreateTable(db *sql.DB, name string, sql string) {
	if TableExists(db, name) {
		log.Printf("Table %s already exists\n", name)
	} else {
		log.Printf("Creating %s table\n", name)
		statement, err := db.Prepare(sql)
		if err != nil {
			log.Panicf("Error preparing %s table statement: %s", name, err.Error())
		}
		result, err := statement.Exec()
		if err != nil {
			log.Panicf("Error creating %s table: %s", name, err.Error())
		}
		log.Println(result)
	}
}
