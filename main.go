package main

import (
	"activitesSummary/pkg/args"
	"activitesSummary/pkg/data"
	"activitesSummary/pkg/input"
	"activitesSummary/pkg/output"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {

	inputArguments := args.ParseArgs()

	if inputArguments.FileName == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if !(inputArguments.Output == args.DefaultOutput || inputArguments.Output == "json") {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if !(inputArguments.AllSortBy == "" ||
		inputArguments.AllSortBy == "AverageSpeed" ||
		inputArguments.AllSortBy == "AverageHR" ||
		inputArguments.AllSortBy == "MaxHR" ||
		inputArguments.AllSortBy == "Calories" ||
		inputArguments.AllSortBy == "Distance" ||
		inputArguments.AllSortBy == "Time") {
		flag.PrintDefaults()
		os.Exit(1)
	}

	file, err := os.Open(inputArguments.FileName)
	if err != nil {
		fmt.Println("Error opening input:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	csvColumnIndex, err := input.GetCsvColumnIndex(reader)
	if err != nil {
		fmt.Println("Error getting CSV column index:", err)
		return
	}

	if inputArguments.ShowActivities {
		// Read the CSV input and print the distinct list of activities
		activities, err := input.ReadActivitiesFromFile(reader)
		if err != nil {
			fmt.Printf("Error reading activities from input: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Activities:")
		for _, activity := range activities {
			fmt.Println("* ", activity)
		}
		os.Exit(0)
	}

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	activities,
		earliestDate,
		latestDate,
		longestActivities,
		err := input.ParseRecords(records, inputArguments, csvColumnIndex)

	if err != nil {
		fmt.Println("Error parsing records:", err)
		return
	}

	data := data.PopulateData(activities, earliestDate, latestDate, longestActivities, inputArguments)

	if inputArguments.Output == "json" {
		output.PrintJson(data, inputArguments)
	}

	output.PrintText(data, inputArguments)

}
