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
	"strings"
)

func main() {

	inputArguments := args.ParseArgs()

	if inputArguments.FileName == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if !(inputArguments.Output == "text" || inputArguments.Output == "json") {
		flag.PrintDefaults()
		os.Exit(1)
	}

	file, err := os.Open(inputArguments.FileName)
	if err != nil {
		fmt.Println("Error opening input:", err)
		return
	}
	defer file.Close()

	if inputArguments.ShowActivities {
		// Read the CSV input and print the distinct list of activities
		activities, err := input.ReadActivitiesFromFile(file)
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

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	activitiesList := strings.Split(inputArguments.SelectedActivities, ",")
	activitiesMap := make(map[string]bool)
	for _, activity := range activitiesList {
		activitiesMap[activity] = true
	}

	activities,
		earliestDate,
		latestDate,
		longestActivities,
		err := input.ParseRecords(records, activitiesMap, inputArguments)

	if err != nil {
		fmt.Println("Error parsing records:", err)
		return
	}

	data := data.PopulateData(activities, earliestDate, latestDate, longestActivities)

	if inputArguments.Output == "json" {
		output.PrintJson(data, inputArguments)
	}

	output.PrintText(data, inputArguments)

}
