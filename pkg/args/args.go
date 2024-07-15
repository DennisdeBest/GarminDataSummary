package args

import (
	"flag"
	"fmt"
	"time"
)

const DefaultSelectedActivities = "All"
const DefaultOutput = "text"

type Args struct {
	FileName           string
	ShowActivities     bool
	Favorites          bool
	All                bool
	SelectedActivities string
	Output             string
	StartDate          *time.Time
	EndDate            *time.Time
}

func ParseArgs() Args {
	var filename string
	flag.StringVar(&filename, "file", "", "path to the input csv input")

	var showActivities bool
	flag.BoolVar(&showActivities, "showActivities", false, "show available activities in input")

	var favorites bool
	flag.BoolVar(&favorites, "favorites", false, "get data only for favorite activities")

	var all bool
	flag.BoolVar(&all, "all", false, "show data for all selected activities")

	var selectedActivities string
	flag.StringVar(&selectedActivities, "activity", DefaultSelectedActivities, "comma-separated list of activities")

	var output string
	flag.StringVar(&output, "output", DefaultOutput, "output destination (text or json)")

	var (
		startDateStr *string = flag.String("startDate", "", "start date/time in ISO 8601 format")
		endDateStr   *string = flag.String("endDate", "", "end date/time in ISO 8601 format")
	)

	flag.Parse()

	var (
		startDate *time.Time = parseDateTime(startDateStr)
		endDate   *time.Time = parseDateTime(endDateStr)
	)

	return Args{
		FileName:           filename,
		ShowActivities:     showActivities,
		All:                all,
		SelectedActivities: selectedActivities,
		Output:             output,
		StartDate:          startDate,
		EndDate:            endDate,
		Favorites:          favorites,
	}
}

func parseDateTime(dateStr *string) *time.Time {
	if *dateStr == "" {
		return nil
	}

	format := "2006-01-02T15:04:05"
	justDateFormat := "2006-01-02"
	loc, _ := time.LoadLocation("UTC")

	date, err := time.ParseInLocation(format, *dateStr, loc)
	if err != nil {
		// Failed to parse as full datetime, try with just date
		date, err = time.ParseInLocation(justDateFormat, *dateStr, loc)
		if err != nil {
			fmt.Println("Error parsing date/time:", err)
			return nil
		}
		if format == justDateFormat {
			date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, loc)
		}
	} else {
		date = time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 0, loc)
	}

	return &date
}
