package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func pingMySQL(t time.Time) {

	// Get DB details from ENV
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASS")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")

	// Open Conenction
	db, err = sql.Open("mysql", user+":"+pass+"@tcp("+host+":"+port+")/")
	if err != nil {
		fmt.Println("[connection failed]", host, "[error]", err.Error())
	}

	// Ping
	err = db.Ping()
	if err != nil {
		fmt.Println("[ping failed]", host, "[error]", err.Error())
	}
	// Go uses odd an odd time formating reference date: https://flaviocopes.com/go-date-time-format/
	fmt.Println("[success]", "[", t.Format("2006-01-02 15:04:05"), "]", "[", host, "]")

	// Close connection to DB
	defer db.Close()

}

func main() {

	doEvery(100*time.Millisecond, pingMySQL)

}
