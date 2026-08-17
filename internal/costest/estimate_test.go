package costest

import (
	"testing"
	"time"
)

func TestLightsailMonthlyCents_smallBundle(t *testing.T) {
	got := LightsailMonthlyCents("small_3_0", DefaultLightsailCatalog())
	if got != 1200 {
		t.Fatalf("LightsailMonthlyCents(small_3_0) = %d, want 1200", got)
	}
}

func TestLightsailMonthlyCents_unknownBundleIsZero(t *testing.T) {
	got := LightsailMonthlyCents("does_not_exist", DefaultLightsailCatalog())
	if got != 0 {
		t.Fatalf("unknown bundle = %d, want 0", got)
	}
}

func TestStaticIPMonthlyCents_attachedIsFree(t *testing.T) {
	if got := StaticIPMonthlyCents(true); got != 0 {
		t.Fatalf("attached IP = %d, want 0", got)
	}
}

func TestStaticIPMonthlyCents_idleCostsThreeDollars(t *testing.T) {
	if got := StaticIPMonthlyCents(false); got != 300 {
		t.Fatalf("idle IP = %d, want 300", got)
	}
}

func TestS3StandardMonthlyCents_oneGiB(t *testing.T) {
	got := S3StandardMonthlyCents(1 << 30)
	if got != 2 {
		t.Fatalf("1 GiB S3 = %d cents, want 2 ($0.023 rounded)", got)
	}
}

func TestMonthToDateCents_proratesByDay(t *testing.T) {
	now := time.Date(2026, 8, 16, 12, 0, 0, 0, time.UTC)
	got := MonthToDateCents(3100, now)
	if got != 1600 {
		t.Fatalf("MTD of 3100 on Aug 16 = %d, want 1600", got)
	}
}

func TestGroupByService_sumsSameService(t *testing.T) {
	got := GroupByService([]Line{
		{Service: "Amazon Lightsail", MonthlyCents: 1200},
		{Service: "Amazon Simple Storage Service", MonthlyCents: 40},
		{Service: "Amazon Lightsail", MonthlyCents: 300},
	})
	if len(got) != 2 {
		t.Fatalf("groups = %d, want 2: %#v", len(got), got)
	}
	if got[0].Service != "Amazon Lightsail" || got[0].MonthlyCents != 1500 {
		t.Fatalf("first group = %+v, want Lightsail 1500", got[0])
	}
}

func TestBudgetBurnBps_halfSpent(t *testing.T) {
	if got := BudgetBurnBps(2500, 5000); got != 5000 {
		t.Fatalf("burn = %d bps, want 5000", got)
	}
}

func TestBudgetBurnBps_zeroBudget(t *testing.T) {
	if got := BudgetBurnBps(100, 0); got != 0 {
		t.Fatalf("zero budget burn = %d, want 0", got)
	}
}
