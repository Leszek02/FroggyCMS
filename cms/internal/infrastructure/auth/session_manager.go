package auth

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	Store *sessions.CookieStore
)

const (
	SessionName     = "Prestapress_session"
	SessionUserID   = "user_id"
	SessionUserType = "user_type"
	SessionAnonID   = "anon_session_id"
)

const (
	UserTypeClient = "client"
	UserTypeAdmin  = "admin"
	UserTypeAnon   = "anonymous"
)

func InitSessionStore(authKey, encryptionKey string) {

	Store = sessions.NewCookieStore(
		[]byte(authKey),
		[]byte(encryptionKey),
	)

	Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 60 * 60 * 7,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}
}
