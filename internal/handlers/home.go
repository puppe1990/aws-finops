package handlers

import (
	"fmt"
	"net/http"

	"github.com/puppe1990/cais/pkg/cais"
	"github.com/puppe1990/cais/pkg/cais/i18n"
	"github.com/puppe1990/cais/pkg/cais/meta"
	inertia "github.com/romsar/gonertia/v3"

	"github.com/puppe1990/aws-finops/internal/finops"
)

type HomeHandler struct {
	renderer *cais.Renderer
	site     meta.Site
	catalog  *i18n.Catalog
	cfg      cais.Config
	inertia  *inertia.Inertia
}

func NewHomeHandler(renderer *cais.Renderer, site meta.Site, catalog *i18n.Catalog, cfg cais.Config, i *inertia.Inertia) *HomeHandler {
	return &HomeHandler{renderer: renderer, site: site, catalog: catalog, cfg: cfg, inertia: i}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	site := meta.ForRequest(h.site, r)
	cat := requestCatalog(r, h.cfg.Locale)
	props := publicProps(h.site, r, h.cfg.Locale)
	props["title"] = cat.T("home.title")
	labels, _ := props["labels"].(map[string]string)
	if labels == nil {
		labels = map[string]string{}
	}
	labels["heading"] = cat.T("home.rails_heading")
	labels["subtitle"] = fmt.Sprintf(cat.T("home.rails_subtitle"), site.AppName)
	labels["stack"] = cat.T("home.stack")
	labels["contact"] = cat.T("home.contact_link")
	labels["login"] = cat.T("auth.login_submit")
	labels["dashboard"] = cat.T("dashboard.title")
	labels["account"] = finops.SeedAWSAccountID()
	if labels["account"] != "" {
		labels["eyebrow"] = fmt.Sprintf(cat.T("home.eyebrow_seeded"), labels["account"])
	} else {
		labels["eyebrow"] = cat.T("home.eyebrow_empty")
	}
	props["labels"] = labels
	if err := h.inertia.Render(w, r, "Home", props); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
