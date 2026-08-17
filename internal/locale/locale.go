package locale

import (
	"net/http"
	"net/url"
	"strings"
)

const CookieName = "cifra_locale"

func Normalize(raw string) string {
	tag := strings.ToLower(strings.TrimSpace(raw))
	tag = strings.ReplaceAll(tag, "_", "-")
	switch {
	case tag == "" || strings.HasPrefix(tag, "en"):
		return "en"
	case strings.HasPrefix(tag, "pt"):
		return "pt"
	default:
		return "en"
	}
}

func FromRequest(r *http.Request, fallback string) string {
	if c, err := r.Cookie(CookieName); err == nil && c.Value != "" {
		return Normalize(c.Value)
	}
	if al := r.Header.Get("Accept-Language"); al != "" {
		first := strings.Split(al, ",")[0]
		first = strings.Split(first, ";")[0]
		if n := Normalize(first); n != "" {
			return n
		}
	}
	return Normalize(fallback)
}

func SetCookie(w http.ResponseWriter, locale string, secure bool) {
	http.SetCookie(w, &http.Cookie{
		Name:     CookieName,
		Value:    Normalize(locale),
		Path:     "/",
		MaxAge:   365 * 24 * 60 * 60,
		HttpOnly: false,
		SameSite: http.SameSiteLaxMode,
		Secure:   secure,
	})
}

func SafeBack(r *http.Request) string {
	ref := r.Header.Get("Referer")
	if ref == "" {
		return "/"
	}
	u, err := url.Parse(ref)
	if err != nil || u.Path == "" {
		return "/"
	}
	if u.Host != "" && u.Host != r.Host {
		return "/"
	}
	if !strings.HasPrefix(u.Path, "/") {
		return "/"
	}
	return u.Path
}
