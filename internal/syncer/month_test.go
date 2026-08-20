package syncer

import (
	"context"
	"testing"
	"time"

	"github.com/puppe1990/aws-finops/internal/awsinv"
	"github.com/puppe1990/aws-finops/internal/finops"
	"github.com/puppe1990/aws-finops/internal/models"
	"github.com/puppe1990/aws-finops/internal/store"
)

type stubMonthCollector struct {
	stubCollector
	lines []models.CostLine
}

func (s stubMonthCollector) CostForMonth(_ context.Context, _ awsinv.Credentials, _ time.Time) ([]models.CostLine, error) {
	return s.lines, nil
}

func TestSyncer_CostForMonth_doesNotWriteCostLines(t *testing.T) {
	s, err := store.NewSQLiteStore(":memory:", "test")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = s.Close() })
	tid, _ := s.CreateTenant("Demo", "demo")
	accID, _ := s.CreateCloudAccount(models.CloudAccount{
		TenantID: tid, AWSAccountID: "111111111111", Alias: "p",
		Region: "us-east-1", AuthMode: finops.AuthModeDefaultChain,
	})
	_ = s.ReplaceCostLines(accID, []models.CostLine{{
		Service: "Stored", MonthlyCents: 100, Source: finops.SourceCE,
		PeriodStart: "2026-08-01", PeriodEnd: "2026-09-01",
	}})

	july := []models.CostLine{{
		Service: "Amazon Lightsail", MonthlyCents: 1947, Source: finops.SourceCE,
		PeriodStart: "2026-07-01", PeriodEnd: "2026-08-01",
	}}
	got, err := New(s, stubMonthCollector{lines: july}).CostForMonth(
		context.Background(), tid, time.Date(2026, 7, 1, 0, 0, 0, 0, time.UTC))
	if err != nil || len(got) != 1 || got[0].MonthlyCents != 1947 {
		t.Fatalf("got=%#v err=%v", got, err)
	}
	stored, _ := s.ListCostLines(accID)
	if len(stored) != 1 || stored[0].Service != "Stored" {
		t.Fatalf("sqlite mutated: %#v", stored)
	}
}
