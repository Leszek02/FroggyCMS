package auth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

type SessionService struct {
	store *sessions.CookieStore
}

func NewSessionService(store *sessions.CookieStore) *SessionService {
	return &SessionService{
		store: store,
	}
}

// TODO: this won't be needed after all (I think)
func (s *SessionService) GenerateSessionID() (string, error) {
	sessionID := make([]byte, 32)

	_, err := rand.Read(sessionID)
	if err != nil {
		return "", fmt.Errorf("failed to generate session ID: %w", err)
	}

	return base64.URLEncoding.EncodeToString(sessionID), nil
}

func (s *SessionService) CreateClientSession(c echo.Context, clientID int64) error {
	session, err := s.store.Get(c.Request(), SessionName)
	if err != nil {
		return fmt.Errorf("failed to get session: %w", err)
	}
	session.Values[SessionUserID] = clientID
	session.Values[SessionUserType] = UserTypeClient

	delete(session.Values, SessionAnonID)

	return session.Save(c.Request(), c.Response())
}

func (s *SessionService) CreateAdminSession(c echo.Context, adminID int32) error {
	session, err := s.store.Get(c.Request(), SessionName)
	if err != nil {
		return fmt.Errorf("failed to get session: %w", err)
	}

	session.Values[SessionUserID] = adminID
	session.Values[SessionUserType] = UserTypeAdmin

	return session.Save(c.Request(), c.Response())
}

func (s *SessionService) CreateAnonymousSession(c echo.Context) (string, error) {
	session, err := s.store.Get(c.Request(), SessionName)
	if err != nil {
		return "", fmt.Errorf("failed to get session: %w", err)
	}

	if anonID, ok := session.Values[SessionAnonID].(string); ok && anonID != "" {
		return anonID, nil
	}

	anonID, err := s.GenerateSessionID()
	if err != nil {
		return "", err
	}

	session.Values[SessionAnonID] = anonID
	session.Values[SessionUserType] = UserTypeAnon

	err = session.Save(c.Request(), c.Response())
	if err != nil {
		return "", fmt.Errorf("failed to save session: %w", err)
	}

	return anonID, nil
}

func (s *SessionService) GetSession(c echo.Context) (*sessions.Session, error) {
	return s.store.Get(c.Request(), SessionName)
}

func (s *SessionService) GetUserID(c echo.Context) (interface{}, bool) {
	session, err := s.GetSession(c)
	if err != nil {
		return nil, false
	}

	userID, ok := session.Values[SessionUserID]
	return userID, ok
}

func (s *SessionService) GetCartItems(c echo.Context) (map[uint64]int, error) {
	session, err := s.GetSession(c)
	if err != nil {
		return nil, err
	}

	cartItems, ok := session.Values["Cart"].(map[uint64]int)
	if !ok {
		return map[uint64]int{}, nil
	}

	return cartItems, nil
}

func (s *SessionService) UpdateCart(c echo.Context, cartItems map[uint64]int) error {
	session, err := s.GetSession(c)
	if err != nil {
		return err
	}

	session.Values["Cart"] = cartItems
	return session.Save(c.Request(), c.Response())
}

func (s *SessionService) AddToCart(c echo.Context, productID uint64, quantity int) error {
	cartItems, err := s.GetCartItems(c)
	if err != nil {
		return err
	}

	if existingQty, exists := cartItems[productID]; exists {
		cartItems[productID] = existingQty + quantity
	} else {
		cartItems[productID] = quantity
	}

	return s.UpdateCart(c, cartItems)
}

func (s *SessionService) RemoveFromCart(c echo.Context, productID uint64) error {
	cartItems, err := s.GetCartItems(c)
	if err != nil {
		return err
	}

	delete(cartItems, productID)

	return s.UpdateCart(c, cartItems)
}

func (s *SessionService) UpdateCartItemQuantity(c echo.Context, productID uint64, quantity int) error {
	cartItems, err := s.GetCartItems(c)
	if err != nil {
		return err
	}

	if quantity <= 0 {
		delete(cartItems, productID)
	} else {
		cartItems[productID] = quantity
	}

	return s.UpdateCart(c, cartItems)
}

func (s *SessionService) GetCartQuantity(c echo.Context) int {
	cartItems, err := s.GetCartItems(c)
	if err != nil {
		return 0
	}

	totalQuantity := 0
	for _, qty := range cartItems {
		totalQuantity += qty
	}

	return totalQuantity
}

func (s *SessionService) ClearCart(c echo.Context) error {
	return s.UpdateCart(c, map[uint64]int{})
}

func (s *SessionService) GetUserType(c echo.Context) string {
	session, err := s.GetSession(c)
	if err != nil {
		return UserTypeAnon
	}

	userType, ok := session.Values[SessionUserType].(string)
	if !ok {
		return UserTypeAnon
	}

	return userType
}

func (s *SessionService) GetAnonymousID(c echo.Context) (string, bool) {
	session, err := s.GetSession(c)
	if err != nil {
		return "", false
	}

	anonID, ok := session.Values[SessionAnonID].(string)
	return anonID, ok
}

func (s *SessionService) IsAuthenticated(c echo.Context) bool {
	userType := s.GetUserType(c)
	return userType == UserTypeClient || userType == UserTypeAdmin
}

func (s *SessionService) IsAdmin(c echo.Context) bool {
	return s.GetUserType(c) == UserTypeAdmin
}

func (s *SessionService) DestroySession(c echo.Context) error {
	session, err := s.GetSession(c)
	if err != nil {
		return err
	}

	session.Options.MaxAge = -1
	for k := range session.Values {
		delete(session.Values, k)
	}

	return session.Save(c.Request(), c.Response())
}
