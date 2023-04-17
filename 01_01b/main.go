package main

import (
	"flag"
	"log"
	"time"
)

var expectedFormat = "2006-01-02"

// parseTime validates and parses a given date string.
func parseTime(target string) time.Time {
	// panic("NOT IMPLEMENTED")
	timeInput,err := time.Parse(expectedFormat,target)
	if err!=nil || time.Now().After(timeInput) {
		log.Panic("invalid Input")
	}
	return timeInput
}

// calcSleeps returns the number of sleeps until the target.
func calcSleeps(target time.Time) float64 {
	// panic("NOT IMPLEMENTED")
	unbtil := time.Until(target)

	return unbtil.Hours() / 24
}

func main() {
	bday := flag.String("bday", "", "Your next bday in YYYY-MM-DD format")
	flag.Parse()
	target := parseTime(*bday)
	log.Printf("You have %d sleeps until your birthday. Hurray!",
		int(calcSleeps(target)))
}
