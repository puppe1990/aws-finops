package handlers

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/puppe1990/aws-finops/internal/awsinv"
	"github.com/puppe1990/aws-finops/internal/crypto"
	"github.com/puppe1990/aws-finops/internal/finops"
	"github.com/puppe1990/aws-finops/internal/models"
	"github.com/puppe1990/aws-finops/internal/store"
	"github.com/puppe1990/aws-finops/internal/syncer"
	"github.com/puppe1990/cais/pkg/cais"
	"github.com/puppe1990/cais/pkg/cais/flash"
	"github.com/puppe1990/cais/pkg/cais/httpx"
	"github.com/puppe1990/cais/pkg/cais/meta"
	"github.com/puppe1990/cais/pkg/cais/validate"
	inertia "github.com/romsar/gonertia/v3"
)

type AccountsHandler struct {
	store     store.Store
	site      meta.Site
	cfg       cais.Config
	inertia   *inertia.Inertia
	syncer    *syncer.Syncer
	appSecret []byte
}

func NewAccountsHandler(s store.Store, site meta.Site, cfg cais.Config, i *inertia.Inertia, appSecret []byte) *AccountsHandler {
	return &AccountsHandler{store: s, site: site, cfg: cfg, inertia: i, appSecret: appSecret}
}

func (h *AccountsHandler) WithSyncer(s *syncer.Syncer) *AccountsHandler {
	h.syncer = s
	return h
}

func (h *AccountsHandler) List(w http.ResponseWriter, r *http.Request) {
	ws, err := loadWorkspace(h.store, r)
	if err != nil {
		h.inertia.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	accounts, err := h.store.ListCloudAccounts(ws.Tenant.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	props := shellProps(h.site, r, h.store, ws)
	props["accounts"] = accountProps(accounts)
	props["policy"] = awsinv.FinOpsIAMPolicy
	_ = h.inertia.Render(w, r, "Accounts", props)
}

func (h *AccountsHandler) Create(w http.ResponseWriter, r *http.Request) {
	ws, err := loadWorkspace(h.store, r)
	if err != nil {
		h.inertia.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	if err := httpx.ParseFormOrJSON(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	awsID := strings.TrimSpace(r.FormValue("aws_account_id"))
	alias := strings.TrimSpace(r.FormValue("alias"))
	region := strings.TrimSpace(r.FormValue("region"))
	mode := strings.TrimSpace(r.FormValue("auth_mode"))
	if region == "" {
		region = finops.DefaultRegion
	}
	if mode == "" {
		mode = finops.AuthModeDefaultChain
	}
	var errs validate.FieldErrors
	if len(awsID) != 12 {
		errs.Add("aws_account_id", "Informe o account ID com 12 dígitos.")
	}
	if alias == "" {
		errs.Add("alias", "Dê um apelido para a conta.")
	}
	if mode == finops.AuthModeAccessKeys && r.FormValue("access_key_id") == "" {
		errs.Add("access_key_id", "Access key obrigatória neste modo.")
	}
	if errs.Any() {
		ve := make(inertia.ValidationErrors)
		for k, v := range errs {
			ve[k] = v
		}
		ctx := inertia.SetValidationErrors(r.Context(), ve)
		accounts, _ := h.store.ListCloudAccounts(ws.Tenant.ID)
		props := shellProps(h.site, r, h.store, ws)
		props["accounts"] = accountProps(accounts)
		props["policy"] = awsinv.FinOpsIAMPolicy
		_ = h.inertia.Render(w, r.WithContext(ctx), "Accounts", props)
		return
	}
	acc := models.CloudAccount{
		TenantID:     ws.Tenant.ID,
		AWSAccountID: awsID,
		Alias:        alias,
		Region:       region,
		AuthMode:     mode,
	}
	if mode == finops.AuthModeAccessKeys {
		acc.AccessKeyID = strings.TrimSpace(r.FormValue("access_key_id"))
		secret := r.FormValue("secret_access_key")
		if secret != "" && len(h.appSecret) > 0 {
			ct, encErr := crypto.Encrypt(h.appSecret, secret)
			if encErr != nil {
				http.Error(w, encErr.Error(), http.StatusInternalServerError)
				return
			}
			acc.SecretCipher = ct
		}
	}
	if _, err := h.store.EnsureCloudAccount(acc); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	flash.Set(w, "notice", "Conta AWS vinculada ao workspace.", h.cfg.CookieSecure())
	h.inertia.Redirect(w, r, "/accounts", http.StatusSeeOther)
}

func (h *AccountsHandler) Sync(w http.ResponseWriter, r *http.Request) {
	ws, err := loadWorkspace(h.store, r)
	if err != nil {
		h.inertia.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	if h.syncer == nil {
		flash.Set(w, "alert", "Coletor AWS não configurado.", h.cfg.CookieSecure())
		h.inertia.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), 40*time.Second)
	defer cancel()
	if err := h.syncer.SyncTenant(ctx, ws.Tenant.ID); err != nil {
		flash.Set(w, "alert", "Falha ao sincronizar: "+err.Error(), h.cfg.CookieSecure())
	} else {
		flash.Set(w, "notice", "Inventário atualizado a partir da AWS.", h.cfg.CookieSecure())
	}
	h.inertia.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}
