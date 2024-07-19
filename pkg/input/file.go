package input

import (
	"activitesSummary/pkg/activity"
	"activitesSummary/pkg/args"
	"activitesSummary/pkg/constants"
	"activitesSummary/pkg/data"
	"activitesSummary/pkg/service"
	"encoding/csv"
	"fmt"
	"strconv"
	"time"
)

func ReadActivitiesFromFile(reader *csv.Reader) ([]string, error) {
	reader.LazyQuotes = true

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

func ParseRecords(records [][]string, args args.Args, csvColumnIndex map[string]int) ([]activity.Activity, time.Time, time.Time, map[string]data.LongestData, error) {
	var activities []activity.Activity
	longestActivities := make(map[string]data.LongestData)
	var earliestDate, latestDate time.Time

	activitiesMap := args.SelectedActivities

	for _, record := range records[1:] { // Skip header row
		if activitiesMap["All"] || activitiesMap[record[csvColumnIndex["ActivityType"]]] {
			currentActivityType := record[csvColumnIndex["ActivityType"]]
			dateTime, err := time.Parse(constants.DateTimeFormat, record[csvColumnIndex["Date"]])      // Parse as date-time
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

			favorite, _ := strconv.ParseBool(record[csvColumnIndex["Favorite"]])
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

			distance := service.ParseFloatData(record[csvColumnIndex["Distance"]])

			//Some activities record in meters, not kilometers
			if currentActivityType == constants.PoolSwim {
				distance = distance / constants.SwimmingFactor
			}

			time, err := service.ParseDuration(record[csvColumnIndex["Time"]])
			if err != nil {
				fmt.Println("Error parsing duration:", err)
				continue
			}

			calories := service.ParseFloatData(record[csvColumnIndex["Calories"]])
			avgStrideLength := service.ParseFloatData(record[csvColumnIndex["AvgStrideLength"]])

			avgHR, _ := strconv.ParseInt(record[csvColumnIndex["AvgHR"]], 10, 64)
			maxHR, _ := strconv.ParseInt(record[csvColumnIndex["MaxHR"]], 10, 64)

			avgRunCadence, _ := strconv.ParseInt(record[csvColumnIndex["AvgRunCadence"]], 10, 64)
			maxRunCadence, _ := strconv.ParseInt(record[csvColumnIndex["MaxRunCadence"]], 10, 64)

			avgPower, _ := strconv.ParseInt(record[csvColumnIndex["AvgPower"]], 10, 64)

			totalAscent, _ := strconv.ParseInt(record[csvColumnIndex["TotalAscent"]], 10, 64)
			totalDescent, _ := strconv.ParseInt(record[csvColumnIndex["TotalDescent"]], 10, 64)

			minElevation, _ := strconv.ParseInt(record[csvColumnIndex["MinElevation"]], 10, 64)
			maxElevation, _ := strconv.ParseInt(record[csvColumnIndex["MaxElevation"]], 10, 64)

			numberofLaps, _ := strconv.ParseInt(record[csvColumnIndex["NumberofLaps"]], 10, 64)

			totalStrokes, _ := strconv.ParseInt(record[csvColumnIndex["TotalStrokes"]], 10, 64)
			totalReps, _ := strconv.ParseInt(record[csvColumnIndex["TotalReps"]], 10, 64)

			activity := activity.Activity{
				ActivityType:    currentActivityType,
				Distance:        distance,
				Time:            time,
				Date:            dateTime,
				Calories:        calories,
				Favorite:        favorite,
				AvgHR:           avgHR,
				MaxHR:           maxHR,
				AvgRunCadence:   avgRunCadence,
				MaxRunCadence:   maxRunCadence,
				TotalAscent:     totalAscent,
				TotalDescent:    totalDescent,
				MinElevation:    minElevation,
				MaxElevation:    maxElevation,
				AvgPower:        avgPower,
				AvgStrideLength: avgStrideLength,
				NumberofLaps:    numberofLaps,
				TotalStrokes:    totalStrokes,
				TotalReps:       totalReps,
				Title:           record[csvColumnIndex["Title"]],
			}
			activities = append(activities, activity)

			longestActivities = data.GetLongestActivities(activity, longestActivities)
		}
	}

	return activities, earliestDate, latestDate, longestActivities, nil
}
