package args

import (
	"flag"
)

const DefaultSelectedActivities = "All"
const DefaultOutput = "text"

type Args struct {
	FileName           string
	ShowActivities     bool
	SelectedActivities string
	Output             string
}

func ParseArgs() Args {
	var filename string
	flag.StringVar(&filename, "file", "", "path to the input csv input")

	var showActivities bool
	flag.BoolVar(&showActivities, "showActivities", false, "show available activities in input")

	var selectedActivities string
	flag.StringVar(&selectedActivities, "activity", DefaultSelectedActivities, "comma-separated list of activities")

	var output string
	flag.StringVar(&output, "output", DefaultOutput, "output destination (text or json)")

	flag.Parse()

	return Args{
		FileName:           filename,
		ShowActivities:     showActivities,
		SelectedActivities: selectedActivities,
		Output:             output,
	}
}
