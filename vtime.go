package vtime

import "time"

type Parseable interface {
	time.Time | int64 | uint64 | string
}

type Time struct {
	Unix uint64
	Nano uint64
}

func FromTime(t time.Time) Time {
	unix := t.Unix()
	nano := t.Nanosecond()
	return Time{Unix: uint64(unix), Nano: uint64(nano)}
}

func Now() Time {
	now := time.Now()
	return FromTime(now)
}

func VTime[T Parseable](t ...T) Time {
	if len(t) <= 0 {
		return Now() // IF NO ARGUMENTS ARE PASSED, RETURN CURRENT TIME
	}

	switch v := any(&t[0]).(type) {
	case *time.Time:
		return FromTime(*v)
	case *int64: // UNIX TIME
		t := time.Unix(*v, 0)
		return FromTime(t)
	case *uint64: // UNIX TIME
		t := time.Unix(int64(*v), 0)
		return FromTime(t)
	}

	// Default to Now()
	return Now()
}
