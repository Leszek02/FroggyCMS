package session

import (
	"prestaprest/internal/infrastructure/auth"

	"github.com/labstack/echo/v4"
)

type SessionMiddleware struct {
	sessionService *auth.SessionService
}

func NewSessionMiddleware(sessionService *auth.SessionService) *SessionMiddleware {
	return &SessionMiddleware{
		sessionService: sessionService,
	}
}

func (m *SessionMiddleware) LoadSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, authenticated := m.sessionService.GetUserID(c)
		userType := m.sessionService.GetUserType(c)

		c.Set("user_id", userID)
		c.Set("user_type", userType)
		c.Set("authenticated", authenticated)

		return next(c)
	}
}

func (m *SessionMiddleware) RequireAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if !m.sessionService.IsAuthenticated(c) {
			return c.Redirect(302, "/login")
		}

		if m.sessionService.GetUserType(c) != auth.UserTypeClient {
			return c.Redirect(302, "/login")
		}

		return next(c)
	}
}

func (m *SessionMiddleware) RequireAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if !m.sessionService.IsAdmin(c) {
			return c.Redirect(302, "/admin/login")
		}

		return next(c)
	}
}

func (m *SessionMiddleware) IsLogged(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if m.sessionService.IsAuthenticated(c) {
			return c.Redirect(302, "/shop/main")
		}
		return next(c)
	}
}

func (m *SessionMiddleware) IsAdminLogged(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if m.sessionService.IsAdmin(c) {
			return c.Redirect(302, "/admin-static/test_admin_panel")
		}
		return next(c)
	}
}
