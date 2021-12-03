package main

import (
	"fmt"
	"net/url"
	"time"
)

// A WeekendAdjustment specifies a day of the week (Sunday = 0, ...).
type WeekendAdjustment int

const (
	NoChange WeekendAdjustment = iota
	Before
	After
)

func main() {
	// adj := NoChange

	now := time.Now().UTC().AddDate(0, 0, 1)
	fmt.Println(now.Format(time.RFC3339))
	d := now.Weekday()
	fmt.Println(d)

	fmt.Println(AdjustForWeekend(now, Before).Weekday())
}

func AdjustForWeekend(now time.Time, adj WeekendAdjustment) time.Time {
	d := now.Weekday()

	if d == time.Saturday {
		switch adj {
		case NoChange:
			return now
		case Before:
			return now.AddDate(0, 0, -1)
		case After:
			return now.AddDate(0, 0, 2)
		}
	} else if d == time.Sunday {
		switch adj {
		case NoChange:
			return now
		case Before:
			return now.AddDate(0, 0, -2)
		case After:
			return now.AddDate(0, 0, 1)
		}
	}
	return now
}

func main__() {
	var u url.URL
	u.Scheme = "https" // "postgresql"
	//u.User = url.UserPassword("{{username}}", "{{password}}")
	u.Host = "127.0.0.1:5432"
	u.Path = "/"
	q := url.Values{}
	q.Add("name", "my mg")
	q.Add("namespace", "demo & default")
	u.RawQuery = q.Encode()

	fmt.Println(u.String())

	//
	//us := u.String()
	//
	//i := strings.Index(us, "://")
	//
	//fmt.Println(us[:i+3] + "{{username}}:{{password}}@" + us[i+3:])
}
