package vtime

import (
	"time"

	"v8.run/go/vtime/internal/parse"
)

type WallClock struct{}

type Parseable interface {
	time.Time | int64 | uint64 | string | WallClock
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
		if len(t) == 1 {
			unix, nano, _ := parse.Parse8601(*v)
			return Time{Unix: unix, Nano: nano, TZ: time.Local}
		} else {
			// TODO: Support custom time formats
			return Now()
		}
	case *WallClock:
		return Now()
	}

	// Default to Now()
	return Now()
}

func UTC[T Parseable](t ...T) Time {
	return VTime(t...).UTC()
}

// January=1...December=12
type Month = time.Month

func (v Time) Time() time.Time {
	return tt(v)
}

func (v Time) In(tz *time.Location) Time {
	if tz == nil {
		tz = time.Local
	}
	v.TZ = tz
	return v
}

func (v Time) UTC() Time {
	v.TZ = time.UTC
	return v
}

func (v Time) Local() Time {
	v.TZ = time.Local
	return v
}

func (v Time) String() string {
	return tt(v).Format(time.RFC3339Nano)
}
