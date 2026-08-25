package handlers

import (
	"testing"
	"time"

	"github.com/puppe1990/aws-finops/internal/awsinv"
	appi18n "github.com/puppe1990/aws-finops/internal/i18n"
	"github.com/puppe1990/aws-finops/internal/models"
)

func TestMonthDeltaBps(t *testing.T) {
	bps, ok := MonthDeltaBps(3215, 5228)
	if !ok || bps != (3215-5228)*10000/5228 {
		t.Fatalf("bps=%d ok=%v", bps, ok)
	}
	if _, ok := MonthDeltaBps(100, 0); ok {
		t.Fatal("prev zero should be skipped")
	}
}

func TestCompareMonthRows_newestFirstWithDelta(t *testing.T) {
	jul := time.Date(2026, 7, 1, 0, 0, 0, 0, time.UTC)
	aug := time.Date(2026, 8, 1, 0, 0, 0, 0, time.UTC)
	rows := compareMonthRows([]awsinv.MonthCost{
		{Query: "2026-07", Period: jul, Cents: 5228},
		{Query: "2026-08", Period: aug, Cents: 3215},
	}, appi18n.DefaultCatalog(), aug)
	if len(rows) != 2 {
		t.Fatalf("len=%d", len(rows))
	}
	if rows[0]["query"] != "2026-08" || rows[0]["usd"] != "US$ 32,15" {
		t.Fatalf("current=%v", rows[0])
	}
	if rows[0]["current"] != true {
		t.Fatal("august should be current")
	}
	want := int64((3215 - 5228) * 10000 / 5228)
	if rows[0]["deltaBps"] != want {
		t.Fatalf("deltaBps=%v want %d", rows[0]["deltaBps"], want)
	}
	if rows[1]["query"] != "2026-07" || rows[1]["deltaBps"] != nil {
		t.Fatalf("july=%v", rows[1])
	}
}

func TestCompareServiceRows_pairsCurrentAndPrevious(t *testing.T) {
	rows := compareServiceRows(
		[]models.CostLine{
			{Service: "Amazon Lightsail", MonthlyCents: 1983},
			{Service: "Amazon S3", MonthlyCents: 164},
		},
		[]models.CostLine{
			{Service: "Amazon Lightsail", MonthlyCents: 3220},
		},
	)
	if len(rows) != 2 {
		t.Fatalf("len=%d", len(rows))
	}
	if rows[0]["name"] != "Amazon Lightsail" {
		t.Fatalf("top=%v", rows[0])
	}
	if rows[0]["currentUSD"] != "US$ 19,83" || rows[0]["previousUSD"] != "US$ 32,20" {
		t.Fatalf("lightsail=%v", rows[0])
	}
	if rows[1]["name"] != "Amazon S3" || rows[1]["previousCents"] != int64(0) {
		t.Fatalf("s3=%v", rows[1])
	}
}
