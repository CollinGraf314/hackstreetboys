package main

import (
    "fmt"
    "net/http"
    "database/sql"
    _"github.com/mattn/go-sqlite3"
    "os"
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

/// Handles the POST requests from the form on the main page
func handler(w http.ResponseWriter, r *http.Request) {
  var (
    name = r.FormValue("name")
    library = r.FormValue("library")
    eventtype = r.FormValue("eventtype")
    date = r.FormValue("date")
    starttime = r.FormValue("starttime")
    endtime = r.FormValue("enttime")
    description = r.FormValue("description")
  )

  sql := InitDB()
  sql.Query("INSERT INTO events VALUES ($1,$2,$3,$4,$5,$6,$7)",
      name, library, eventtype, date, starttime,
      endtime, description)
}

/// Handler for "/"
func mainHandler(w http.ResponseWriter, r *http.Request) {
fmt.Println("Recieved on /")
  file, err := os.Open("../index.html")
  if (err != nil) {
    fmt.Println("An error occured while trying to open the 'index.html' file");
    return;
  } else {
    fmt.Println("Opened file successfully")
  }

  var htmlBuffer []byte = make([]byte, 1024 * 1024)
  nBytes, errRead := file.Read(htmlBuffer)

  if (errRead != nil) {
    fmt.Println("An error occured while trying to read from the index.html file")
    return;
  } else {
    fmt.Println("Read from the file successfully")
    fmt.Printf("%d bytes read\n", nBytes)
  }

  strContents := string(htmlBuffer[:nBytes])
  fmt.Println(strContents)
  fmt.Fprint(w, strContents)
}

var __db *sql.DB

/// Idempotent function which creates a singleton instance of the database
func InitDB() *sql.DB {
  if (__db == nil) {
	   __db, err := sql.Open("sqlite3", "../realdb.db")
	   if err != nil { panic(err) }
	   if __db == nil { panic("db nil") }
  }

	return __db
}

/// Queries the database for all events at a given library
func libEvents (db *sql.DB, LID int) *sql.Rows {
	sql_libEvents :=`SELECT * FROM events WHERE  libId = $1`
	rows, err := db.Query(sql_libEvents, LID)
	if err != nil { panic(err) }
	return rows
}

func main() {

    http.HandleFunc("/submitForm", handler)
    http.HandleFunc("/", mainHandler)
    http.ListenAndServe(":8080", nil)
}
