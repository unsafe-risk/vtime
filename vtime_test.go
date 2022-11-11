package vtime

import (
	"testing"
	"time"
	_ "time/tzdata"
)

func TestUnix(t *testing.T) {
	v := Unix(1668174735).UTC() // Fri Nov 11 2022 13:52:15 GMT+0000
	if v.Year() != 2022 {
		t.Errorf("Expected 2022, got %d", v.Year())
	}
	if v.Month() != time.November {
		t.Errorf("Expected November, got %s", v.Month())
	}
	if v.Day() != 11 {
		t.Errorf("Expected 11, got %d", v.Day())
	}
	if v.Hour() != 13 {
		t.Errorf("Expected 13, got %d", v.Hour())
	}
	if v.Minute() != 52 {
		t.Errorf("Expected 52, got %d", v.Minute())
	}
	if v.Second() != 15 {
		t.Errorf("Expected 15, got %d", v.Second())
	}

	Seoul, err := time.LoadLocation("Asia/Seoul") // "time/tzdata" is required
	if err != nil {
		t.Error(err)
	}
	v = v.In(Seoul) // 	Fri Nov 11 2022 22:52:15 GMT+0900 (한국 표준시)
	if v.Hour() != 22 {
		t.Error("Expected 22, got", v.Hour())
	}
	if v.Minute() != 52 {
		t.Error("Expected 52, got", v.Minute())
	}
	if v.Second() != 15 {
		t.Error("Expected 15, got", v.Second())
	}
	if v.Nanosecond() != 0 {
		t.Error("Expected 0, got", v.Nanosecond())
	}
	d := v.Date() // 2022-11-11
	if d.Year != 2022 {
		t.Error("Expected 2022, got", d.Year)
	}
	if d.Month != time.November {
		t.Error("Expected 11, got", d.Month)
	}
	if d.Day != 11 {
		t.Error("Expected 11, got", d.Day)
	}
}

func TestSet(t *testing.T) {
	// Fri Nov 11 2022 13:52:15 GMT+0000
	v := Now().UTC().
		SetYear(2022).
		SetMonth(time.November).
		SetDay(11).
		SetHour(13).
		SetMinute(52).
		SetSecond(15)
	if v.Year() != 2022 {
		t.Errorf("Expected 2022, got %d", v.Year())
	}
	if v.Month() != time.November {
		t.Errorf("Expected November, got %s", v.Month())
	}
	if v.Day() != 11 {
		t.Errorf("Expected 11, got %d", v.Day())
	}
	if v.Hour() != 13 {
		t.Errorf("Expected 13, got %d", v.Hour())
	}
	if v.Minute() != 52 {
		t.Errorf("Expected 52, got %d", v.Minute())
	}
	if v.Second() != 15 {
		t.Errorf("Expected 15, got %d", v.Second())
	}

	Seoul, err := time.LoadLocation("Asia/Seoul") // "time/tzdata" is required
	if err != nil {
		t.Error(err)
	}
	v = v.In(Seoul) // 	Fri Nov 11 2022 22:52:15 GMT+0900 (한국 표준시)
	if v.Hour() != 22 {
		t.Error("Expected 22, got", v.Hour())
	}
	if v.Minute() != 52 {
		t.Error("Expected 52, got", v.Minute())
	}
	if v.Second() != 15 {
		t.Error("Expected 15, got", v.Second())
	}
	d := v.Date() // 2022-11-11
	if d.Year != 2022 {
		t.Error("Expected 2022, got", d.Year)
	}
	if d.Month != time.November {
		t.Error("Expected 11, got", d.Month)
	}
	if d.Day != 11 {
		t.Error("Expected 11, got", d.Day)
	}
}
