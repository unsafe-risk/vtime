package vtime

import "time"

func (v *Time) Nanosecond(s ...int64) int64 {
	if len(s) >= 1 {
		*v = ft(time.Unix(int64(v.Unix), int64(s[0])))
	}
	return int64(v.Nano)
}

func (v *Time) Microsecond(s ...int64) int64 {
	const m = nano / micro
	if len(s) >= 1 {
		*v = ft(time.Unix(int64(v.Unix), int64(s[0])*m))
	}
	return int64(v.Nano) / m
}

func (v *Time) Millisecond(s ...int64) int64 {
	const m = nano / milli
	if len(s) >= 1 {
		*v = ft(time.Unix(int64(v.Unix), int64(s[0])*m))
	}
	return int64(v.Nano) / m
}

func (v *Time) Second(s ...int64) int64 {
	t := tt(*v)
	if len(s) >= 1 {
		YY := t.Year()
		MM := t.Month()
		DD := t.Day()
		HH := t.Hour()
		mm := t.Minute()
		ss := int(s[0])
		*v = ft(time.Date(YY, MM, DD, HH, mm, ss, int(v.Nano), t.Location()))
	}
	return int64(t.Second())
}

func (v *Time) Minute(s ...int64) int64 {
	t := tt(*v)
	if len(s) >= 1 {
		YY := t.Year()
		MM := t.Month()
		DD := t.Day()
		HH := t.Hour()
		mm := int(s[0])
		ss := t.Second()
		*v = ft(time.Date(YY, MM, DD, HH, mm, ss, int(v.Nano), t.Location()))
	}
	return int64(t.Minute())
}

func (v *Time) Hour(s ...int64) int64 {
	t := tt(*v)
	if len(s) >= 1 {
		YY := t.Year()
		MM := t.Month()
		DD := t.Day()
		HH := int(s[0])
		mm := t.Minute()
		ss := t.Second()
		*v = ft(time.Date(YY, MM, DD, HH, mm, ss, int(v.Nano), t.Location()))
	}
	return int64(t.Hour())
}

func (v *Time) Day(s ...int64) int64 {
	t := tt(*v)
	if len(s) >= 1 {
		YY := t.Year()
		MM := t.Month()
		DD := int(s[0])
		HH := t.Hour()
		mm := t.Minute()
		ss := t.Second()
		*v = ft(time.Date(YY, MM, DD, HH, mm, ss, int(v.Nano), t.Location()))
	}
	return int64(t.Day())
}

type Month = time.Month

func (v *Time) Month(s ...Month) Month {
	t := tt(*v)
	if len(s) >= 1 {
		YY := t.Year()
		MM := s[0]
		DD := t.Day()
		HH := t.Hour()
		mm := t.Minute()
		ss := t.Second()
		*v = ft(time.Date(YY, MM, DD, HH, mm, ss, int(v.Nano), t.Location()))
	}
	return t.Month()
}

func (v *Time) Year(s ...int64) int64 {
	t := tt(*v)
	if len(s) >= 1 {
		YY := int(s[0])
		MM := t.Month()
		DD := t.Day()
		HH := t.Hour()
		mm := t.Minute()
		ss := t.Second()
		*v = ft(time.Date(YY, MM, DD, HH, mm, ss, int(v.Nano), t.Location()))
	}
	return int64(t.Year())
}
