package args

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

const DefaultSelectedActivities = "All"
const DefaultOutput = "cli"

type Args struct {
	FileName           string
	ShowActivities     bool
	Favorites          bool
	All                bool
	AllSummaries       bool
	AllSortBy          string
	HideFields         map[string]bool
	SelectedActivities map[string]bool
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

	var allSummaries bool
	flag.BoolVar(&allSummaries, "allSummaries", false, "show summaries for all selected data types")

	var selectedActivities string
	flag.StringVar(&selectedActivities, "activities", DefaultSelectedActivities, "comma-separated list of activities (ex \"Running,Treadmill Running, Cycling\") default: All")

	var allSortBy string
	flag.StringVar(&allSortBy, "allSortBy", "", "Sort the all activities tables by a column [AverageSpeed, Distance, Time, AverageHR, MaxHR, Calories]")

	var hideFields string
	flag.StringVar(&hideFields, "hideFields", "", "Hide fields for the all activities tables [Title]")

	var output string
	flag.StringVar(&output, "output", DefaultOutput, "output type [cli, json] default: cli")

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
		AllSummaries:       allSummaries,
		SelectedActivities: mapCommaSeparatedValuesString(selectedActivities),
		HideFields:         mapCommaSeparatedValuesString(hideFields),
		Output:             output,
		StartDate:          startDate,
		EndDate:            endDate,
		Favorites:          favorites,
		AllSortBy:          allSortBy,
	}
}

func mapCommaSeparatedValuesString(values string) map[string]bool {
	list := strings.Split(values, ",")
	output := make(map[string]bool)
	for _, value := range list {
		value = strings.TrimSpace(value)
		output[value] = true
	}

	return output
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
