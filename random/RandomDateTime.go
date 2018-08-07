package random

import (
	"time"
)

type TRandomDateTime struct{}

var RandomDateTime *TRandomDateTime = &TRandomDateTime{}

func (c *TRandomDateTime) NextDate(min time.Time, max time.Time) time.Time {
	diff := max.Unix() - min.Unix()
	if diff <= 0 {
		return min
	}

	diff = RandomLong.NextLong(0, diff)
	date := min.Add(time.Duration(diff) * time.Second)

	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, min.Location())
}

func (c *TRandomDateTime) NextDateTime(min time.Time, max time.Time) time.Time {
	diff := max.Unix() - min.Unix()
	if diff <= 0 {
		return min
	}

	diff = RandomLong.NextLong(0, diff)
	return min.Add(time.Duration(diff) * time.Second)
}

func (c *TRandomDateTime) UpdateDateTime(value time.Time, interval int64) time.Time {
	if interval == 0 {
		interval = 10 * 24 * 3600
	}
	if interval < 0 {
		return value
	}

	// Days to milliseconds
	unixTime := value.Unix() + RandomLong.NextLong(-interval, interval)
	return time.Unix(unixTime, 0)
}
