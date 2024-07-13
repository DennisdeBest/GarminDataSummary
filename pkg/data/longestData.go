package data

import (
	"activitesSummary/pkg/activity"
	"time"
)

type LongestData struct {
	Type         string
	Date         time.Time
	Distance     float64
	AverageSpeed float64
	Duration     time.Duration
	Calories     float64
}

func GetLongestActivities(currentActivity activity.Activity, longestActivities map[string]LongestData) map[string]LongestData {
	if longestActivity, exists := longestActivities[currentActivity.Type]; !exists || currentActivity.Distance > longestActivity.Distance {
		longestActivities[currentActivity.Type] = LongestData{
			Type:         currentActivity.Type,
			Duration:     currentActivity.Duration,
			Date:         currentActivity.Date,
			Distance:     currentActivity.Distance,
			Calories:     currentActivity.Calories,
			AverageSpeed: currentActivity.Distance / currentActivity.Duration.Hours(),
		}
	}
	return longestActivities
}
