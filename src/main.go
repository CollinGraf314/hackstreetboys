package main

import (
    "fmt"
    "net/http"
    "database/sql"
    _"github.com/mattn/go-sqlite3"
    "os"
    "strings"
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
	LibID  string
	Req    string
	Name   string
	Stime  string
	Etime  string
	Date   string
	Desc   string
}

func readHtml(filename, requests, hosts string) string {
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

  strContents = strings.Replace(strContents, "___REQUESTED_EVENTS_ELEMENTS___", requests, 1)
  strContents = strings.Replace(strContents, "___HOSTED_EVENTS_ELEMENTS___", hosts, 1)


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

  table := InitDB()
  if (table == nil) {
    fmt.Println("ERROR: SQL is NIL")
    return;
  }

  table.Query("INSERT INTO events (Name, LibId, Req, `date`, Stime, Etime, `desc`) VALUES (\"$1\",$2,\"$3\",\"$4\",\"$5\",\"$6\",\"$7\");",
      name, library, eventtype, date, starttime,
      endtime, description)

  requests := libEvents(table, library, "request")
  hosts := libEvents(table, library, "host")

  content := readHtml("../submitForm.html", requests, hosts)
  fmt.Fprint(w, content)
}


/// Handler for "/"
func mainHandler(w http.ResponseWriter, r *http.Request) {

  table := InitDB()
  if (table == nil) {
    fmt.Println("ERROR: SQL is NIL")
    return
  }

  requests := libEvents(table, "Guildhall Public", "request")
  hosts := libEvents(table, "Guildhall Public", "host")

  strContents := readHtml("../index.html", requests, hosts)
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
func libEvents(db *sql.DB, LID string, eType string) string {
	sql_libEvents :=`SELECT * FROM events WHERE libId = $1 and Req = $2`
	rows, err := db.Query(sql_libEvents, LID, eType)
	if err != nil { panic(err) }
  defer rows.Close()

  var element string

  for rows.Next() {
    var name, library, eventtype, date, starttime, endtime, description string

      err = rows.Scan(&name, &library, &eventtype, &date, &starttime, &endtime, &description)

      element += fmt.Sprintf(`
        <tr>
          <td>%s</td>
          <td>%s</td>
          <td>%s</td>
          <td>%s</td>
          <td>%s</td>
          <td>%s</td>
          <td>%s</td>
        </tr>
      `, name, library, eventtype, date, starttime, endtime, description)
  }

	return element
}



func main() {
    // http.HandlerFunc("/res/Icon.png", func (w http.ResponseWriter, r *http.Request) {
    //   http.ServeFile(w,r,"/res/Icon.png")
    // })
    // http.HandlerFunc("/libraryLocation.png", func (w http.ResponseWriter, r *http.Request) {
    //   http.ServeFile(w,r,"/libraryLocation.png")
    // })
    // http.HandlerFunc("/myLocation.png", func (w http.ResponseWriter, r *http.Request) {
    //   http.ServeFile(w,r,"/myLocation.png")
    // })
    // http.HandlerFunc("/scriptfile.js", func (w http.ResponseWriter, r *http.Request) {
    //   http.ServeFile(w,r,"/scriptfile.js")
    // })
    // http.HandlerFunc("/style.css", func (w http.ResponseWriter, r *http.Request) {
    //   http.ServeFile(w,r,"/style.css")
    // })
    // http.HandleFunc("/index.html", func (w http.ResponseWriter, r *http.Request) {
    //   http.ServeFile(w,r,"/index.html")
    // })
    // http.HandleFunc("/submitForm.html", handler)
    //
    // http.ListenAndServe(":8080", nil)

    http.HandleFunc("index.html", func(w http.ResponseWriter, r *http.Request) {
      http.ServeFile(w,r,"index.html")
    })
    http.HandleFunc("/res/Icon.png", func(w http.ResponseWriter, r *http.Request) {
      http.ServeFile(w,r,"/res/Icon.png")
    })
    http.HandleFunc("/libraryLocation.png", func(w http.ResponseWriter, r *http.Request) {
      http.ServeFile(w,r,"/libraryLocation.png")
    })
    http.HandleFunc("/myLocation.png", func(w http.ResponseWriter, r *http.Request) {
      http.ServeFile(w,r,"/myLocation.png")
    })
    http.HandleFunc("/scriptfile.js", func(w http.ResponseWriter, r *http.Request) {
      http.ServeFile(w,r,"/scriptfile.js")
    })
    http.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
      http.ServeFile(w,r,"/style.css")
    })
    http.HandleFunc("/submitForm.html", handler)

    http.ListenAndServe(":8080", nil)
}
