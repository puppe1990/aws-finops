package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/puppe1990/cais/pkg/cais"
	"github.com/puppe1990/cais/pkg/cais/flash"
	"github.com/puppe1990/cais/pkg/cais/meta"
	inertia "github.com/romsar/gonertia/v3"

	"github.com/puppe1990/aws-finops/internal/awsinv"
	"github.com/puppe1990/aws-finops/internal/models"
	"github.com/puppe1990/aws-finops/internal/store"
	"github.com/puppe1990/aws-finops/internal/syncer"
)

type CompareHandler struct {
	store   store.Store
	site    meta.Site
	cfg     cais.Config
	inertia *inertia.Inertia
	syncer  *syncer.Syncer
	now     func() time.Time
}

func NewCompareHandler(s store.Store, site meta.Site, cfg cais.Config, i *inertia.Inertia) *CompareHandler {
	return &CompareHandler{store: s, site: site, cfg: cfg, inertia: i}
}

func (h *CompareHandler) WithSyncer(s *syncer.Syncer) *CompareHandler {
	h.syncer = s
	return h
}

func (h *CompareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ws, err := loadWorkspace(h.store, r)
	props := inertia.Props{
		"site":     meta.ForRequest(h.site, r),
		"months":   []any{},
		"services": []any{},
		"ceDenied": false,
		"flash":    inertia.Flash{},
	}
	if msg, ok := flash.MessageFromRequest(r); ok {
		props["flash"] = inertia.Flash{msg.Kind: msg.Message}
	}
	if err != nil {
		_ = h.inertia.Render(w, r, "Compare", props)
		return
	}

	now := time.Now().UTC()
	if h.now != nil {
		now = h.now()
	}
	window := awsinv.LookbackMonths(now)
	from, to := window[0], window[len(window)-1]
	cat := requestCatalog(r, h.cfg.Locale)

	var lines []models.CostLine
	ceDenied := false
	if h.syncer != nil {
		ctx, cancel := context.WithTimeout(r.Context(), 25*time.Second)
		var overlayErr error
		lines, overlayErr = h.syncer.CostByMonth(ctx, ws.Tenant.ID, from, to)
		cancel()
		if overlayErr != nil {
			if awsinv.IsAccessDenied(overlayErr) {
				ceDenied = true
			}
			lines = nil
		}
	}

	buckets := awsinv.FoldMonthlyLines(from, to, lines)

	for k, v := range shellProps(h.site, r, h.store, ws) {
		props[k] = v
	}
	props["months"] = compareMonthRows(buckets, cat, now)
	props["services"] = compareServiceHistory(buckets)
	props["ceDenied"] = ceDenied
	_ = h.inertia.Render(w, r, "Compare", props)
}
