package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/puppe1990/cais/pkg/cais"
	"github.com/puppe1990/cais/pkg/cais/session"

	"github.com/puppe1990/aws-finops/internal/finops"
	"github.com/puppe1990/aws-finops/internal/seed"
)

func TestDashboardHandler_showsSeededAccount(t *testing.T) {
	s := setupTestStore(t)
	uid, err := s.CreateUser("ops@example.com", "hash")
	if err != nil {
		t.Fatal(err)
	}
	t.Setenv(finops.SeedAccountEnv, "111111111111")
	if err := seed.EnsurePrimaryWorkspace(s, uid); err != nil {
		t.Fatal(err)
	}

	h := NewDashboardHandler(setupTestRenderer(t), s, testSite(), cais.Config{}, setupTestInertia(t))
	req := inertiaRequest(http.MethodGet, "/dashboard", nil)
	req = session.WithUserID(req, uid)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("status = %d body=%s", rr.Code, rr.Body.String())
	}
	assertInertiaComponent(t, rr, "Dashboard")
	accounts, ok := assertInertiaProp(t, rr, "accounts").([]any)
	if !ok || len(accounts) != 1 {
		t.Fatalf("accounts = %#v", accounts)
	}
	first, _ := accounts[0].(map[string]any)
	if first["awsAccountId"] != "111111111111" {
		t.Fatalf("seeded account = %#v", first)
	}
}
