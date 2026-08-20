package awsinv

import (
	"fmt"
	"time"
)

const LedgerMonthLayout = "2006-01"
const LedgerMonthLookback = 12

type LedgerMonth struct {
	Period    time.Time
	IsCurrent bool
	Prev      string
	Next      string
	Query     string
}

func MonthBounds(t time.Time) (string, string) {
	t = t.UTC()
	start := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC)
	return start.Format("2006-01-02"), start.AddDate(0, 1, 0).Format("2006-01-02")
}

func MonthLabelKey(t time.Time) string {
	return fmt.Sprintf("dash.m%02d", int(t.UTC().Month()))
}

func ParseLedgerMonth(raw string, now time.Time) LedgerMonth {
	now = time.Date(now.UTC().Year(), now.UTC().Month(), 1, 0, 0, 0, 0, time.UTC)
	floor := now.AddDate(0, -LedgerMonthLookback, 0)
	period := now
	if parsed, err := time.Parse(LedgerMonthLayout, raw); err == nil {
		parsed = time.Date(parsed.Year(), parsed.Month(), 1, 0, 0, 0, 0, time.UTC)
		if !parsed.After(now) && !parsed.Before(floor) {
			period = parsed
		}
	}
	lm := LedgerMonth{
		Period:    period,
		IsCurrent: period.Equal(now),
		Query:     period.Format(LedgerMonthLayout),
	}
	if period.After(floor) {
		lm.Prev = period.AddDate(0, -1, 0).Format(LedgerMonthLayout)
	}
	if !lm.IsCurrent {
		lm.Next = period.AddDate(0, 1, 0).Format(LedgerMonthLayout)
	}
	return lm
}
