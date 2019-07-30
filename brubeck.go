package main

import (
  "fmt"
  "os"
  "time"
  "strconv"
)

func main() {
  argCount := len(os.Args[1:])

  if argCount == 0 {
    fmt.Println(time.Now().Unix())
  } else if argCount == 1 {
    timestamp, err := strconv.ParseInt(os.Args[1], 10, 64)
    if err != nil {
        panic(err)
    }
    tm := time.Unix(timestamp, 0)
    fmt.Println(tm)
  } else if argCount == 3 && os.Args[3] == "ago" {
    delta, err := strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Println(err)
        os.Exit(2)
    }
    start := time.Now()
    switch os.Args[2] {
    case "week":
      before := start.AddDate(0, 0, -delta * 7)
      // moving print out or switch results in undefined before
      fmt.Println(before.Unix())
    }
  }
}
