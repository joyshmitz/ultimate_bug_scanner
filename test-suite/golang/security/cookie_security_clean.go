package security

import "net/http"

type cleanCookieContext interface {
	SetCookie(name string, value string, maxAge int, path string, domain string, secure bool, httpOnly bool)
}

func secureSessionCookie(w http.ResponseWriter, sessionID string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})
}

func secureCrossSiteCookie(w http.ResponseWriter, token string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	})
}

func rawSetCookieWithFlags(w http.ResponseWriter, jwt string) {
	w.Header().Set("Set-Cookie", "jwt="+jwt+"; Path=/; HttpOnly; Secure; SameSite=Lax")
}

func frameworkCookieWithRequiredBooleans(c cleanCookieContext, authToken string) {
	c.SetCookie("auth_token", authToken, 3600, "/", "", true, true)
}

func publicPreferenceCookie(w http.ResponseWriter, theme string) {
	http.SetCookie(w, &http.Cookie{
		Name:  "theme",
		Value: theme,
		Path:  "/",
	})
}
