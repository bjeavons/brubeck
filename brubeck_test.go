package main

import (
	"fmt"
	"testing"
	"time"
)

func TestBrubeck(t *testing.T) {
	// Test too many arguments.
	args := []string{"one", "two", "three", "four"}
	emptyResult, _ := brubeck(args)
	if emptyResult != "" {
		t.Errorf("Got %s, expected empty string result", emptyResult)
	}

	// Test empty arguments.
	args = []string{}
	ans, err := brubeck(args)
	if err != nil {
		t.Error("Got error, expected none")
	}
	if len(ans) == 0 {
		t.Error("Got empty answer, expected length > 0")
	}

	// Test 3 arguments.
	args = []string{"1587079799", "in", "pst"}
	ans, err = brubeck(args)
	if err != nil {
		t.Error("Got error, expected none")
	}
	if len(ans) == 0 {
		t.Error("Got empty answer, expected length > 0")
	}
	args = []string{"3", "days", "ago"}
	ans, err = brubeck(args)
	if err != nil {
		t.Error("Got error, expected none")
	}
	if len(ans) == 0 {
		t.Error("Got empty answer, expected length > 0")
	}
}

func TestTimeChange(t *testing.T) {
	// Test error result.
	_, errResult := timeChange(time.Now(), 0, "Nods")
	if errResult == nil {
		t.Errorf("Got no error, expected one")
	}

	// Test day, week, and year forward and backward.
	start := time.Unix(1587079799, 0)
	tomorrow := start.AddDate(0, 0, 1)
	yesterday := start.AddDate(0, 0, -1)
	nextWeek := start.AddDate(0, 0, 7)
	lastWeek := start.AddDate(0, 0, -7)
	nextMonth := start.AddDate(0, 1, 0)
	lastMonth := start.AddDate(0, -1, 0)
	nextYear := start.AddDate(1, 0, 0)
	lastYear := start.AddDate(-1, 0, 0)
	var tests = []struct {
		start time.Time
		amt   int
		unit  string
		want  time.Time
	}{
		{start, 1, "day", tomorrow},
		{start, 1, "Day", tomorrow},
		{start, 1, "days", tomorrow},
		{start, 1, "d", tomorrow},
		{start, -1, "Day", yesterday},
		{start, 1, "week", nextWeek},
		{start, -1, "week", lastWeek},
		{start, 1, "month", nextMonth},
		{start, -1, "month", lastMonth},
		{start, 1, "year", nextYear},
		{start, -1, "year", lastYear},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%s", tt.amt, tt.unit)
		t.Run(testname, func(t *testing.T) {
			ans, _ := timeChange(tt.start, tt.amt, tt.unit)
			if ans != tt.want {
				t.Errorf("Got %s, expected %s", ans, tt.want)
			}
		})
	}
}

func TestTimeConvert(t *testing.T) {
	time := int64(1587079799)
	var tests = []struct {
		time int64
		abr  string
		want string
	}{
		{time, "pdt", "America/Los_Angeles"},
		{time, "mst", "America/Denver"},
		{time, "cdt", "America/Chicago"},
		{time, "est", "America/New_York"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%s", tt.time, tt.abr)
		t.Run(testname, func(t *testing.T) {
			ans, _ := timeConvert(tt.time, tt.abr)
			if ans.Location().String() != tt.want {
				t.Errorf("Got %s, expected %s", ans.Location().String(), tt.want)
			}
		})
	}
}
