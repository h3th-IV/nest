package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var dB *sql.DB

func InitDB() (*sql.DB, error) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("SQL_USER"), os.Getenv("SQL_PASSWORD"), os.Getenv("SQL_HOST"), os.Getenv("SQL_PORT"), os.Getenv("SQL_DBNAME"))
	var err error
	dB, err = sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = dB.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Connection to Database was successful")
	return dB, nil
}

func Close() {
	dB.Close()
}
