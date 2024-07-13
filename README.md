# Activities Summary

A Go application that processes a CSV of various activities and provides a summary.

[Download the latest version here](https://github.com/DennisdeBest/GarminDataSummary/releases/latest)

## Getting the data

The program needs an exported CSV file of the garmin data. To get it go
to [Garmin Dashboard Activities](https://connect.garmin.com/modern/activities)
scroll all the way down, if not only the initial visible data will be exported. After that click on the `export csv`
button in the top right corner.


## Usage

To run the application, use the following command:

```shell
go run main.go -input ./activities.csv
```

Replace `./activities.csv` with the path to your CSV file.

This command will read the specified CSV file, and the app will process various activities in the file.

The command can give you the list of activity types available in the file with the `-showActivities` flag.

The `activities` flag can take a comma-separated list of activities or "All" to process all activities. For
example, `-activities="Run,Treadmill Running"` or `-activities="All"`.

By default the output is rendered as tables but can also be set to json with `-output=json`

### Examples

> List the available activities in the exported file

```shell
go run main.go -input ./activities.csv -showActivities                      
Activities:
*  Running
*  Treadmill Running
*  Cycling
*  Other
*  Elliptical
*  Kayaking
*  Walking
*  Pool Swim
*  Breathwork
*  Assistance Requested
```

> Show the data for 3 different activities

```shell
go run main.go -input ./testInput/activities.csv -activity "Running,Cycling,Treadmill Running"

 --- Summary 

              STATISTIC              |   VALUE    
-------------------------------------+------------
  Total Activities                   |       251  
  Total Distance (km)                |   1304.00  
  Total Time                         | 114:19:34  
  Total Calories                     |     87441  
                                     |            
  Average Distance per Activity (km) |      5.20  
  Average Time per Activity          | 00:27:20   
  Average Speed (km/h)               |     11.41  
  Average Calories per Activity      |    348.37  
                                     |            
  Average Daily Activities           |      0.34  
  Average Daily Distance (km)        |      1.75  
  Average Daily Time (h)             | 00:09:11   
  Average Daily Calories             |    117.06  

 --- Longest Activities 

      ACTIVITY      |       DATE       | DISTANCE | DURATION | AVERAGE | CALORIES  
--------------------+------------------+----------+----------+---------+-----------
  Running           | March 31, 2024   |    15.01 | 01:44:38 |    8.61 |  1240.00  
  Treadmill Running | November 8, 2022 |    14.08 | 01:31:22 |    9.25 |  1305.00  
  Cycling           | August 15, 2022  |    36.95 | 02:30:53 |   14.69 |   889.00  

Activities range from June 26, 2022 to July 12, 2024 a total of 747 days between those dates
```

> Output as json for all activities

```shell
go run main.go -input ./testInput/activities.csv  -output json 
{
  "Summary": {
    "TotalActivities": 309,
    "TotalDistance": 1515.8129999999985,
    "TotalDuration": 667084000000000,
    "TotalCalories": 113842,
    "AverageDistance": 4.905543689320384,
    "AverageDuration": 2158847896440,
    "AverageSpeed": 8.180269351386023,
    "AverageCalories": 368.42071197411,
    "AverageDailyActivities": 0.41365461847389556,
    "AverageDailyDistance": 2.0292008032128495,
    "AverageDailyTime": 893017402945,
    "AverageDailyCalories": 152.39892904953146
  },
  "Longest": [
    {
      "Type": "Other",
      "Date": "2022-07-16T00:00:00Z",
      "Distance": 1.14,
      "AverageSpeed": 1.722198908938313,
      "Duration": 2383000000000,
      "Calories": 388
    },
    {
      "Type": "Pool Swim",
      "Date": "2022-07-02T00:00:00Z",
      "Distance": 1.022,
      "AverageSpeed": 2.1060103033772184,
      "Duration": 1747000000000,
      "Calories": 309
    },
    {
      "Type": "Breathwork",
      "Date": "2022-07-12T00:00:00Z",
      "Distance": 0,
      "AverageSpeed": 0,
      "Duration": 334000000000,
      "Calories": 0
    },
    {
      "Type": "Treadmill Running",
      "Date": "2022-11-08T00:00:00Z",
      "Distance": 14.08,
      "AverageSpeed": 9.246260488872673,
      "Duration": 5482000000000,
      "Calories": 1305
    },
    {
      "Type": "Cycling",
      "Date": "2022-08-15T00:00:00Z",
      "Distance": 36.95,
      "AverageSpeed": 14.69347177731139,
      "Duration": 9053000000000,
      "Calories": 889
    },
    {
      "Type": "Elliptical",
      "Date": "2022-10-01T00:00:00Z",
      "Distance": 0,
      "AverageSpeed": 0,
      "Duration": 3916000000000,
      "Calories": 801
    },
    {
      "Type": "Kayaking",
      "Date": "2022-07-24T00:00:00Z",
      "Distance": 9.2,
      "AverageSpeed": 5.142857142857143,
      "Duration": 6440000000000,
      "Calories": 696
    },
    {
      "Type": "Assistance Requested",
      "Date": "2022-06-26T00:00:00Z",
      "Distance": 0.09,
      "AverageSpeed": 7.199999999999999,
      "Duration": 45000000000,
      "Calories": 1
    },
    {
      "Type": "Running",
      "Date": "2024-03-31T00:00:00Z",
      "Distance": 15.01,
      "AverageSpeed": 8.607199745141765,
      "Duration": 6278000000000,
      "Calories": 1240
    },
    {
      "Type": "Walking",
      "Date": "2022-11-05T00:00:00Z",
      "Distance": 13.06,
      "AverageSpeed": 2.789108382274426,
      "Duration": 16857000000000,
      "Calories": 1567
    }
  ]
}
```

## Development

### Prerequisites

- Go version 1.21 or newer

### Setup

1. Clone this repository to your local machine.
2. Navigate to the directory containing the source code.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](LICENSE)