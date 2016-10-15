package readdata

/*
 * Note: This is temporary. I don't believe that this will be used at the end
         delete in cleanup if not used
*/

import "os"
import "fmt"
import "encoding/csv"

func ReadData ()  {
  file, openError := os.Open("../data.json")

  if (openError != nil) {
    panic(openError)
  }

  var data []byte = make([]byte, 100)
  nBytes, _ := file.Read(data)
  fmt.Println("%d bytes read", nBytes)
  fmt.Println("Data:")
  for i:=0;i< nBytes;i++ {
    fmt.Print("%v", data[i])
  }
}
