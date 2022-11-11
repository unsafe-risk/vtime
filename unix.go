package vtime

import "time"

const (
	nano  = 1e9
	micro = 1e6
	milli = 1e3
)

func Unix(ts int64) Time {
	return ft(time.Unix(ts, 0))
}

func UnixMilli(ts int64) Time {
	return ft(time.UnixMilli(ts))
}

func UnixMicro(ts int64) Time {
	return ft(time.UnixMicro(ts))
}

func UnixNano(ts int64) Time {
	sec := ts / nano
	nano := ts % nano
	return ft(time.Unix(sec, nano))
}
