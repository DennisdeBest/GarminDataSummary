package output

import (
	"activitesSummary/pkg/constants"
	"activitesSummary/pkg/data"
	"activitesSummary/pkg/service"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
)

func PrintText(data *data.Data) {
	fmt.Printf("\n --- Summary \n\n")

	table := tablewriter.NewWriter(os.Stdout)
	table.SetColWidth(40)
	table.SetHeader([]string{"Statistic", "Value"})
	table.SetBorder(false)
	table.AppendBulk([][]string{
		{"Total Activities", fmt.Sprintf("%d", data.Summary.TotalActivities)},
		{"Total Distance (km)", fmt.Sprintf("%.2f", data.Summary.TotalDistance)},
		{"Total Time", service.FormatDuration(data.Summary.TotalDuration)},
		{"Total Calories", fmt.Sprintf("%.0f", data.Summary.TotalCalories)},
		{"", ""},
		{"Average Distance per Activity (km)", fmt.Sprintf("%.2f", data.Summary.AverageDistance)},
		{"Average Time per Activity", service.FormatDuration(data.Summary.AverageDuration)},
		{"Average Speed (km/h)", fmt.Sprintf("%.2f", data.Summary.AverageSpeed)},
		{"Average Calories per Activity", fmt.Sprintf("%.2f", data.Summary.AverageCalories)},
		{"", ""},
		{"Average Daily Activities", fmt.Sprintf("%.2f", data.Summary.AverageDailyActivities)},
		{"Average Daily Distance (km)", fmt.Sprintf("%.2f", data.Summary.AverageDailyDistance)},
		{"Average Daily Time (h)", fmt.Sprintf("%s", service.FormatDuration(data.Summary.AverageDailyTime))},
		{"Average Daily Calories", fmt.Sprintf("%.2f", data.Summary.AverageDailyCalories)},
	})

	table.Render()

	fmt.Printf("\n --- Longest Activities \n\n")

	longestActivitiesTable := tablewriter.NewWriter(os.Stdout)
	longestActivitiesTable.SetHeader([]string{"Activity", "Date", "Distance", "Duration", "Average", "Calories"})
	longestActivitiesTable.SetBorder(false)

	for _, longActivity := range data.Longest {
		longActivityAvgSpeed := longActivity.Distance / longActivity.Duration.Hours()
		longestActivitiesTable.AppendBulk([][]string{
			{
				longActivity.Type,
				longActivity.Date.Format(constants.DateTimeFormat),
				fmt.Sprintf("%.2f", longActivity.Distance),
				service.FormatDuration(longActivity.Duration),
				fmt.Sprintf("%.2f", longActivityAvgSpeed),
				fmt.Sprintf("%.2f", longActivity.Calories),
			},
		})
	}
	longestActivitiesTable.Render()

	fmt.Printf(
		"\nActivities range from %s to %s.\n"+
			"A total of %d days between those dates, %d of which had activities\n"+
			"You are active %.2f%% of days",
		data.EarliestDate.Format(constants.DateTimeFormat),
		data.LatestDate.Format(constants.DateTimeFormat),
		data.Days,
		data.ActivityDays,
		(float64(data.ActivityDays)/float64(data.Days))*100,
	)

}
