package handlers

import "fmt"

func formatUSD(cents int64) string {
	sign := ""
	if cents < 0 {
		sign = "-"
		cents = -cents
	}
	return fmt.Sprintf("%sUS$ %d,%02d", sign, cents/100, cents%100)
}

func resourceKindLabel(kind string) string {
	switch kind {
	case "lightsail_instance":
		return "Lightsail"
	case "lightsail_static_ip":
		return "IP estático"
	case "s3_bucket":
		return "S3"
	default:
		return kind
	}
}
