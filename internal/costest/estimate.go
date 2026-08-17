package costest

import (
	"math"
	"sort"
	"time"
)

type Line struct {
	Service      string
	MonthlyCents int64
}

func DefaultLightsailCatalog() map[string]int64 {
	return map[string]int64{
		"nano_3_0": 500, "micro_3_0": 700, "small_3_0": 1200,
		"medium_3_0": 2400, "large_3_0": 4400, "xlarge_3_0": 8400,
		"2xlarge_3_0": 16400, "4xlarge_3_0": 38400,
	}
}

func LightsailMonthlyCents(bundleID string, catalog map[string]int64) int64 {
	if catalog == nil {
		return 0
	}
	return catalog[bundleID]
}

func StaticIPMonthlyCents(attached bool) int64 {
	if attached {
		return 0
	}
	return 300
}

func S3StandardMonthlyCents(bytes int64) int64 {
	if bytes <= 0 {
		return 0
	}
	gb := float64(bytes) / (1 << 30)
	return int64(math.Round(gb * 0.023 * 100))
}

func MonthToDateCents(monthly int64, now time.Time) int64 {
	days := daysInMonth(now)
	if days == 0 {
		return 0
	}
	return monthly * int64(now.Day()) / int64(days)
}

func daysInMonth(now time.Time) int {
	next := time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, 0, now.Location())
	return next.AddDate(0, 0, -1).Day()
}

func GroupByService(lines []Line) []Line {
	sums := map[string]int64{}
	for _, line := range lines {
		sums[line.Service] += line.MonthlyCents
	}
	out := make([]Line, 0, len(sums))
	for service, cents := range sums {
		out = append(out, Line{Service: service, MonthlyCents: cents})
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].MonthlyCents == out[j].MonthlyCents {
			return out[i].Service < out[j].Service
		}
		return out[i].MonthlyCents > out[j].MonthlyCents
	})
	return out
}

func BudgetBurnBps(spent, budget int64) int64 {
	if budget <= 0 {
		return 0
	}
	return spent * 10000 / budget
}

func SumCents(values []int64) int64 {
	var total int64
	for _, v := range values {
		total += v
	}
	return total
}
