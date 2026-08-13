package db

import (
	"database/sql"
	// "flag"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Wdym != nil")
	}
}

func ConnectDB() (*sql.DB, error) {
	dbUser, dbUserErr := os.LookupEnv("DBUSER")
	if !dbUserErr {
		log.Fatal(dbUserErr)
	}
	dbName, dbNameErr := os.LookupEnv("DBNAME")
	if !dbNameErr {
		log.Fatal(dbNameErr)
	}
	dbPassword, dbPassErr := os.LookupEnv("DBPASSWORD")
	if !dbPassErr {
		log.Fatal(dbPassErr)
	}
	connStr := "user=" + dbUser + " dbname=" + dbName + " password=" + dbPassword + " sslmode=disable"
	db, dbErr := sql.Open("postgres", connStr)
	
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	dbErr = db.Ping()
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	pool := db
	return pool, nil
}
