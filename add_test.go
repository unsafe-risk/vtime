package vtime_test

import (
	"testing"
	"time"

	"v8.run/go/vtime"
)

func TestAdd(t *testing.T) {
	// Fri Nov 11 2022 14:38:59 GMT+0000
	v := vtime.Unix(1668177539).UTC()

	t.Run("seconds", func(t *testing.T) {
		v0 := v.Add(10, "seconds")
		if v0.Second() != 9 {
			t.Errorf("Expected 9, got %d", v0.Second())
		}
		if v0.Minute() != 39 {
			t.Errorf("Expected 39, got %d", v0.Minute())
		}
		if v0.Hour() != 14 {
			t.Errorf("Expected 14, got %d", v0.Hour())
		}
	})

	t.Run("minutes", func(t *testing.T) {
		v1 := v.Add(10, "minutes")
		if v1.Second() != 59 {
			t.Errorf("Expected 59, got %d", v1.Second())
		}
		if v1.Minute() != 48 {
			t.Errorf("Expected 48, got %d", v1.Minute())
		}
		if v1.Hour() != 14 {
			t.Errorf("Expected 14, got %d", v1.Hour())
		}
	})

	t.Run("hours", func(t *testing.T) {
		v2 := v.Add(10, "hours")
		if v2.Second() != 59 {
			t.Errorf("Expected 59, got %d", v2.Second())
		}
		if v2.Minute() != 38 {
			t.Errorf("Expected 38, got %d", v2.Minute())
		}
		if v2.Hour() != 0 {
			t.Errorf("Expected 0, got %d", v2.Hour())
		}

		d2 := v2.Date()
		if d2.Year != 2022 {
			t.Errorf("Expected 2022, got %d", d2.Year)
		}
		if d2.Month != time.November {
			t.Errorf("Expected 11, got %d", d2.Month)
		}
		if d2.Day != 12 {
			t.Errorf("Expected 12, got %d", d2.Day)
		}
	})

	t.Run("days", func(t *testing.T) {
		v3 := v.Add(10, "days")
		if v3.Second() != 59 {
			t.Errorf("Expected 59, got %d", v3.Second())
		}
		if v3.Minute() != 38 {
			t.Errorf("Expected 38, got %d", v3.Minute())
		}
		if v3.Hour() != 14 {
			t.Errorf("Expected 14, got %d", v3.Hour())
		}

		d3 := v3.Date()
		if d3.Year != 2022 {
			t.Errorf("Expected 2022, got %d", d3.Year)
		}
		if d3.Month != time.November {
			t.Errorf("Expected 11, got %d", d3.Month)
		}
		if d3.Day != 21 {
			t.Errorf("Expected 21, got %d", d3.Day)
		}
	})

	t.Run("weeks", func(t *testing.T) {
		v4 := v.Add(1, "weeks")
		if v4.Second() != 59 {
			t.Errorf("Expected 59, got %d", v4.Second())
		}
		if v4.Minute() != 38 {
			t.Errorf("Expected 38, got %d", v4.Minute())
		}
		if v4.Hour() != 14 {
			t.Errorf("Expected 14, got %d", v4.Hour())
		}

		d4 := v4.Date()
		if d4.Year != 2022 {
			t.Errorf("Expected 2022, got %d", d4.Year)
		}
		if d4.Month != time.November {
			t.Errorf("Expected 11, got %d", d4.Month)
		}
		if d4.Day != 18 {
			t.Errorf("Expected 18, got %d", d4.Day)
		}
	})

	t.Run("months", func(t *testing.T) {
		v5 := v.Add(1, "months")
		if v5.Second() != 59 {
			t.Errorf("Expected 59, got %d", v5.Second())
		}
		if v5.Minute() != 38 {
			t.Errorf("Expected 38, got %d", v5.Minute())
		}
		if v5.Hour() != 14 {
			t.Errorf("Expected 14, got %d", v5.Hour())
		}

		d5 := v5.Date()
		if d5.Year != 2022 {
			t.Errorf("Expected 2022, got %d", d5.Year)
		}
		if d5.Month != time.December {
			t.Errorf("Expected 12, got %d", d5.Month)
		}
		if d5.Day != 11 {
			t.Errorf("Expected 11, got %d", d5.Day)
		}
	})

	t.Run("years", func(t *testing.T) {
		v6 := v.Add(1, "year")
		if v6.Second() != 59 {
			t.Errorf("Expected 59, got %d", v6.Second())
		}
		if v6.Minute() != 38 {
			t.Errorf("Expected 38, got %d", v6.Minute())
		}
		if v6.Hour() != 14 {
			t.Errorf("Expected 14, got %d", v6.Hour())
		}

		d6 := v6.Date()
		if d6.Year != 2023 {
			t.Errorf("Expected 2023, got %d", d6.Year)
		}
		if d6.Month != time.November {
			t.Errorf("Expected 11, got %d", d6.Month)
		}
		if d6.Day != 11 {
			t.Errorf("Expected 11, got %d", d6.Day)
		}
	})

	t.Run("invalid", func(t *testing.T) {
		v7 := v.Add(-100000, "invalid")
		if v7.Equal(v) == false {
			t.Errorf("Expected %v, got %v", v, v7)
		}
	})
}

func TestSub(t *testing.T) {
	// Fri Nov 11 2022 14:38:59 GMT+0000
	v := vtime.Unix(1668177539).UTC()

	t.Run("seconds", func(t *testing.T) {
		v1 := v.Subtract(10, "seconds")
		if v1.Second() != 49 {
			t.Errorf("Expected 49, got %d", v1.Second())
		}
		if v1.Minute() != 38 {
			t.Errorf("Expected 38, got %d", v1.Minute())
		}
		if v1.Hour() != 14 {
			t.Errorf("Expected 14, got %d", v1.Hour())
		}

		d1 := v1.Date()
		if d1.Year != 2022 {
			t.Errorf("Expected 2022, got %d", d1.Year)
		}
		if d1.Month != time.November {
			t.Errorf("Expected 11, got %d", d1.Month)
		}
		if d1.Day != 11 {
			t.Errorf("Expected 11, got %d", d1.Day)
		}
	})

	t.Run("minutes", func(t *testing.T) {
		v2 := v.Subtract(10, "minutes")
		if v2.Second() != 59 {
			t.Errorf("Expected 59, got %d", v2.Second())
		}
		if v2.Minute() != 28 {
			t.Errorf("Expected 28, got %d", v2.Minute())
		}
		if v2.Hour() != 14 {
			t.Errorf("Expected 14, got %d", v2.Hour())
		}

		d2 := v2.Date()
		if d2.Year != 2022 {
			t.Errorf("Expected 2022, got %d", d2.Year)
		}
		if d2.Month != time.November {
			t.Errorf("Expected 11, got %d", d2.Month)
		}
		if d2.Day != 11 {
			t.Errorf("Expected 11, got %d", d2.Day)
		}
	})

	t.Run("hours", func(t *testing.T) {
		v3 := v.Subtract(10, "hours")
		if v3.Second() != 59 {
			t.Errorf("Expected 59, got %d", v3.Second())
		}
		if v3.Minute() != 38 {
			t.Errorf("Expected 38, got %d", v3.Minute())
		}
		if v3.Hour() != 4 {
			t.Errorf("Expected 4, got %d", v3.Hour())
		}

		d3 := v3.Date()
		if d3.Year != 2022 {
			t.Errorf("Expected 2022, got %d", d3.Year)
		}
		if d3.Month != time.November {
			t.Errorf("Expected 11, got %d", d3.Month)
		}
		if d3.Day != 11 {
			t.Errorf("Expected 11, got %d", d3.Day)
		}
	})

	t.Run("days", func(t *testing.T) {
		v4 := v.Subtract(10, "days")
		if v4.Second() != 59 {
			t.Errorf("Expected 59, got %d", v4.Second())
		}
		if v4.Minute() != 38 {
			t.Errorf("Expected 38, got %d", v4.Minute())
		}
		if v4.Hour() != 14 {
			t.Errorf("Expected 14, got %d", v4.Hour())
		}

		d4 := v4.Date()
		if d4.Year != 2022 {
			t.Errorf("Expected 2022, got %d", d4.Year)
		}
		if d4.Month != time.November {
			t.Errorf("Expected 11, got %d", d4.Month)
		}
		if d4.Day != 1 {
			t.Errorf("Expected 1, got %d", d4.Day)
		}
	})

	t.Run("months", func(t *testing.T) {
		v5 := v.Subtract(10, "months")
		if v5.Second() != 59 {
			t.Errorf("Expected 59, got %d", v5.Second())
		}
		if v5.Minute() != 38 {
			t.Errorf("Expected 38, got %d", v5.Minute())
		}
		if v5.Hour() != 14 {
			t.Errorf("Expected 14, got %d", v5.Hour())
		}

		d5 := v5.Date()
		if d5.Year != 2022 {
			t.Errorf("Expected 2022, got %d", d5.Year)
		}
		if d5.Month != time.January {
			t.Errorf("Expected 1, got %d", d5.Month)
		}
		if d5.Day != 11 {
			t.Errorf("Expected 11, got %d", d5.Day)
		}
	})

	t.Run("years", func(t *testing.T) {
		v6 := v.Subtract(10, "years")
		if v6.Second() != 59 {
			t.Errorf("Expected 59, got %d", v6.Second())
		}
		if v6.Minute() != 38 {
			t.Errorf("Expected 38, got %d", v6.Minute())
		}
		if v6.Hour() != 14 {
			t.Errorf("Expected 14, got %d", v6.Hour())
		}

		d6 := v6.Date()
		if d6.Year != 2012 {
			t.Errorf("Expected 2012, got %d", d6.Year)
		}
		if d6.Month != time.November {
			t.Errorf("Expected 11, got %d", d6.Month)
		}
		if d6.Day != 11 {
			t.Errorf("Expected 11, got %d", d6.Day)
		}
	})

	t.Run("invalid", func(t *testing.T) {
		v7 := v.Subtract(10000, "invalid")
		if !v7.Equal(v) {
			t.Errorf("Expected %v, got %v", v, v7)
		}
	})
}

func TestAddAddDuration(t *testing.T) {
	// Fri Nov 11 2022 14:38:59 GMT+0000
	v := vtime.Unix(1668177539).UTC()

	v0 := v.AddDuration(10 * time.Second)
	if v0.Second() != 9 {
		t.Errorf("Expected 9, got %d", v0.Second())
	}
	if v0.Minute() != 39 {
		t.Errorf("Expected 39, got %d", v0.Minute())
	}
	if v0.Hour() != 14 {
		t.Errorf("Expected 14, got %d", v0.Hour())
	}

	v1 := vtime.Unix(1668177539 + 10).UTC()
	if !v0.Equal(v1) {
		t.Errorf("Expected %v, got %v", v1, v0)
	}
}
