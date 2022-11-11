package vtime

import "time"

func (v Time) Add(x int64, unit string) Time {
	t := tt(v)

	switch unit {
	case "ns", "nanosecond", "nanoseconds":
		t = t.Add(time.Duration(x) * time.Nanosecond)
	case "us", "microsecond", "microseconds":
		t = t.Add(time.Duration(x) * time.Microsecond)
	case "ms", "millisecond", "milliseconds":
		t = t.Add(time.Duration(x) * time.Millisecond)
	case "s", "second", "seconds":
		t = t.Add(time.Duration(x) * time.Second)
	case "m", "minute", "minutes":
		t = t.Add(time.Duration(x) * time.Minute)
	case "h", "hour", "hours":
		t = t.Add(time.Duration(x) * time.Hour)
	case "d", "day", "days":
		t = t.Add(time.Duration(x) * 24 * time.Hour)
	case "w", "week", "weeks":
		t = t.Add(time.Duration(x) * 7 * 24 * time.Hour)
	case "y", "year", "years":
		t = t.AddDate(int(x), 0, 0)
	case "M", "month", "months":
		t = t.AddDate(0, int(x), 0)
	default:
		return v
	}

	return ft(t)
}

func (v Time) Subtract(x int64, unit string) Time {
	return v.Add(-x, unit)
}
