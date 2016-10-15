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

func readHtml(filename string) string {
  file, err := os.Open(filename)
  if (err != nil) {
    fmt.Printf("An error occured while trying to open the '%s' file\n", filename);
    return "";
  } else {
    fmt.Println("Opened file successfully")
  }

  var htmlBuffer []byte = make([]byte, 1024 * 1024)
  nBytes, errRead := file.Read(htmlBuffer)

  if (errRead != nil) {
    fmt.Printf("An error occured while trying to read from the %s file\n", filename)
    return "";
  } else {
    fmt.Println("Read from the file successfully")
    fmt.Printf("%d bytes read\n", nBytes)
  }

  strContents := string(htmlBuffer[:nBytes])
  return strContents
}

/// Handles the GET requests from the form on the main page
func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Handler on /submitForm")
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
  if (sql == nil) {
    fmt.Println("ERROR: SQL is NIL")
    return;
  }

  sql.Query("INSERT INTO events (Name, LibId, Req, `date`, Stime, Etime, `desc`) VALUES ($1,$2,$3,$4,$5,$6,$7)",
      name, library, eventtype, date, starttime,
      endtime, description)

  content := readHtml("../submitForm.html")
  fmt.Println(content)
  fmt.Fprint(w, content)
}

/// Handler for "/"
func mainHandler(w http.ResponseWriter, r *http.Request) {
  strContents := readHtml("../index.html")
  fmt.Println(strContents)
  fmt.Fprint(w, strContents)
}

/// Idempotent function which creates a singleton instance of the database
func InitDB() *sql.DB {
	__db, err := sql.Open("sqlite3", "../realdb.db")
	if err != nil { panic(err) }
  if __db == nil { panic("db nil") }

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
