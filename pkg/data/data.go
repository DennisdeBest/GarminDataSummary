package data

import (
	"activitesSummary/pkg/activity"
	"activitesSummary/pkg/args"
	"activitesSummary/pkg/constants"
	"sort"
	"time"
)

type Data struct {
	Summary               SummaryData
	ActivityTypeSummaries map[string]SummaryData
	Longest               []LongestData
	SortedActivities      map[string][]activity.Activity
	EarliestDate          time.Time
	LatestDate            time.Time
	Days                  int
	ActivityDays          int
}

func PopulateData(
	activities []activity.Activity,
	earliestDate time.Time,
	latestDate time.Time,
	longestActivities map[string]LongestData,
	args args.Args,
) *Data {
	totalActivities := len(activities)
	var totalDistance float64
	var totalCalories float64
	var totalAverageHR int64
	var totalMaxHR int64
	var totalDuration time.Duration
	var sortedActivities = make(map[string][]activity.Activity)
	uniqueDates := make(map[string]bool)
	for _, activity := range activities {
		totalDistance += activity.Distance
		totalDuration += activity.Time
		totalCalories += activity.Calories
		totalMaxHR += activity.MaxHR
		totalAverageHR += activity.AvgHR
		dateString := activity.Date.Format(constants.DateFormat)
		uniqueDates[dateString] = true
		sortedActivities[activity.ActivityType] = append(sortedActivities[activity.ActivityType], activity)

		if args.AllSortBy == "Distance" {
			sort.Sort(ByDistance(sortedActivities[activity.ActivityType]))
		}

		if args.AllSortBy == "Time" {
			sort.Sort(ByTime(sortedActivities[activity.ActivityType]))
		}

		if args.AllSortBy == "Calories" {
			sort.Sort(ByCalories(sortedActivities[activity.ActivityType]))
		}

		if args.AllSortBy == "AverageSpeed" {
			sort.Sort(ByAverageSpeed(sortedActivities[activity.ActivityType]))
		}

		if args.AllSortBy == "AverageHR" {
			sort.Sort(ByAverageHR(sortedActivities[activity.ActivityType]))
		}

		if args.AllSortBy == "MaxHR" {
			sort.Sort(ByMaxHR(sortedActivities[activity.ActivityType]))
		}
	}

	//Summary data per activity type
	var activityTypeSummaries = make(map[string]SummaryData)
	for activityType, activitiesForType := range sortedActivities {

		var activityTypeTotalDistance float64
		var activityTypeTotalCalories float64
		var activityTypeTotalAverageHR int64
		var activityTypeTotalMaxHR int64
		var activityTypeTotalDuration time.Duration
		activityTypeUniqueDates := make(map[string]bool)

		for _, activityForType := range activitiesForType {
			activityTypeTotalDistance += activityForType.Distance
			activityTypeTotalDuration += activityForType.Time
			activityTypeTotalCalories += activityForType.Calories
			activityTypeTotalMaxHR += activityForType.MaxHR
			activityTypeTotalAverageHR += activityForType.AvgHR
			activityTypeDateString := activityForType.Date.Format(constants.DateFormat)
			activityTypeUniqueDates[activityTypeDateString] = true
		}

		totalActivitiesForType := len(activitiesForType)

		durationDenominator := time.Duration(totalActivitiesForType)
		if durationDenominator == 0 {
			durationDenominator = 1
		}

		floatDenominator := float64(totalActivitiesForType)
		if floatDenominator == 0 {
			floatDenominator = 1
		}

		activityTypeAvgDistance := activityTypeTotalDistance / floatDenominator
		activityTypeAvgSpeed := activityTypeTotalDistance / activityTypeTotalDuration.Hours()
		activityTypeAvgDuration := activityTypeTotalDuration / durationDenominator
		activityTypeAvgCalories := activityTypeTotalCalories / floatDenominator
		activityTypeAvgHr := float64(activityTypeTotalAverageHR) / floatDenominator
		activityTypeAvgMaxHr := float64(activityTypeTotalMaxHR) / floatDenominator

		days := latestDate.Sub(earliestDate).Hours() / 24

		activityTypeDailyAverageActivities := float64(totalActivitiesForType) / days
		activityTypeDailyAverageDistance := activityTypeTotalDistance / days
		activityTypeDailyAverageTime := time.Duration(activityTypeTotalDuration.Hours() / days * float64(time.Hour))
		activityTypeDailyAverageCalories := float64(activityTypeTotalCalories) / days

		summary := SummaryData{
			TotalActivities:        totalActivitiesForType,
			TotalDistance:          activityTypeTotalDistance,
			TotalTime:              activityTypeTotalDuration,
			TotalCalories:          activityTypeTotalCalories,
			AverageDistance:        activityTypeAvgDistance,
			AverageCalories:        activityTypeAvgCalories,
			AverageSpeed:           activityTypeAvgSpeed,
			AverageTime:            activityTypeAvgDuration,
			AverageHR:              activityTypeAvgHr,
			AverageMaxHR:           activityTypeAvgMaxHr,
			AverageDailyDistance:   activityTypeDailyAverageDistance,
			AverageDailyCalories:   activityTypeDailyAverageCalories,
			AverageDailyTime:       activityTypeDailyAverageTime,
			AverageDailyActivities: activityTypeDailyAverageActivities,
		}

		activityTypeSummaries[activityType] = summary

		activityTypeTotalDistance = 0
		activityTypeTotalDuration = 0
		activityTypeTotalCalories = 0
		activityTypeTotalMaxHR = 0
		activityTypeTotalAverageHR = 0
	}

	//Total summary data

	numUniqueDays := len(uniqueDates)

	durationDenominator := time.Duration(totalActivities)
	if durationDenominator == 0 {
		durationDenominator = 1
	}

	floatDenominator := float64(totalActivities)
	if floatDenominator == 0 {
		floatDenominator = 1
	}

	avgDistance := totalDistance / float64(totalActivities)
	avgSpeed := totalDistance / totalDuration.Hours()
	avgDuration := totalDuration / durationDenominator
	avgCalories := totalCalories / floatDenominator
	avgHr := float64(totalAverageHR) / floatDenominator
	avgMaxHr := float64(totalMaxHR) / floatDenominator

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
			TotalTime:              totalDuration,
			TotalCalories:          totalCalories,
			AverageDistance:        avgDistance,
			AverageCalories:        avgCalories,
			AverageSpeed:           avgSpeed,
			AverageTime:            avgDuration,
			AverageHR:              avgHr,
			AverageMaxHR:           avgMaxHr,
			AverageDailyDistance:   dailyAverageDistance,
			AverageDailyCalories:   dailyAverageCalories,
			AverageDailyTime:       dailyAverageTime,
			AverageDailyActivities: dailyAverageActivities,
		},
		ActivityTypeSummaries: activityTypeSummaries,
		Longest:               longestDataSlice,
		SortedActivities:      sortedActivities,
		EarliestDate:          earliestDate,
		LatestDate:            latestDate,
		Days:                  int(latestDate.Sub(earliestDate).Hours() / 24),
		ActivityDays:          numUniqueDays,
	}
	return data
}
