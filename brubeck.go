package main

import (
  "fmt"
  "os"
  "time"
  "strconv"
  "errors"
  "strings"
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
  } else if argCount == 3 && os.Args[2] == "in" {
    timestamp, err := strconv.ParseInt(os.Args[1], 10, 64)
    if err != nil {
        panic(err)
    }
    tm, err := time_convert(timestamp, os.Args[3])
    if err != nil {
        panic(err)
    }
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
    fmt.Println(tm.Unix())
  }
}

func time_change(amt int, unit string) (time.Time, error) {
  start := time.Now()
  switch unit {
  case "day", "days", "d":
    return start.AddDate(0, 0, amt), nil
  case "week", "weeks", "w":
    return start.AddDate(0, 0, amt * 7), nil
  case "month", "months", "m":
    return start.AddDate(0, amt, 0), nil
  case "year", "years", "y":
    return start.AddDate(amt, 0, 0), nil
  }
  return start, errors.New("Not implemented")
}

func time_convert(timestamp int64, abr string) (time.Time, error) {
  var loc string
  switch strings.ToLower(abr) {
  case "pst", "pdt":
    loc = "America/Los_Angeles"
  case "mst", "mdt":
    loc = "America/Denver"
  case "cst", "cdt":
    loc = "America/Chicago"
  case "est", "edt":
    loc = "America/New_York"
  }
  location, err := time.LoadLocation(loc)
  if err != nil {
      panic(err)
  }
  tm := time.Unix(timestamp, 0)
  return tm.In(location), nil
}
