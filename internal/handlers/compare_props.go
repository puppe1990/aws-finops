package handlers

import (
	"sort"
	"time"

	"github.com/puppe1990/cais/pkg/cais/i18n"

	"github.com/puppe1990/aws-finops/internal/awsinv"
	"github.com/puppe1990/aws-finops/internal/models"
)

func MonthDeltaBps(curr, prev int64) (int64, bool) {
	if prev == 0 {
		return 0, false
	}
	return (curr - prev) * 10000 / prev, true
}

func compareMonthRows(months []awsinv.MonthCost, cat *i18n.Catalog, now time.Time) []map[string]any {
	now = time.Date(now.UTC().Year(), now.UTC().Month(), 1, 0, 0, 0, 0, time.UTC)
	n := len(months)
	out := make([]map[string]any, n)
	for i := n - 1; i >= 0; i-- {
		m := months[i]
		row := map[string]any{
			"query":   m.Query,
			"label":   ledgerMonthLabel(cat, awsinv.LedgerMonth{Period: m.Period}),
			"cents":   m.Cents,
			"usd":     formatUSD(m.Cents),
			"current": m.Period.Equal(now),
		}
		if i > 0 {
			if bps, ok := MonthDeltaBps(m.Cents, months[i-1].Cents); ok {
				row["deltaBps"] = bps
				row["deltaUSD"] = formatUSD(m.Cents - months[i-1].Cents)
			}
		}
		out[n-1-i] = row
	}
	return out
}

func compareServiceRows(current, previous []models.CostLine) []map[string]any {
	curr := serviceSums(current)
	prev := serviceSums(previous)
	names := make([]string, 0, len(curr)+len(prev))
	seen := map[string]bool{}
	for name := range curr {
		names = append(names, name)
		seen[name] = true
	}
	for name := range prev {
		if !seen[name] {
			names = append(names, name)
		}
	}
	sort.Slice(names, func(i, j int) bool {
		if curr[names[i]] == curr[names[j]] {
			return names[i] < names[j]
		}
		return curr[names[i]] > curr[names[j]]
	})
	out := make([]map[string]any, 0, len(names))
	for _, name := range names {
		c, p := curr[name], prev[name]
		if c == 0 && p == 0 {
			continue
		}
		row := map[string]any{
			"name":          name,
			"currentCents":  c,
			"previousCents": p,
			"currentUSD":    formatUSD(c),
			"previousUSD":   formatUSD(p),
		}
		if bps, ok := MonthDeltaBps(c, p); ok {
			row["deltaBps"] = bps
		}
		out = append(out, row)
	}
	return out
}

func serviceSums(lines []models.CostLine) map[string]int64 {
	out := map[string]int64{}
	for _, line := range lines {
		if line.Service == "" {
			continue
		}
		out[line.Service] += line.MonthlyCents
	}
	return out
}
