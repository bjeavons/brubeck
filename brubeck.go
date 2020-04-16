package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	tm, err := brubeck(os.Args[1:])
	if err != nil {
		panic(err)
	}
	fmt.Println(tm)
}

func brubeck(args []string) (string, error) {
	argCount := len(args)

	switch {
	case argCount == 0:
		return strconv.FormatInt(time.Now().Unix(), 10), nil
	case argCount == 1 && (len(args[0]) == 10 || len(args[0]) == 13):
		timestamp, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return "", err
		}
		if len(args[0]) == 13 {
			timestamp = timestamp / 1000
		}
		tm := time.Unix(timestamp, 0)
		return tm.String(), nil
	case argCount == 3 && args[1] == "in":
		timestamp, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return "", err
		}
		if len(args[0]) == 13 {
			timestamp = timestamp / 1000
		}
		tm, err := timeConvert(timestamp, args[2])
		if err != nil {
			return "", err
		}
		return tm.String(), nil
	case argCount == 3 && (args[2] == "ago" || args[2] == "later"):
		amt, err := strconv.Atoi(args[0])
		if err != nil {
			return "", err
		}
		if args[2] == "ago" {
			amt = -amt
		}
		tm, err := timeChange(amt, args[1])
		if err != nil {
			return "", err
		}
		return strconv.FormatInt(tm.Unix(), 10), nil
	}
	return "", nil
}

func timeChange(amt int, unit string) (time.Time, error) {
	start := time.Now()
	switch unit {
	case "day", "days", "d":
		return start.AddDate(0, 0, amt), nil
	case "week", "weeks", "w":
		return start.AddDate(0, 0, amt*7), nil
	case "month", "months", "m":
		return start.AddDate(0, amt, 0), nil
	case "year", "years", "y":
		return start.AddDate(amt, 0, 0), nil
	}
	return start, errors.New("not implemented")
}

func timeConvert(timestamp int64, abr string) (time.Time, error) {
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
	tm := time.Unix(timestamp, 0)
	location, err := time.LoadLocation(loc)
	if err != nil {
		return tm, err
	}
	return tm.In(location), nil
}
