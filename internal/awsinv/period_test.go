package awsinv

import (
	"testing"
	"time"
)

func TestMonthBounds_july2026(t *testing.T) {
	start, end := MonthBounds(time.Date(2026, 7, 15, 12, 0, 0, 0, time.UTC))
	if start != "2026-07-01" || end != "2026-08-01" {
		t.Fatalf("bounds = %s %s", start, end)
	}
}

func TestParseLedgerMonth(t *testing.T) {
	now := time.Date(2026, 8, 20, 0, 0, 0, 0, time.UTC)

	cur := ParseLedgerMonth("", now)
	if !cur.IsCurrent || cur.Query != "2026-08" || cur.Prev != "2026-07" || cur.Next != "" {
		t.Fatalf("empty = %#v", cur)
	}

	jul := ParseLedgerMonth("2026-07", now)
	if jul.IsCurrent || jul.Query != "2026-07" || jul.Prev != "2026-06" || jul.Next != "2026-08" {
		t.Fatalf("july = %#v", jul)
	}

	fut := ParseLedgerMonth("2026-09", now)
	if !fut.IsCurrent || fut.Query != "2026-08" {
		t.Fatalf("future = %#v", fut)
	}

	old := ParseLedgerMonth("2025-07", now) // 13 months back
	if !old.IsCurrent || old.Query != "2026-08" {
		t.Fatalf("too old = %#v", old)
	}

	floor := ParseLedgerMonth("2025-08", now) // exactly 12 months
	if floor.IsCurrent || floor.Query != "2025-08" || floor.Prev != "" {
		t.Fatalf("floor = %#v", floor)
	}
}

func TestMonthLabelKey(t *testing.T) {
	if MonthLabelKey(time.Date(2026, 8, 1, 0, 0, 0, 0, time.UTC)) != "dash.m08" {
		t.Fatal(MonthLabelKey(time.Date(2026, 8, 1, 0, 0, 0, 0, time.UTC)))
	}
}

func TestNextMonth(t *testing.T) {
	got := NextMonth(time.Date(2026, 8, 24, 15, 4, 0, 0, time.UTC))
	want := time.Date(2026, 9, 1, 0, 0, 0, 0, time.UTC)
	if !got.Equal(want) {
		t.Fatalf("aug = %v want %v", got, want)
	}

	got = NextMonth(time.Date(2026, 12, 31, 23, 0, 0, 0, time.UTC))
	want = time.Date(2027, 1, 1, 0, 0, 0, 0, time.UTC)
	if !got.Equal(want) {
		t.Fatalf("dec = %v want %v", got, want)
	}
}

func TestForecastAPIStart_clampsFutureMonthToToday(t *testing.T) {
	today := time.Date(2026, 8, 24, 15, 0, 0, 0, time.UTC)
	if got := ForecastAPIStart("2026-09-01", today); got != "2026-08-24" {
		t.Fatalf("next month start = %s", got)
	}
	if got := ForecastAPIStart("2026-08-01", today); got != "2026-08-01" {
		t.Fatalf("current month start = %s", got)
	}
}
