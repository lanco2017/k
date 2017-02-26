package main

import (
	_ "github.com/lib/pq"
	"log"
    "database/sql"
	"os"
//	http://l-lin.github.io/2015/01/31/Golang-Deploy_to_heroku
//	https://godoc.org/github.com/lib/pq
// https://www.opsdash.com/blog/postgres-full-text-search-golang.html
//	https://astaxie.gitbooks.io/build-web-application-with-golang/content/zh/05.4.html
)



// Connect to Heroku database using the OS env DATABASE_URL
func connect() *sql.DB {
	dbUrl := os.Getenv("DATABASE_URL")
	database, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalf("[x] Could not open the connection to the database. Reason: %s", err.Error())
	}
	return database
}

