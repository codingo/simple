package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// read input date ranges from command-line arguments
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <start_date> <end_date>")
		os.Exit(1)
	}
	startDateStr := os.Args[1]
	endDateStr := os.Args[2]

	// parse input dates to time.Time
	startDate, err := time.Parse("02/01/2006", startDateStr)
	if err != nil {
		panic(err)
	}
	endDate, err := time.Parse("02/01/2006", endDateStr)
	if err != nil {
		panic(err)
	}

	// output date information for each day in the range
	for d := startDate; !d.After(endDate); d = d.AddDate(0, 0, 1) {
		fmt.Printf("Date: %s\n", d.Format("02/01/2006"))
		fmt.Printf("ISO 8601 format: %s\n", d.Format("2006-01-02"))
		fmt.Printf("Unix timestamp: %d\n", d.Unix())
		fmt.Printf("Julian day: %d\n", d.YearDay())
		fmt.Printf("Excel serial number: %d\n", int(d.Sub(time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC)).Hours() / 24))
		fmt.Printf("SQL format: '%s'\n", d.Format("2006-01-02"))
		fmt.Printf("JSON format: {\"date\": \"%s\"}\n", d.Format("02/01/2006"))
		fmt.Printf("XML format: <date>%s</date>\n", d.Format("02/01/2006"))
		fmt.Printf("YAML format: date: %s\n", d.Format("02/01/2006"))
		fmt.Printf("TXT format: %s\n", d.Format("02-01-2006"))
		fmt.Printf("HTML format: <time datetime=\"%s\">%s</time>\n", d.Format("2006-01-02"), d.Format("02/01/2006"))
		fmt.Println()
	}
}
