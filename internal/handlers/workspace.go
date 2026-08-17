package handlers

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/puppe1990/aws-finops/internal/finops"
	"github.com/puppe1990/aws-finops/internal/models"
	"github.com/puppe1990/aws-finops/internal/store"
	"github.com/puppe1990/cais/pkg/cais/flash"
	"github.com/puppe1990/cais/pkg/cais/meta"
	"github.com/puppe1990/cais/pkg/cais/session"
	inertia "github.com/romsar/gonertia/v3"
)

type workspace struct {
	User   models.User
	Tenant models.Tenant
	Role   string
}

func loadWorkspace(s store.Store, r *http.Request) (workspace, error) {
	uid, ok := session.UserID(r)
	if !ok {
		return workspace{}, errNoSession
	}
	user, err := s.FindUserByID(uid)
	if err != nil {
		return workspace{}, err
	}
	tenants, err := s.ListTenantsForUser(uid)
	if err != nil {
		return workspace{}, err
	}
	if len(tenants) == 0 {
		return workspace{User: user}, errNoTenant
	}
	tenant := tenants[0]
	if user.ActiveTenantID != 0 {
		for _, t := range tenants {
			if t.ID == user.ActiveTenantID {
				tenant = t
				break
			}
		}
	}
	role, _, err := s.MembershipRole(tenant.ID, uid)
	if err != nil {
		return workspace{}, err
	}
	if user.ActiveTenantID != tenant.ID {
		_ = s.SetActiveTenant(uid, tenant.ID)
	}
	return workspace{User: user, Tenant: tenant, Role: role}, nil
}

func createPersonalTenant(s store.Store, userID int64, email string) error {
	slug := slugFromEmail(email, userID)
	id, err := s.CreateTenant("Workspace", slug)
	if err != nil {
		return err
	}
	if err := s.AddMember(id, userID, finops.RoleOwner); err != nil {
		return err
	}
	return s.SetActiveTenant(userID, id)
}

func slugFromEmail(email string, userID int64) string {
	local := strings.ToLower(strings.Split(email, "@")[0])
	re := regexp.MustCompile(`[^a-z0-9]+`)
	local = strings.Trim(re.ReplaceAllString(local, "-"), "-")
	if local == "" {
		local = "ws"
	}
	return fmt.Sprintf("%s-%d", local, userID)
}

func shellProps(hSite meta.Site, r *http.Request, s store.Store, ws workspace) inertia.Props {
	tenants, _ := s.ListTenantsForUser(ws.User.ID)
	props := inertia.Props{
		"site":       meta.ForRequest(hSite, r),
		"userEmail":  ws.User.Email,
		"tenant":     map[string]any{"id": ws.Tenant.ID, "name": ws.Tenant.Name, "slug": ws.Tenant.Slug, "role": ws.Role},
		"tenants":    tenantsProps(tenants),
		"primaryAws": finops.SeedAWSAccountID(),
	}
	if msg, ok := flash.MessageFromRequest(r); ok {
		props["flash"] = inertia.Flash{msg.Kind: msg.Message}
	}
	return props
}

func tenantsProps(tenants []models.Tenant) []map[string]any {
	out := make([]map[string]any, 0, len(tenants))
	for _, t := range tenants {
		out = append(out, map[string]any{"id": t.ID, "name": t.Name, "slug": t.Slug})
	}
	return out
}

var (
	errNoSession = fmt.Errorf("not signed in")
	errNoTenant  = fmt.Errorf("no tenant")
)
