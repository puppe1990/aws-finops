package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/puppe1990/cais/pkg/cais"
	"github.com/puppe1990/cais/pkg/cais/flash"
	"github.com/puppe1990/cais/pkg/cais/session"

	"github.com/puppe1990/aws-finops/internal/finops"
	"github.com/puppe1990/aws-finops/internal/models"
	"github.com/puppe1990/aws-finops/internal/seed"
	"github.com/puppe1990/aws-finops/internal/syncer"
)

func TestDashboardHandler_InertiaComponent(t *testing.T) {
	h := NewDashboardHandler(setupTestRenderer(t), setupTestStore(t), testSite(), cais.Config{}, setupTestInertia(t))

	req := inertiaRequest(http.MethodGet, "/dashboard", nil)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("status = %d, want %d", rr.Code, http.StatusOK)
	}
	assertInertiaComponent(t, rr, "Dashboard")
}

func TestDashboardHandler_includesFlashProp(t *testing.T) {
	h := NewDashboardHandler(setupTestRenderer(t), setupTestStore(t), testSite(), cais.Config{}, setupTestInertia(t))

	req := inertiaRequest(http.MethodGet, "/dashboard", nil)
	req = flash.WithMessage(req, flash.Message{Kind: "notice", Message: "Welcome back!"})
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rr.Code)
	}
	flashProp, ok := assertInertiaProp(t, rr, "flash").(map[string]any)
	if !ok || flashProp["notice"] != "Welcome back!" {
		t.Errorf("props.flash missing notice: %v", flashProp)
	}
}

func TestDashboardHandler_pastMonthUsesOverlayNotSQLite(t *testing.T) {
	s := setupTestStore(t)
	uid, err := s.CreateUser("ops@example.com", "hash")
	if err != nil {
		t.Fatal(err)
	}
	t.Setenv(finops.SeedAccountEnv, "111111111111")
	if err := seed.EnsurePrimaryWorkspace(s, uid); err != nil {
		t.Fatal(err)
	}
	tenant, _ := s.FindTenantBySlug(finops.PrimaryTenantSlug)
	accounts, _ := s.ListCloudAccounts(tenant.ID)
	_ = s.ReplaceCostLines(accounts[0].ID, []models.CostLine{{
		Service: "Stored August", MonthlyCents: 999, Source: finops.SourceCE,
		PeriodStart: "2026-08-01", PeriodEnd: "2026-09-01",
	}})
	_ = s.ReplaceFindings(accounts[0].ID, []models.Finding{{
		Kind: finops.FindingCEDenied, Severity: "warning",
	}})

	col := stubMonthCollector{lines: []models.CostLine{{
		Service: "Amazon Lightsail", MonthlyCents: 1947, Source: finops.SourceCE,
		PeriodStart: "2026-07-01", PeriodEnd: "2026-08-01",
	}}}

	h := NewDashboardHandler(setupTestRenderer(t), s, testSite(), cais.Config{}, setupTestInertia(t)).
		WithSyncer(syncer.New(s, col))
	h.now = func() time.Time { return time.Date(2026, 8, 20, 0, 0, 0, 0, time.UTC) }

	req := inertiaRequest(http.MethodGet, "/dashboard?month=2026-07", nil)
	req = session.WithUserID(req, uid)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("status=%d body=%s", rr.Code, rr.Body.String())
	}
	if assertInertiaProp(t, rr, "isCurrent") != false {
		t.Fatal("isCurrent")
	}
	if assertInertiaProp(t, rr, "month") != "2026-07" {
		t.Fatal("month")
	}
	summary := assertInertiaProp(t, rr, "summary").(map[string]any)
	if summary["monthlyUSD"] != "US$ 19,47" {
		t.Fatalf("usd=%v", summary["monthlyUSD"])
	}
	findings, _ := assertInertiaProp(t, rr, "findings").([]any)
	if len(findings) != 0 {
		t.Fatalf("findings=%v", findings)
	}
	stored, _ := s.ListCostLines(accounts[0].ID)
	if stored[0].Service != "Stored August" {
		t.Fatalf("sqlite=%#v", stored)
	}
}
