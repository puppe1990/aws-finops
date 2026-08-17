package locale

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNormalize_mapsPtBRAndEn(t *testing.T) {
	if got := Normalize("pt-BR"); got != "pt" {
		t.Fatalf("pt-BR = %q", got)
	}
	if got := Normalize("en-US"); got != "en" {
		t.Fatalf("en-US = %q", got)
	}
	if got := Normalize("fr"); got != "en" {
		t.Fatalf("unknown = %q, want en", got)
	}
}

func TestFromRequest_prefersCookie(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Accept-Language", "en")
	req.AddCookie(&http.Cookie{Name: CookieName, Value: "pt"})
	if got := FromRequest(req, "en"); got != "pt" {
		t.Fatalf("got %q, want pt", got)
	}
}

func TestFromRequest_usesAcceptLanguage(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Accept-Language", "pt-BR,pt;q=0.9")
	if got := FromRequest(req, "en"); got != "pt" {
		t.Fatalf("got %q, want pt", got)
	}
}

func TestSetCookie_writesNormalizedLocale(t *testing.T) {
	rr := httptest.NewRecorder()
	SetCookie(rr, "pt-BR", false)
	res := rr.Result()
	defer func() { _ = res.Body.Close() }()
	found := false
	for _, c := range res.Cookies() {
		if c.Name == CookieName && c.Value == "pt" {
			found = true
		}
	}
	if !found {
		t.Fatalf("cookie not set: %#v", res.Cookies())
	}
}

func TestSafeBack_blocksExternalHost(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/locale", nil)
	req.Host = "cifra.test"
	req.Header.Set("Referer", "https://evil.example/phish")
	if got := SafeBack(req); got != "/" {
		t.Fatalf("got %q", got)
	}
}

func TestSafeBack_keepsLocalPath(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/locale", nil)
	req.Host = "cifra.test"
	req.Header.Set("Referer", "https://cifra.test/dashboard")
	if got := SafeBack(req); got != "/dashboard" {
		t.Fatalf("got %q", got)
	}
}
