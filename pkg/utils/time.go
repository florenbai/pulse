package utils

import "time"

func GetTimeAgo(timeRange string) time.Time {
	now := time.Now()
	switch timeRange {
	case "30m":
		return now.Add(-30 * time.Minute)
	case "1h":
		return now.Add(-1 * time.Hour)
	case "2h":
		return now.Add(-2 * time.Hour)
	case "8h":
		return now.Add(-8 * time.Hour)
	case "1d":
		return now.AddDate(0, 0, -1)
	case "2d":
		return now.AddDate(0, 0, -2)
	case "7d":
		return now.AddDate(0, 0, -7)
	case "14d":
		return now.AddDate(0, 0, -14)
	case "30d":
		return now.AddDate(0, 0, -30)
	case "3month":
		return now.AddDate(0, -3, 0)
	case "6month":
		return now.AddDate(0, -6, 0)
	}
	return now.AddDate(0, 0, -1)
}

func GetMonthDays(month string) (int, error) {
	t, err := time.Parse("2006-01", month)
	if err != nil {
		return 0, err
	}
	year := t.Year()
	nextMonth := t.Month() + 1
	if nextMonth > time.December {
		nextMonth = time.January
		year++
	}

	nextMonthFirstDay := time.Date(year, nextMonth, 1, 0, 0, 0, 0, time.UTC)
	lastDay := nextMonthFirstDay.AddDate(0, 0, -1)
	return lastDay.Day(), nil
}
