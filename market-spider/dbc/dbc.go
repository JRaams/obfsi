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
		return
	}

	log.Printf("Creating %s table\n", name)
	Exec(db, sql)
}

func Exec(db *sql.DB, sql string) {
	statement, err := db.Prepare(sql)
	if err != nil {
		log.Panicf("Error preparing statement: %s", err.Error())
	}
	_, err = statement.Exec()
	if err != nil {
		log.Panicf("Error executing sql: %s", err.Error())
	}
}
