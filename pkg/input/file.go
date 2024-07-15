package input

import (
	"activitesSummary/pkg/activity"
	"activitesSummary/pkg/args"
	"activitesSummary/pkg/constants"
	"activitesSummary/pkg/data"
	"activitesSummary/pkg/service"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

func ReadActivitiesFromFile(file *os.File) ([]string, error) {

	reader := csv.NewReader(file)
	reader.LazyQuotes = true

	_, err := reader.Read() // Read and discard the header
	if err != nil {
		return nil, err
	}

	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// Use a map to find unique activities
	activities := make(map[string]bool)
	for _, line := range lines {
		activity := line[0]
		activities[activity] = true
	}

	// Convert map keys to slice
	var uniqueActivities []string
	for activity := range activities {
		uniqueActivities = append(uniqueActivities, activity)
	}

	return uniqueActivities, nil
}

func ParseRecords(records [][]string, activitiesMap map[string]bool, args args.Args) ([]activity.Activity, time.Time, time.Time, map[string]data.LongestData, error) {
	var activities []activity.Activity
	longestActivities := make(map[string]data.LongestData)
	var earliestDate, latestDate time.Time

	for _, record := range records[1:] { // Skip header row
		if activitiesMap["All"] || activitiesMap[record[CsvColumnIndex["ActivityType"]]] {
			currentActivityType := record[CsvColumnIndex["ActivityType"]]
			dateTime, err := time.Parse(constants.DateTimeFormat, record[CsvColumnIndex["Date"]])      // Parse as date-time
			date := time.Date(dateTime.Year(), dateTime.Month(), dateTime.Day(), 0, 0, 0, 0, time.UTC) // Keep only date part

			if err != nil {
				fmt.Println("Error parsing date:", err)
				continue
			}

			if args.StartDate != nil && dateTime.Before(*args.StartDate) {
				continue
			}

			if args.EndDate != nil && dateTime.After(*args.EndDate) {
				continue
			}

			favorite, _ := strconv.ParseBool(record[CsvColumnIndex["Favorite"]])
			if args.Favorites == true && favorite == false {
				continue
			}

			// Update earliest and latest dates
			if earliestDate.IsZero() || date.Before(earliestDate) {
				earliestDate = dateTime
			}
			if latestDate.IsZero() || date.After(latestDate) {
				latestDate = dateTime
			}

			distance := service.ParseFloatData(record[CsvColumnIndex["Distance"]])

			//Some activities record in meters, not kilometers
			if currentActivityType == constants.PoolSwim {
				distance = distance / constants.SwimmingFactor
			}

			time, err := service.ParseDuration(record[CsvColumnIndex["Time"]])
			if err != nil {
				fmt.Println("Error parsing duration:", err)
				continue
			}

			calories := service.ParseFloatData(record[CsvColumnIndex["Calories"]])

			avgHR, _ := strconv.ParseInt(record[CsvColumnIndex["AvgHR"]], 10, 64)
			maxHR, _ := strconv.ParseInt(record[CsvColumnIndex["MaxHR"]], 10, 64)

			activity := activity.Activity{
				ActivityType: currentActivityType,
				Distance:     distance,
				Time:         time,
				Date:         dateTime,
				Calories:     calories,
				Favorite:     favorite,
				AvgHR:        avgHR,
				MaxHR:        maxHR,
			}
			activities = append(activities, activity)

			longestActivities = data.GetLongestActivities(activity, longestActivities)
		}
	}

	return activities, earliestDate, latestDate, longestActivities, nil
}
