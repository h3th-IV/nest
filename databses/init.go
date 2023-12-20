package databses

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var dB *sql.DB

func InitDB(dataSrc string) (*sql.DB, error) {
	var err error
	dB, err = sql.Open("mysql", dataSrc)
	if err != nil {
		log.Fatal(err)
	}
	return dB, dB.Ping()
}

func Close() {
	dB.Close()
}
