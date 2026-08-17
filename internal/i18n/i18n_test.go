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
