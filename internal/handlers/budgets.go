package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/puppe1990/cais/pkg/cais"
	"github.com/puppe1990/cais/pkg/cais/flash"
	"github.com/puppe1990/cais/pkg/cais/httpx"
	"github.com/puppe1990/cais/pkg/cais/meta"
	inertia "github.com/romsar/gonertia/v3"

	"github.com/puppe1990/aws-finops/internal/models"
	"github.com/puppe1990/aws-finops/internal/store"
)

type BudgetsHandler struct {
	store   store.Store
	site    meta.Site
	cfg     cais.Config
	inertia *inertia.Inertia
}

func NewBudgetsHandler(s store.Store, site meta.Site, cfg cais.Config, i *inertia.Inertia) *BudgetsHandler {
	return &BudgetsHandler{store: s, site: site, cfg: cfg, inertia: i}
}

func (h *BudgetsHandler) List(w http.ResponseWriter, r *http.Request) {
	ws, err := loadWorkspace(h.store, r)
	if err != nil {
		h.inertia.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	view, err := buildTenantView(h.store, ws.Tenant.ID, requestCatalog(r, h.cfg.Locale))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	props := shellProps(h.site, r, h.store, ws)
	props["budgets"] = view.Budgets
	props["spentUSD"] = view.Summary["monthlyUSD"]
	_ = h.inertia.Render(w, r, "Budgets", props)
}

func (h *BudgetsHandler) Create(w http.ResponseWriter, r *http.Request) {
	ws, err := loadWorkspace(h.store, r)
	if err != nil {
		h.inertia.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	if err := httpx.ParseFormOrJSON(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	name := strings.TrimSpace(r.FormValue("name"))
	amount, _ := strconv.ParseFloat(strings.ReplaceAll(r.FormValue("amount_usd"), ",", "."), 64)
	if name == "" || amount <= 0 {
		flash.Set(w, "alert", requestCatalog(r, h.cfg.Locale).T("bud.need_fields"), h.cfg.CookieSecure())
		h.inertia.Redirect(w, r, "/budgets", http.StatusSeeOther)
		return
	}
	if _, err := h.store.CreateBudget(models.Budget{
		TenantID:    ws.Tenant.ID,
		Name:        name,
		AmountCents: int64(amount * 100),
		Period:      "monthly",
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	flash.Set(w, "notice", requestCatalog(r, h.cfg.Locale).T("bud.created"), h.cfg.CookieSecure())
	h.inertia.Redirect(w, r, "/budgets", http.StatusSeeOther)
}
