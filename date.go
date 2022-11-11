package vtime

import "time"

type Date struct {
	Year  int64
	Month Month
	Day   int64
}

func (v Time) SetDate(d Date) Time {
	t := tt(v)
	YY := int(d.Year)
	MM := time.Month(d.Month)
	DD := int(d.Day)
	HH := t.Hour()
	mm := t.Minute()
	ss := t.Second()
	nv := ft(time.Date(YY, MM, DD, HH, mm, ss, int(v.Nano), t.Location()))
	return nv
}

func (v Time) Date(tz *time.Location) Date {
	if tz == nil {
		tz = time.Local
	}
	t := tt(v).In(tz)
	return Date{
		Year:  int64(t.Year()),
		Month: Month(t.Month()),
		Day:   int64(t.Day()),
	}
}

func (v Date) Time(tz *time.Location) Time {
	if tz == nil {
		tz = time.Local
	}
	t := time.Date(int(v.Year), time.Month(v.Month), int(v.Day), 0, 0, 0, 0, tz)
	return ft(t)
}
