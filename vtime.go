package vtime

import "time"

type Parseable interface {
	time.Time | int64 | uint64 | string
}

type Time struct {
	Unix int64
	Nano int64
	TZ   *time.Location
}

func ft(t time.Time) Time {
	unix := t.Unix()
	nano := t.Nanosecond()
	return Time{Unix: unix, Nano: int64(nano), TZ: t.Location()}
}

func tt(v Time) time.Time {
	if v.TZ == nil {
		v.TZ = time.Local
	}
	return time.Unix(int64(v.Unix), int64(v.Nano)).In(v.TZ)
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

// January=1...December=12
type Month = time.Month

func (v Time) Time() time.Time {
	return tt(v)
}

func (v Time) UTC() Time {
	v.TZ = time.UTC
	return v
}
