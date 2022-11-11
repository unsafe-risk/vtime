package vtime

// Min returns the minimum of given times.
// If no times are given, it returns the current time.
func Min(tt ...Time) Time {
	if len(tt) == 0 {
		return Now()
	}

	m := tt[0]
	for i := 1; i < len(tt); i++ {
		if tt[i].Unix < m.Unix {
			m = tt[i]
		}

		if tt[i].Unix == m.Unix && tt[i].Nano < m.Nano {
			m = tt[i]
		}
	}

	return m
}

// Max returns the maximum of given times.
// If no times are given, it returns the current time.
func Max(tt ...Time) Time {
	if len(tt) == 0 {
		return Now()
	}

	m := tt[0]
	for i := 1; i < len(tt); i++ {
		if tt[i].Unix > m.Unix {
			m = tt[i]
		}

		if tt[i].Unix == m.Unix && tt[i].Nano > m.Nano {
			m = tt[i]
		}
	}

	return m
}

// After returns true if t is after u.
func After(t, u Time) bool {
	return t.Unix > u.Unix || (t.Unix == u.Unix && t.Nano > u.Nano)
}

// Before returns true if t is before u.
func Before(t, u Time) bool {
	return t.Unix < u.Unix || (t.Unix == u.Unix && t.Nano < u.Nano)
}
