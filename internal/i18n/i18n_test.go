package i18n

import "testing"

func TestDefaultCatalog_english(t *testing.T) {
	c := DefaultCatalog()
	if got := c.T("auth.welcome"); got != "Welcome!" {
		t.Errorf("T(auth.welcome) = %q", got)
	}
	if got := c.T("auth.signup_prompt"); got != "Don't have an account?" {
		t.Errorf("T(auth.signup_prompt) = %q", got)
	}
}

func TestLabels_dashboardBudgetLine(t *testing.T) {
	en := Labels("en")
	if en["dash.budget_line"] == "" {
		t.Fatal("missing dash.budget_line")
	}
	if _, ok := en["dash.costliest"]; ok {
		t.Fatal("dash.costliest should be removed")
	}
	if _, ok := en["dash.col_month"]; ok {
		t.Fatal("dash.col_month should be removed")
	}
}

func TestLabels_ledgerMonthKeys(t *testing.T) {
	en := Labels("en")
	pt := Labels("pt-BR")
	if en["dash.month_fmt"] != "%s %d" || pt["dash.month_fmt"] != "%s %d" {
		t.Fatalf("month_fmt en=%q pt=%q", en["dash.month_fmt"], pt["dash.month_fmt"])
	}
	if en["dash.m08"] != "Aug" || pt["dash.m08"] != "ago" {
		t.Fatalf("m08 en=%q pt=%q", en["dash.m08"], pt["dash.m08"])
	}
}

func TestLabels_sameKeysEnAndPt(t *testing.T) {
	en := Labels("en")
	pt := Labels("pt-BR")
	if len(en) == 0 || len(pt) == 0 {
		t.Fatal("empty catalog")
	}
	for k := range en {
		if _, ok := pt[k]; !ok {
			t.Errorf("pt missing key %q", k)
		}
	}
	for k := range pt {
		if _, ok := en[k]; !ok {
			t.Errorf("en missing key %q", k)
		}
	}
}

func TestNewCatalog_portuguese(t *testing.T) {
	c := NewCatalog("pt-BR")
	if got := c.T("auth.welcome"); got != "Bem-vindo!" {
		t.Errorf("T(auth.welcome) = %q", got)
	}
	if c.HTMLLang() != "pt-BR" {
		t.Errorf("HTMLLang() = %q, want pt-BR", c.HTMLLang())
	}
}
