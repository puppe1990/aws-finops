package handlers

import (
	"net/http"

	"github.com/puppe1990/cais/pkg/cais"
	"github.com/puppe1990/cais/pkg/cais/httpx"
	inertia "github.com/romsar/gonertia/v3"

	"github.com/puppe1990/aws-finops/internal/locale"
)

type LocaleHandler struct {
	cfg     cais.Config
	inertia *inertia.Inertia
}

func NewLocaleHandler(cfg cais.Config, i *inertia.Inertia) *LocaleHandler {
	return &LocaleHandler{cfg: cfg, inertia: i}
}

func (h *LocaleHandler) Post(w http.ResponseWriter, r *http.Request) {
	if err := httpx.ParseFormOrJSON(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	locale.SetCookie(w, r.FormValue("locale"), h.cfg.CookieSecure())
	h.inertia.Redirect(w, r, locale.SafeBack(r), http.StatusSeeOther)
}
