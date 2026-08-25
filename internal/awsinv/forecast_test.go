package awsinv

import "testing"

func TestForecastCentsForPeriod_picksMatchingMonth(t *testing.T) {
	cents := ForecastCentsForPeriod("2026-09-01", []ForecastBucket{
		{Start: "2026-08-24", Amount: "12.00"},
		{Start: "2026-09-01", Amount: "38.50"},
	}, "50.50")
	if cents != 3850 {
		t.Fatalf("cents = %d", cents)
	}
}

func TestForecastCentsForPeriod_matchesDatePrefix(t *testing.T) {
	cents := ForecastCentsForPeriod("2026-09-01", []ForecastBucket{
		{Start: "2026-09-01T00:00:00Z", Amount: "38.50"},
	}, "99.00")
	if cents != 3850 {
		t.Fatalf("cents = %d", cents)
	}
}

func TestForecastCentsForPeriod_skipsTotalWhenMonthMissing(t *testing.T) {
	cents := ForecastCentsForPeriod("2026-09-01", []ForecastBucket{
		{Start: "2026-08-24", Amount: "12.00"},
	}, "38.50")
	if cents != 0 {
		t.Fatalf("cents = %d", cents)
	}
}

func TestForecastCentsForPeriod_usesTotalWhenQueryIsThatMonth(t *testing.T) {
	cents := ForecastCentsForPeriod("2026-09-01", nil, "38.50")
	if cents != 3850 {
		t.Fatalf("cents = %d", cents)
	}
}
