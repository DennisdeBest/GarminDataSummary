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
		if activitiesMap["All"] || activitiesMap[record[0]] {
			currentActivityType := record[0]
			dateTime, err := time.Parse(constants.DateTimeFormat, record[1])                           // Parse as date-time
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

			// Update earliest and latest dates
			if earliestDate.IsZero() || date.Before(earliestDate) {
				earliestDate = dateTime
			}
			if latestDate.IsZero() || date.After(latestDate) {
				latestDate = dateTime
			}

			distance := service.ParseFloatData(record[4])

			//Some activities record in meters, not kilometers
			if currentActivityType == constants.PoolSwim {
				distance = distance / constants.SwimmingFactor
			}

			duration, err := service.ParseDuration(record[6])
			if err != nil {
				fmt.Println("Error parsing duration:", err)
				continue
			}

			calories := service.ParseFloatData(record[5])

			activity := activity.Activity{
				Type:     record[0],
				Distance: distance,
				Duration: duration,
				Date:     date,
				Calories: calories,
			}
			activities = append(activities, activity)

			longestActivities = data.GetLongestActivities(activity, longestActivities)
		}
	}

	return activities, earliestDate, latestDate, longestActivities, nil
}
