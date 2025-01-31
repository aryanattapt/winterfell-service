package pkg

import (
	"log"
	"time"
)

func GenerateCurrentTimeStamp() string {
	return time.Now().Format(time.RFC3339Nano)
}

func FormatTime(data time.Time, format string) string {
	return data.Format(format)
}

func ParseTime(data string, format string) (time.Time, error) {
	return time.Parse(format, data)
}

func ParseAndFormatTime(data string, fromFormat string, toFormat string) (string, error) {
	parsedTime, err := time.Parse(fromFormat, data)
	if err != nil {
		return "", err
	}

	return parsedTime.Format(toFormat), nil
}

func CompareIsoDateStringToNow(isoDateString string) (result int, err error) {
	// Define the layout for the ISO 8601 date string
	layout := time.RFC3339

	// Parse the ISO 8601 date string
	parsedTime, err := time.Parse(layout, isoDateString)
	if err != nil {
		return
	}

	// Get the current date
	now := time.Now().UTC()

	// Create new time.Time objects representing the date part only
	parsedDate := time.Date(parsedTime.Year(), parsedTime.Month(), parsedTime.Day(), 0, 0, 0, 0, time.UTC)
	currentDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	log.Println("parsed date: ", parsedDate)
	log.Println("current date: ", currentDate)

	// Compare the dates
	if parsedDate.Before(currentDate) {
		result = -1
	} else if parsedDate.After(currentDate) {
		result = 1
	} else {
		result = 0
	}
	return
}
