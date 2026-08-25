package awsinv

import "strings"

type ForecastBucket struct {
	Start  string
	Amount string
}

func ForecastCentsForPeriod(monthStart string, buckets []ForecastBucket, total string) int64 {
	for _, b := range buckets {
		if strings.HasPrefix(b.Start, monthStart) {
			return parseUSDCents(b.Amount)
		}
	}
	if len(buckets) == 0 {
		return parseUSDCents(total)
	}
	return 0
}
