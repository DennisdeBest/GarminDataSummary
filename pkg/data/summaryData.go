package data

import "time"

type SummaryData struct {
	TotalActivities        int
	TotalDistance          float64
	TotalDuration          time.Duration
	TotalCalories          float64
	AverageDistance        float64
	AverageDuration        time.Duration
	AverageSpeed           float64
	AverageCalories        float64
	AverageDailyActivities float64
	AverageDailyDistance   float64
	AverageDailyTime       time.Duration
	AverageDailyCalories   float64
}
