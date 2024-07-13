package activity

import "time"

type Activity struct {
	Type     string
	Distance float64
	Duration time.Duration
	Date     time.Time
	Calories float64
}
