package main

import (
    "fmt"
    "net/http"
    "database/sql"

     _"github.com/mattn/go-sqlite3"
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

type Event struct {
	ID     int
	LibID  int
	Req    bool
	Name   string
	Stime  string
	Etime  string
	Date   string
	Desc   string
}

func handler(w http.ResponseWriter, r *http.Request) {

}

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil { panic(err) }
	if db == nil { panic("db nil") }
	return db
}

func CreateLibraryTable(db *sql.DB) {
	// create table if not exists
	sql_table := `
	CREATE TABLE IF NOT EXISTS libraries(
		Id INT NOT NULL AUtO_INCREMENT,
		Name VARCHAR(255),
		Town VARCHAR(255),
    Mailing VARCHAR(255),
    Street VARCHAR(255),
    Phone VARCHAR(255),
    Web VARCHAR(255),
    PRIMARY KEY (ID)
	);
	`

	_, err := db.Exec(sql_table)
	if err != nil { panic(err) }
}

func createEventTable(db *sql.DB){
	sql_table :=`CREATE TABLE IF NOT EXISTS events(
	id INT
	libId INT
	req  BOOl
	name TEXT
	sTime TEXT
	eTime TEXT
	date TEXT
	desc TEXT
	)`

	_, err := db.Exec(sql_table)
	if err != nil { panic(err) }
}



func libEvents (db *sql.DB, int LID){
	sql_libEvents :=`SELECT * FROM events WHERE  libId = $1`, LID
	rows, err := db.Query(sql_libEvents)
	if err != nil { panic(err) }
	return rows
}

func main() {

    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
