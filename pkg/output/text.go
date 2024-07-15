package output

import (
	"activitesSummary/pkg/args"
	"activitesSummary/pkg/constants"
	"activitesSummary/pkg/data"
	"activitesSummary/pkg/service"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
)

func PrintText(data *data.Data, args args.Args) {
	fmt.Printf("\n --- Summary \n\n")

	table := tablewriter.NewWriter(os.Stdout)
	table.SetColWidth(40)
	table.SetHeader([]string{"Statistic", "Value"})
	table.SetBorder(false)
	table.AppendBulk([][]string{
		{"Total Activities", fmt.Sprintf("%d", data.Summary.TotalActivities)},
		{"Total Distance (km)", fmt.Sprintf("%.2f", data.Summary.TotalDistance)},
		{"Total Time", service.FormatDuration(data.Summary.TotalTime)},
		{"Total Calories", fmt.Sprintf("%.0f", data.Summary.TotalCalories)},
		{"", ""},
		{"Average Distance per Activity (km)", fmt.Sprintf("%.2f", data.Summary.AverageDistance)},
		{"Average Time per Activity", service.FormatDuration(data.Summary.AverageTime)},
		{"Average Speed (km/h)", fmt.Sprintf("%.2f", data.Summary.AverageSpeed)},
		{"Average Calories per Activity", fmt.Sprintf("%.2f", data.Summary.AverageCalories)},
		{"Average HR per Activity", fmt.Sprintf("%.2f", data.Summary.AverageHR)},
		{"Average Max HR per Activity", fmt.Sprintf("%.2f", data.Summary.AverageMaxHR)},
		{"", ""},
		{"Average Daily Activities", fmt.Sprintf("%.2f", data.Summary.AverageDailyActivities)},
		{"Average Daily Distance (km)", fmt.Sprintf("%.2f", data.Summary.AverageDailyDistance)},
		{"Average Daily Time (h:m:s)", fmt.Sprintf("%s", service.FormatDuration(data.Summary.AverageDailyTime))},
		{"Average Daily Calories", fmt.Sprintf("%.2f", data.Summary.AverageDailyCalories)},
	})

	table.Render()

	fmt.Printf("\n --- Longest Activities \n\n")

	longestActivitiesTable := tablewriter.NewWriter(os.Stdout)
	longestActivitiesTable.SetHeader([]string{"Activity", "Date", "Distance", "Duration", "Average Speed (km/h)", "Calories", "Average HR", "Max HR"})
	longestActivitiesTable.SetBorder(false)

	for _, longActivity := range data.Longest {
		longActivityAvgSpeed := longActivity.Distance / longActivity.Time.Hours()
		longestActivitiesTable.AppendBulk([][]string{
			{
				longActivity.ActivityType,
				longActivity.Date.Format(constants.DateTimeFormat),
				fmt.Sprintf("%.2f", longActivity.Distance),
				service.FormatDuration(longActivity.Time),
				fmt.Sprintf("%.2f", longActivityAvgSpeed),
				fmt.Sprintf("%.0f", longActivity.Calories),
				fmt.Sprintf("%d", longActivity.AvgHR),
				fmt.Sprintf("%d", longActivity.MaxHR),
			},
		})
	}
	longestActivitiesTable.Render()

	if args.All {
		fmt.Printf("\n --- All Activities \n\n")
		for activityType, activities := range data.SortedActivities {
			fmt.Printf("\n\t --- %s\n\n", activityType)
			activityTypeAllTable := tablewriter.NewWriter(os.Stdout)
			activityTypeAllTable.SetHeader([]string{"Date", "Distance", "Duration", "Average Speed (km/h)", "Calories", "Average HR", "Max HR"})
			activityTypeAllTable.SetBorder(false)
			for _, activity := range activities {
				activityAvgSpeed := activity.Distance / activity.Time.Hours()
				activityTypeAllTable.AppendBulk([][]string{
					{
						activity.Date.Format(constants.DateTimeFormat),
						fmt.Sprintf("%.2f", activity.Distance),
						service.FormatDuration(activity.Time),
						fmt.Sprintf("%.2f", activityAvgSpeed),
						fmt.Sprintf("%.0f", activity.Calories),
						fmt.Sprintf("%d", activity.AvgHR),
						fmt.Sprintf("%d", activity.MaxHR),
					},
				})
			}
			activityTypeAllTable.Render()
		}
		fmt.Println()
	}

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
