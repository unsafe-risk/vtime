package vtime

import "time"

func (v Time) SetNanosecond(s int64) Time {
	nv := ft(time.Unix(int64(v.Unix), int64(s)))
	return nv
}

func (v Time) Nanosecond() int64 {
	return int64(v.Nano)
}

func (v Time) SetMicrosecond(s int64) Time {
	const m = nano / micro
	nv := ft(time.Unix(int64(v.Unix), int64(s)*m))
	return nv
}

func (v Time) Microsecond() int64 {
	const m = nano / micro
	return int64(v.Nano) / m
}

func (v Time) SetMillisecond(s int64) Time {
	const m = nano / milli
	nv := ft(time.Unix(int64(v.Unix), int64(s)*m))
	return nv
}

func (v Time) Millisecond() int64 {
	const m = nano / milli
	return int64(v.Nano) / m
}

func (v Time) SetSecond(s int64) Time {
	t := tt(v)
	YY := t.Year()
	MM := t.Month()
	DD := t.Day()
	HH := t.Hour()
	mm := t.Minute()
	ss := int(s)
	nv := ft(time.Date(YY, MM, DD, HH, mm, ss, int(v.Nano), t.Location()))
	return nv
}

func (v Time) Second() int64 {
	t := tt(v)
	return int64(t.Second())
}

func (v Time) SetMinute(s int64) Time {
	t := tt(v)
	YY := t.Year()
	MM := t.Month()
	DD := t.Day()
	HH := t.Hour()
	mm := int(s)
	ss := t.Second()
	nv := ft(time.Date(YY, MM, DD, HH, mm, ss, int(v.Nano), t.Location()))
	return nv
}

func (v Time) Minute() int64 {
	t := tt(v)
	return int64(t.Minute())
}

func (v Time) SetHour(s int64) Time {
	t := tt(v)
	YY := t.Year()
	MM := t.Month()
	DD := t.Day()
	HH := int(s)
	mm := t.Minute()
	ss := t.Second()
	nv := ft(time.Date(YY, MM, DD, HH, mm, ss, int(v.Nano), t.Location()))
	return nv
}

func (v Time) Hour() int64 {
	t := tt(v)
	return int64(t.Hour())
}

func (v Time) SetDay(s int64) Time {
	t := tt(v)
	YY := t.Year()
	MM := t.Month()
	DD := int(s)
	HH := t.Hour()
	mm := t.Minute()
	ss := t.Second()
	nv := ft(time.Date(YY, MM, DD, HH, mm, ss, int(v.Nano), t.Location()))
	return nv
}

func (v Time) Day() int64 {
	t := tt(v)
	return int64(t.Day())
}

func (v Time) SetMonth(s Month) Time {
	t := tt(v)
	YY := t.Year()
	MM := s
	DD := t.Day()
	HH := t.Hour()
	mm := t.Minute()
	ss := t.Second()
	nv := ft(time.Date(YY, MM, DD, HH, mm, ss, int(v.Nano), t.Location()))
	return nv
}

func (v Time) Month() Month {
	t := tt(v)
	return Month(t.Month())
}

func (v Time) SetYear(s int64) Time {
	t := tt(v)
	YY := int(s)
	MM := t.Month()
	DD := t.Day()
	HH := t.Hour()
	mm := t.Minute()
	ss := t.Second()
	nv := ft(time.Date(YY, MM, DD, HH, mm, ss, int(v.Nano), t.Location()))
	return nv
}

func (v Time) Year() int64 {
	t := tt(v)
	return int64(t.Year())
}
