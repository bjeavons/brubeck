package main

import (
  "fmt"
  "os"
  "time"
  "strconv"
)

func main() {

  if len(os.Args) > 1 {
    timestamp, err := strconv.ParseInt(os.Args[1], 10, 64)
    if err != nil {
        panic(err)
    }
    tm := time.Unix(timestamp, 0)
    fmt.Println(tm)
  } else {
    fmt.Println(time.Now().Unix())
  }
}
