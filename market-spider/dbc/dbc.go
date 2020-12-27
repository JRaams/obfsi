package dbc

import (
	"database/sql"
	"fmt"
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

// TableExists checks whether a table exists
func TableExists(db *sql.DB, name string) bool {
	SQL := `SELECT EXISTS (
		SELECT FROM information_schema.tables
		WHERE  table_schema = 'public'
		AND    table_name   = '%s'
 );`
	rows, err := db.Query(fmt.Sprintf(SQL, name))
	if err != nil {
		log.Panicf("Error checking if %s table exists: %s", name, err.Error())
	}
	defer rows.Close()

	rows.Next()
	var exists bool
	if err = rows.Scan(&exists); err != nil {
		log.Panicf("Error checking if %s table exists: %s", name, err.Error())
	}
	return exists
}

// CreateTable creates a table according
func CreateTable(db *sql.DB, name string, sql string) {
	if TableExists(db, name) {
		return
	}

	log.Printf("Creating %s table\n", name)
	Exec(db, sql)
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
