package security

import "net/http"

type cookieContext interface {
	SetCookie(name string, value string, maxAge int, path string, domain string, secure bool, httpOnly bool)
}

func missingCookieFlags(w http.ResponseWriter, sessionID string) {
	http.SetCookie(w, &http.Cookie{
		Name:  "session_id",
		Value: sessionID,
		Path:  "/",
	})
}

func sameSiteNoneWithoutSecure(w http.ResponseWriter, token string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	})
}

func rawSetCookieMissingFlags(w http.ResponseWriter, jwt string) {
	w.Header().Set("Set-Cookie", "jwt="+jwt+"; Path=/; SameSite=Lax")
}

func frameworkCookieExplicitlyUnsafe(c cookieContext, authToken string) {
	c.SetCookie("auth_token", authToken, 3600, "/", "", false, true)
}
