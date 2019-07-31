package main

import (
  "fmt"
  "os"
  "time"
  "strconv"
  "errors"
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
  } else if argCount == 3 && (os.Args[3] == "ago" || os.Args[3] == "later") {
    amt, err := strconv.Atoi(os.Args[1])
    if err != nil {
        panic(err)
    }
    if os.Args[3] == "ago" {
      amt = -amt
    }
    tm, err := time_change(amt, os.Args[2])
    if err != nil {
        panic(err)
    }
    fmt.Println(tm)
  }
}

func time_change(amt int, unit string) (int64, error) {
  start := time.Now()
  switch unit {
  case "day", "days", "d":
    return start.AddDate(0, 0, amt).Unix(), nil
  case "week", "weeks", "w":
    return start.AddDate(0, 0, amt * 7).Unix(), nil
  case "month", "months", "m":
    return start.AddDate(0, amt, 0).Unix(), nil
  case "year", "years", "y":
    return start.AddDate(amt, 0, 0).Unix(), nil
  }
  return -1, errors.New("Not implemented")
}
