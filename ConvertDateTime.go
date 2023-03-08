package main

import (
	"fmt"
	"time"
)

func main() {
	// get input date from user
	var inputDate string
	fmt.Print("Enter date (dd/mm/yyyy): ")
	fmt.Scan(&inputDate)

	// parse input date to time.Time
	t, err := time.Parse("02/01/2006", inputDate)
	if err != nil {
		panic(err)
	}

	// output date in various formats
	fmt.Printf("ISO 8601 format: %s\n", t.Format("2006-01-02"))
	fmt.Printf("Unix timestamp: %d\n", t.Unix())
	fmt.Printf("Julian day: %d\n", t.YearDay())
	fmt.Printf("Excel serial number: %d\n", int(t.Sub(time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC)).Hours() / 24))
	fmt.Printf("SQL format: '%s'\n", t.Format("2006-01-02"))
	fmt.Printf("JSON format: {\"date\": \"%s\"}\n", t.Format("02/01/2006"))
	fmt.Printf("XML format: <date>%s</date>\n", t.Format("02/01/2006"))
	fmt.Printf("YAML format: date: %s\n", t.Format("02/01/2006"))
	fmt.Printf("TXT format: %s\n", t.Format("02-01-2006"))
	fmt.Printf("HTML format: <time datetime=\"%s\">%s</time>\n", t.Format("2006-01-02"), t.Format("02/01/2006"))
}
