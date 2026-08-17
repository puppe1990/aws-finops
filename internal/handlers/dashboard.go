package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/puppe1990/cais/pkg/cais"
	"github.com/puppe1990/cais/pkg/cais/flash"
	"github.com/puppe1990/cais/pkg/cais/meta"
	inertia "github.com/romsar/gonertia/v3"

	"github.com/puppe1990/aws-finops/internal/costest"
	"github.com/puppe1990/aws-finops/internal/finops"
	"github.com/puppe1990/aws-finops/internal/models"
	"github.com/puppe1990/aws-finops/internal/store"
	"github.com/puppe1990/aws-finops/internal/syncer"
)

type DashboardHandler struct {
	renderer *cais.Renderer
	store    store.Store
	site     meta.Site
	cfg      cais.Config
	inertia  *inertia.Inertia
	syncer   *syncer.Syncer
}

func NewDashboardHandler(renderer *cais.Renderer, s store.Store, site meta.Site, cfg cais.Config, i *inertia.Inertia) *DashboardHandler {
	return &DashboardHandler{renderer: renderer, store: s, site: site, cfg: cfg, inertia: i}
}

func (h *DashboardHandler) WithSyncer(s *syncer.Syncer) *DashboardHandler {
	h.syncer = s
	return h
}

func (h *DashboardHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ws, err := loadWorkspace(h.store, r)
	props := inertia.Props{
		"site":          meta.ForRequest(h.site, r),
		"totalContacts": int64(0),
		"env":           h.cfg.Env,
		"summary":       map[string]any{},
		"services":      []any{},
		"resources":     []any{},
		"findings":      []any{},
		"budgets":       []any{},
		"accounts":      []any{},
		"lastSync":      nil,
	}
	if msg, ok := flash.MessageFromRequest(r); ok {
		props["flash"] = inertia.Flash{msg.Kind: msg.Message}
	}
	if err != nil {
		_ = h.inertia.Render(w, r, "Dashboard", props)
		return
	}

	if h.syncer != nil {
		resources, _ := h.store.ListResourcesForTenant(ws.Tenant.ID)
		if len(resources) == 0 {
			ctx, cancel := context.WithTimeout(r.Context(), 25*time.Second)
			_ = h.syncer.SyncTenant(ctx, ws.Tenant.ID)
			cancel()
		}
	}

	for k, v := range shellProps(h.site, r, h.store, ws) {
		props[k] = v
	}
	view, err := buildTenantView(h.store, ws.Tenant.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	props["summary"] = view.Summary
	props["services"] = view.Services
	props["resources"] = view.Resources
	props["findings"] = view.Findings
	props["budgets"] = view.Budgets
	props["accounts"] = view.Accounts
	props["lastSync"] = view.LastSync
	_ = h.inertia.Render(w, r, "Dashboard", props)
}

type tenantView struct {
	Summary   map[string]any
	Services  []map[string]any
	Resources []map[string]any
	Findings  []map[string]any
	Budgets   []map[string]any
	Accounts  []map[string]any
	LastSync  map[string]any
}

func buildTenantView(s store.Store, tenantID int64) (tenantView, error) {
	resources, err := s.ListResourcesForTenant(tenantID)
	if err != nil {
		return tenantView{}, err
	}
	accounts, err := s.ListCloudAccounts(tenantID)
	if err != nil {
		return tenantView{}, err
	}
	findings, err := s.ListFindingsForTenant(tenantID)
	if err != nil {
		return tenantView{}, err
	}
	budgets, err := s.ListBudgets(tenantID)
	if err != nil {
		return tenantView{}, err
	}

	var monthly int64
	var lines []costest.Line
	source := finops.SourceEstimate
	var lastSync map[string]any
	for _, acc := range accounts {
		costLines, err := s.ListCostLines(acc.ID)
		if err != nil {
			return tenantView{}, err
		}
		for _, line := range costLines {
			monthly += line.MonthlyCents
			lines = append(lines, costest.Line{Service: line.Service, MonthlyCents: line.MonthlyCents})
			if line.Source == finops.SourceCE {
				source = finops.SourceCE
			}
		}
		if run, err := s.LastSyncRun(acc.ID); err == nil && run.ID != 0 {
			lastSync = map[string]any{
				"status":  run.Status,
				"source":  run.Source,
				"warning": run.Warning,
				"error":   run.Error,
				"at":      run.StartedAt.Format(time.RFC3339),
			}
		}
	}
	if monthly == 0 {
		for _, r := range resources {
			monthly += r.MonthlyCents
			lines = append(lines, costest.Line{Service: resourceKindLabel(r.Kind), MonthlyCents: r.MonthlyCents})
		}
	}
	now := time.Now()
	mtd := monthly
	if source != finops.SourceCE {
		mtd = costest.MonthToDateCents(monthly, now)
	}

	services := []map[string]any{}
	for _, g := range costest.GroupByService(lines) {
		services = append(services, map[string]any{
			"name": g.Service, "cents": g.MonthlyCents, "usd": formatUSD(g.MonthlyCents),
		})
	}

	return tenantView{
		Summary: map[string]any{
			"monthlyCents":  monthly,
			"monthlyUSD":    formatUSD(monthly),
			"mtdCents":      mtd,
			"mtdUSD":        formatUSD(mtd),
			"source":        source,
			"accountCount":  len(accounts),
			"resourceCount": len(resources),
			"ceDenied":      hasFinding(findings, finops.FindingCEDenied),
		},
		Services:  services,
		Resources: resourceProps(resources),
		Findings:  findingProps(findings),
		Budgets:   budgetProps(budgets, monthly),
		Accounts:  accountProps(accounts),
		LastSync:  lastSync,
	}, nil
}

func resourceProps(resources []models.CloudResource) []map[string]any {
	out := make([]map[string]any, 0, len(resources))
	for _, r := range resources {
		out = append(out, map[string]any{
			"id":     r.ID,
			"kind":   r.Kind,
			"label":  resourceKindLabel(r.Kind),
			"name":   r.Name,
			"region": r.Region,
			"state":  r.State,
			"usd":    formatUSD(r.MonthlyCents),
			"cents":  r.MonthlyCents,
			"source": r.Source,
		})
	}
	return out
}

func findingProps(findings []models.Finding) []map[string]any {
	out := make([]map[string]any, 0, len(findings))
	for _, f := range findings {
		out = append(out, map[string]any{
			"kind": f.Kind, "severity": f.Severity, "title": f.Title, "detail": f.Detail,
		})
	}
	return out
}

func budgetProps(budgets []models.Budget, spent int64) []map[string]any {
	out := make([]map[string]any, 0, len(budgets))
	for _, b := range budgets {
		out = append(out, map[string]any{
			"id":      b.ID,
			"name":    b.Name,
			"amount":  formatUSD(b.AmountCents),
			"spent":   formatUSD(spent),
			"burnBps": costest.BudgetBurnBps(spent, b.AmountCents),
			"period":  b.Period,
			"over":    spent > b.AmountCents && b.AmountCents > 0,
		})
	}
	return out
}

func accountProps(accounts []models.CloudAccount) []map[string]any {
	out := make([]map[string]any, 0, len(accounts))
	for _, a := range accounts {
		out = append(out, map[string]any{
			"id":           a.ID,
			"awsAccountId": a.AWSAccountID,
			"alias":        a.Alias,
			"region":       a.Region,
			"authMode":     a.AuthMode,
			"primary":      a.IsPrimary,
		})
	}
	return out
}

func hasFinding(findings []models.Finding, kind string) bool {
	for _, f := range findings {
		if f.Kind == kind {
			return true
		}
	}
	return false
}
