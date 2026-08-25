package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/puppe1990/cais/pkg/cais"
	"github.com/puppe1990/cais/pkg/cais/session"

	"github.com/puppe1990/aws-finops/internal/finops"
	"github.com/puppe1990/aws-finops/internal/models"
	"github.com/puppe1990/aws-finops/internal/seed"
	"github.com/puppe1990/aws-finops/internal/syncer"
)

func TestCompareHandler_InertiaComponent(t *testing.T) {
	h := NewCompareHandler(setupTestStore(t), testSite(), cais.Config{}, setupTestInertia(t))
	req := inertiaRequest(http.MethodGet, "/compare", nil)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("status=%d", rr.Code)
	}
	assertInertiaComponent(t, rr, "Compare")
}

func TestCompareHandler_monthlyRowsFromCostExplorer(t *testing.T) {
	s := setupTestStore(t)
	uid, err := s.CreateUser("ops@example.com", "hash")
	if err != nil {
		t.Fatal(err)
	}
	t.Setenv(finops.SeedAccountEnv, "111111111111")
	if err := seed.EnsurePrimaryWorkspace(s, uid); err != nil {
		t.Fatal(err)
	}

	col := stubRangeCollector{rangeLines: []models.CostLine{
		{Service: "Amazon Lightsail", MonthlyCents: 3220, Source: finops.SourceCE, PeriodStart: "2026-07-01", PeriodEnd: "2026-08-01"},
		{Service: "Amazon Lightsail", MonthlyCents: 1983, Source: finops.SourceCE, PeriodStart: "2026-08-01", PeriodEnd: "2026-09-01"},
	}}
	h := NewCompareHandler(s, testSite(), cais.Config{}, setupTestInertia(t)).
		WithSyncer(syncer.New(s, col))
	h.now = func() time.Time { return time.Date(2026, 8, 20, 0, 0, 0, 0, time.UTC) }

	req := inertiaRequest(http.MethodGet, "/compare", nil)
	req = session.WithUserID(req, uid)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("status=%d body=%s", rr.Code, rr.Body.String())
	}
	assertInertiaComponent(t, rr, "Compare")
	months, _ := assertInertiaProp(t, rr, "months").([]any)
	if len(months) != 13 {
		t.Fatalf("months=%d", len(months))
	}
	first := months[0].(map[string]any)
	if first["query"] != "2026-08" || first["usd"] != "US$ 19,83" {
		t.Fatalf("current=%v", first)
	}
	services, _ := assertInertiaProp(t, rr, "services").([]any)
	if len(services) == 0 {
		t.Fatal("expected service compare rows")
	}
	svc := services[0].(map[string]any)
	if svc["name"] != "Amazon Lightsail" || svc["currentUSD"] != "US$ 19,83" {
		t.Fatalf("service=%v", svc)
	}
	series, _ := svc["months"].([]any)
	if len(series) < 2 {
		t.Fatalf("service months=%d", len(series))
	}
	firstMonth := series[0].(map[string]any)
	if firstMonth["query"] != "2026-07" {
		t.Fatalf("first service month=%v", firstMonth)
	}
}
