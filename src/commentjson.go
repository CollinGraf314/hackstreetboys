package commentjson

import "fmt"

/// Transforms the database query results into json (can be emmitted by the server)
func Export(name string, library string, eventtype string, date string, starttime string, endtime string, description string) string {
  /*
    Export as
    {
      name: $name,
      library: $library,
      eventtype: $eventtype,
      date: $date,
      starttime: $starttime,
      endtime: $endtime,
      description: $description
    }
  */

  return fmt.Sprintf(`
      {
        "name": "%s",
        "library": "%s",
        "eventtype": "%s",
        "date": "%s",
        "starttime": "%s",
        "endtime": "%s",
        "description": "%s"
      }
    `, name, library, eventtype, date, starttime, endtime, description);
}
