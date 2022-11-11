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

// Equal returns true if t is equal to u.
func (v Time) Equal(u Time) bool {
	return v.Unix == u.Unix && v.Nano == u.Nano
}

// Equal returns true if t is equal to u.
func (v Date) Equal(u Date) bool {
	return v.Year == u.Year && v.Month == u.Month && v.Day == u.Day
}
