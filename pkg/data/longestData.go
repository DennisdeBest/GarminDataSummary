package data

import (
	"activitesSummary/pkg/activity"
	"time"
)

type LongestData struct {
	ActivityType string
	Date         time.Time
	Distance     float64
	AverageSpeed float64
	Time         time.Duration
	Calories     float64
	AvgHR        int64
	MaxHR        int64
}

func GetLongestActivities(currentActivity activity.Activity, longestActivities map[string]LongestData) map[string]LongestData {
	if longestActivity, exists := longestActivities[currentActivity.ActivityType]; !exists || currentActivity.Distance > longestActivity.Distance {
		longestActivities[currentActivity.ActivityType] = LongestData{
			ActivityType: currentActivity.ActivityType,
			Time:         currentActivity.Time,
			Date:         currentActivity.Date,
			Distance:     currentActivity.Distance,
			Calories:     currentActivity.Calories,
			AverageSpeed: currentActivity.Distance / currentActivity.Time.Hours(),
			AvgHR:        currentActivity.AvgHR,
			MaxHR:        currentActivity.MaxHR,
		}
	}
	return longestActivities
}
