package handlers

import (
	"net/http"

	"github.com/puppe1990/cais/pkg/cais/meta"
	inertia "github.com/romsar/gonertia/v3"

	"github.com/puppe1990/aws-finops/internal/awsinv"
	"github.com/puppe1990/aws-finops/internal/finops"
	"github.com/puppe1990/aws-finops/internal/store"
)

type SettingsHandler struct {
	store   store.Store
	site    meta.Site
	inertia *inertia.Inertia
}

func NewSettingsHandler(s store.Store, site meta.Site, i *inertia.Inertia) *SettingsHandler {
	return &SettingsHandler{store: s, site: site, inertia: i}
}

func (h *SettingsHandler) Get(w http.ResponseWriter, r *http.Request) {
	ws, err := loadWorkspace(h.store, r)
	if err != nil {
		h.inertia.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	props := shellProps(h.site, r, h.store, ws)
	props["policy"] = awsinv.FinOpsIAMPolicy
	props["seededAccount"] = finops.SeedAWSAccountID()
	_ = h.inertia.Render(w, r, "Settings", props)
}
