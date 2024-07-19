package input

import (
	"encoding/csv"
	"strings"
)

// Map the adapted names to actual CSV header names.
var csvHeaderNames = map[string]string{
	"ActivityType":           "Activity Type",
	"Date":                   "Date",
	"Favorite":               "Favorite",
	"Title":                  "Title",
	"Distance":               "Distance",
	"Calories":               "Calories",
	"Time":                   "Time",
	"AvgHR":                  "Avg HR",
	"MaxHR":                  "Max HR",
	"AvgRunCadence":          "Avg Run Cadence",
	"MaxRunCadence":          "Max Run Cadence",
	"AvgPace":                "Avg Pace",
	"BestPace":               "Best Pace",
	"TotalAscent":            "Total Ascent",
	"TotalDescent":           "Total Descent",
	"AvgStrideLength":        "Avg Stride Length",
	"AvgVerticalRatio":       "Avg Vertical Ratio",
	"AvgVerticalOscillation": "Avg Vertical Oscillation",
	"AvgGroundContactTime":   "Avg Ground Contact Time",
	"TrainingStressScore":    "Training Stress Score®",
	"Grit":                   "Grit",
	"Flow":                   "Flow",
	"TotalStrokes":           "Total Strokes",
	"AvgSwolf":               "Avg. Swolf",
	"AvgStrokeRate":          "Avg Stroke Rate",
	"TotalReps":              "Total Reps",
	"Decompression":          "Decompression",
	"BestLapTime":            "Best Lap Time",
	"NumberofLaps":           "Number of Laps",
	"MaxTemp":                "Max Temp",
	"AvgResp":                "Avg Resp",
	"MinResp":                "Min Resp",
	"MaxResp":                "Max Resp",
	"StressChange":           "Stress Change",
	"StressStart":            "Stress Start",
	"StressEnd":              "Stress End",
	"AvgStress":              "Avg Stress",
	"MaxStress":              "Max Stress",
	"MovingTime":             "Moving Time",
	"ElapsedTime":            "Elapsed Time",
	"MinElevation":           "Min Elevation",
	"MaxElevation":           "Max Elevation",
}

// GetCsvColumnIndex Get the column indexes for the header names
func GetCsvColumnIndex(reader *csv.Reader) (map[string]int, error) {
	headers, err := reader.Read()
	if err != nil {
		return nil, err
	}

	csvColumnIndex := make(map[string]int)
	for i, header := range headers {
		header = strings.Trim(header, " \t\n\r")
		for adapted, actual := range csvHeaderNames {
			if actual == header {
				csvColumnIndex[adapted] = i
				break
			}
		}
	}
	return csvColumnIndex, nil
}
