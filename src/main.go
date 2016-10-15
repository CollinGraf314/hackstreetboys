package main

import (
    "fmt"
    "net/http"
    "database/sql"
    "github.com/mattn/go-sqlite3"
)

type Libraries struct {
  ID      int
  Name    string
  Town    string
  Mailing string
  Street  string
  Phone   string
  Web     string
}

func handler(w http.ResponseWriter, r *http.Request) {

}

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil { panic(err) }
	if db == nil { panic("db nil") }
	return db
}

func CreateTable(db *sql.DB) {
	// create table if not exists
	sql_table := `
	CREATE TABLE IF NOT EXISTS items(
		Id TEXT NOT NULL PRIMARY KEY,
		Name TEXT,
		Phone TEXT,
		InsertedDatetime DATETIME
	);
	`

	_, err := db.Exec(sql_table)
	if err != nil { panic(err) }
}

func main() {

    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
