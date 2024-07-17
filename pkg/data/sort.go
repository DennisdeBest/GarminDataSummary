package data

import "activitesSummary/pkg/activity"

type ByDistance []activity.Activity

func (a ByDistance) Len() int {
	return len(a)
}

func (a ByDistance) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByDistance) Less(i, j int) bool {
	return a[i].Distance > a[j].Distance
}

type ByAverageSpeed []activity.Activity

func (a ByAverageSpeed) Len() int {
	return len(a)
}

func (a ByAverageSpeed) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAverageSpeed) Less(i, j int) bool {
	return (a[i].Distance / a[i].Time.Hours()) > (a[j].Distance / a[j].Time.Hours())
}

type ByAverageHR []activity.Activity

func (a ByAverageHR) Len() int {
	return len(a)
}

func (a ByAverageHR) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAverageHR) Less(i, j int) bool {
	return a[i].AvgHR > a[j].AvgHR
}

type ByMaxHR []activity.Activity

func (a ByMaxHR) Len() int {
	return len(a)
}

func (a ByMaxHR) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByMaxHR) Less(i, j int) bool {
	return a[i].AvgHR > a[j].AvgHR
}

type ByDate []activity.Activity

func (a ByDate) Len() int {
	return len(a)
}

func (a ByDate) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByDate) Less(i, j int) bool {
	return a[i].Date.Before(a[j].Date)
}

type ByTime []activity.Activity

func (a ByTime) Len() int {
	return len(a)
}

func (a ByTime) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByTime) Less(i, j int) bool {
	return a[i].Time > a[j].Time
}

type ByCalories []activity.Activity

func (a ByCalories) Len() int {
	return len(a)
}

func (a ByCalories) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByCalories) Less(i, j int) bool {
	return a[i].Calories > a[j].Calories
}
