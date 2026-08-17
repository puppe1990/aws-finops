package handlers

import (
	"net/http"

	"github.com/puppe1990/aws-finops/internal/store"
	"github.com/puppe1990/cais/pkg/cais/meta"
	inertia "github.com/romsar/gonertia/v3"
)

type ResourcesHandler struct {
	store   store.Store
	site    meta.Site
	inertia *inertia.Inertia
}

func NewResourcesHandler(s store.Store, site meta.Site, i *inertia.Inertia) *ResourcesHandler {
	return &ResourcesHandler{store: s, site: site, inertia: i}
}

func (h *ResourcesHandler) List(w http.ResponseWriter, r *http.Request) {
	ws, err := loadWorkspace(h.store, r)
	if err != nil {
		h.inertia.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	resources, err := h.store.ListResourcesForTenant(ws.Tenant.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	props := shellProps(h.site, r, h.store, ws)
	props["resources"] = resourceProps(resources)
	_ = h.inertia.Render(w, r, "Resources", props)
}
