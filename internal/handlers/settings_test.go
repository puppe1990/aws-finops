package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/puppe1990/cais/pkg/cais/session"

	"github.com/puppe1990/aws-finops/internal/awsinv"
	"github.com/puppe1990/aws-finops/internal/seed"
)

func TestSettingsHandler_includesPolicyAndCloudShell(t *testing.T) {
	s := setupTestStore(t)
	uid, err := s.CreateUser("ops@example.com", "hash")
	if err != nil {
		t.Fatal(err)
	}
	if err := seed.EnsurePrimaryWorkspace(s, uid); err != nil {
		t.Fatal(err)
	}

	h := NewSettingsHandler(s, testSite(), setupTestInertia(t))
	req := inertiaRequest(http.MethodGet, "/settings", nil)
	req = session.WithUserID(req, uid)
	rr := httptest.NewRecorder()
	h.Get(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("status = %d body=%s", rr.Code, rr.Body.String())
	}
	assertInertiaComponent(t, rr, "Settings")
	if got := assertInertiaProp(t, rr, "policy"); got != awsinv.FinOpsIAMPolicy {
		t.Errorf("policy = %v", got)
	}
	if got := assertInertiaProp(t, rr, "cloudShell"); got != awsinv.CloudShellCommand() {
		t.Errorf("cloudShell = %v", got)
	}
}
