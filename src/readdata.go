package readdata

/*
 * Note: This is temporary. I don't believe that this will be used at the end
         delete in cleanup if not used
*/

import "os"
import "fmt"

func ReadData () {
  file, openError := os.Open("../data.json")

  if (openError != nil) {
    log.Fatal("An error occured while attempting to open the data file (does it exist?)")
    return _,nil;
  }

  var data []byte = make([]byte, 100);
  nBytes, readError := file.Read(data)
  fmt.Println("%d bytes read", nBytes);
  fmt.Println("Data:")
  for i:=0;i<nBytes;i++ {
    fmt.Print("%v", data[i])
  }
}
