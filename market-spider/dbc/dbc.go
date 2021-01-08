package dbc

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq" // postgres import
)

const defaultDbSource = "host=0.0.0.0 port=5432 user=obfsi dbname=obfsi sslmode=disable"

// GetDBCon returns a connection to the (postgres) database
func GetDBCon() *sql.DB {
	psqlconn := os.Getenv("DBSOURCE")
	if len(psqlconn) == 0 {
		log.Printf("DBSOURCE env not specified, using default: %s", defaultDbSource)
		psqlconn = defaultDbSource
	}

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Successfully connected to db!")
	return db
}

// Exec executes a specified SQL string in a prepared statement
func Exec(db *sql.DB, sql string) {
	statement, err := db.Prepare(sql)
	if err != nil {
		log.Panicf("Error preparing statement: %s", err.Error())
	}
	defer statement.Close()

	if _, err = statement.Exec(); err != nil {
		log.Panicf("Error executing sql: %s", err.Error())
	}
}
