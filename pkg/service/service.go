package service

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func ParseDuration(s string) (time.Duration, error) {
	parts := strings.Split(s, ":")
	if len(parts) != 3 {
		return 0, errors.New("invalid time format")
	}
	h, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, err
	}
	m, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, err
	}
	sec, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		return 0, err
	}
	return time.Duration(h)*time.Hour + time.Duration(m)*time.Minute + time.Duration(sec)*time.Second, nil
}

func FormatDuration(d time.Duration) string {
	d = d.Round(time.Second)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}

func ParseFloatData(data string) float64 {
	if data == "--" || data == "NaN" {
		return (float64(0))
	}

	data = strings.Replace(data, ",", "", -1)

	result, err := strconv.ParseFloat(data, 64)
	if err != nil {
		fmt.Println("Error parsing distance:", err)
	}

	return result
}
