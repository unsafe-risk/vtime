package vtime

import "time"

type Parseable interface {
	time.Time | int64 | uint64 | string
}

type Time struct {
	Unix int64
	Nano uint64
}

func ft(t time.Time) Time {
	unix := t.Unix()
	nano := t.Nanosecond()
	return Time{Unix: unix, Nano: uint64(nano)}
}

func tt(v Time) time.Time {
	return time.Unix(int64(v.Unix), int64(v.Nano))
}

func Now() Time {
	now := time.Now()
	return ft(now)
}

func VTime[T Parseable](t ...T) Time {
	if len(t) <= 0 {
		return Now() // IF NO ARGUMENTS ARE PASSED, RETURN CURRENT TIME
	}

	switch v := any(&t[0]).(type) {
	case *time.Time:
		return ft(*v)
	case *int64: // UNIX TIME MILLISECONDS
		return UnixMilli(*v)
	case *uint64: // UNIX TIME MILLISECONDS
		return UnixMilli(int64(*v))
	case *string:
		// TODO: PARSE STRING
	}

	// Default to Now()
	return Now()
}

func (v *Time) Time() time.Time {
	return tt(*v)
}
