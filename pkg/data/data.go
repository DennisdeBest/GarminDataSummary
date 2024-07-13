package data

import (
	"activitesSummary/pkg/activity"
	"activitesSummary/pkg/constants"
	"time"
)

type Data struct {
	Summary      SummaryData
	Longest      []LongestData
	EarliestDate time.Time
	LatestDate   time.Time
	Days         int
	ActivityDays int
}

func PopulateData(
	activities []activity.Activity,
	earliestDate time.Time,
	latestDate time.Time,
	longestActivities map[string]LongestData,
) *Data {
	totalActivities := len(activities)
	var totalDistance float64
	var totalCalories float64
	var totalDuration time.Duration
	uniqueDates := make(map[string]bool)
	for _, activity := range activities {
		totalDistance += activity.Distance
		totalDuration += activity.Duration
		totalCalories += activity.Calories
		dateString := activity.Date.Format(constants.DateFormat)
		uniqueDates[dateString] = true
	}

	numUniqueDays := len(uniqueDates)

	durationDenominator := time.Duration(totalActivities)
	if durationDenominator == 0 {
		durationDenominator = 1
	}

	caloriesDenominator := float64(totalActivities)
	if caloriesDenominator == 0 {
		caloriesDenominator = 1
	}

	avgDistance := totalDistance / float64(totalActivities)
	avgSpeed := totalDistance / totalDuration.Hours()
	avgDuration := totalDuration / durationDenominator
	avgCalories := totalCalories / caloriesDenominator

	days := latestDate.Sub(earliestDate).Hours() / 24

	dailyAverageActivities := float64(totalActivities) / days
	dailyAverageDistance := totalDistance / days
	dailyAverageTime := time.Duration(totalDuration.Hours() / days * float64(time.Hour))
	dailyAverageCalories := float64(totalCalories) / days

	var longestDataSlice []LongestData
	for _, value := range longestActivities {
		longestDataSlice = append(longestDataSlice, value)
	}

	data := &Data{
		Summary: SummaryData{
			TotalActivities:        totalActivities,
			TotalDistance:          totalDistance,
			TotalDuration:          totalDuration,
			TotalCalories:          totalCalories,
			AverageDistance:        avgDistance,
			AverageCalories:        avgCalories,
			AverageSpeed:           avgSpeed,
			AverageDuration:        avgDuration,
			AverageDailyDistance:   dailyAverageDistance,
			AverageDailyCalories:   dailyAverageCalories,
			AverageDailyTime:       dailyAverageTime,
			AverageDailyActivities: dailyAverageActivities,
		},
		Longest:      longestDataSlice,
		EarliestDate: earliestDate,
		LatestDate:   latestDate,
		Days:         int(latestDate.Sub(earliestDate).Hours() / 24),
		ActivityDays: numUniqueDays,
	}
	return data
}
