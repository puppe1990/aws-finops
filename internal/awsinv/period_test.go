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
