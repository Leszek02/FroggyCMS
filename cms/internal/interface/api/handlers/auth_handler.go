package handler

import (
	"errors"
	"fmt"
	"net/http"
	"net/mail"
	"prestaprest/internal/application/service"
	"prestaprest/internal/domain/entity"
	"prestaprest/internal/infrastructure/auth"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	clientService        *service.ClientService
	administratorService *service.AdministratorService
	sessionService       *auth.SessionService
	passwordManager      *auth.PasswordHasher
}

func NewAuthHandler(clientService *service.ClientService,
	administratorService *service.AdministratorService,
	sessionService *auth.SessionService) *AuthHandler {
	return &AuthHandler{
		clientService:        clientService,
		administratorService: administratorService,
		sessionService:       sessionService,
		passwordManager:      auth.NewPasswordHasher(),
	}
}

func (h *AuthHandler) Register(c echo.Context) error {
	fmt.Println("Auth: Register")
	if err := h.validateForm(c, true); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("Validation error: %s", err),
		})
	}
	client := new(entity.Client)
	if err := c.Bind(client); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("Failed to parse request body: %s", err),
		})
	}

	hashedPassword, err := h.passwordManager.HashPassword(client.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to process password",
		})
	}
	client.Password = hashedPassword

	fmt.Println("Creating client: ", client)
	err = h.clientService.CreateClient(client)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	err = h.sessionService.CreateClientSession(c, int64(client.Client_ID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create session",
		})
	}

	return c.Redirect(302, "/shop/main")
}

func (h *AuthHandler) Login(c echo.Context) error {
	var loginData struct {
		Login    string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
	}
	fmt.Println("Auth: Login")
	if err := h.validateForm(c, false); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("Validation error: %s", err),
		})
	}
	if err := c.Bind(&loginData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request",
		})
	}

	hashedPassword, err := h.passwordManager.HashPassword(loginData.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to process password",
		})
	}
	loginData.Password = hashedPassword

	client, err := h.clientService.Authenticate(loginData.Login, loginData.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Invalid credentials",
		})
	}
	fmt.Println("Authenticated client: ", client.Client_ID)
	err = h.sessionService.CreateClientSession(c, int64(client.Client_ID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create session",
		})
	}
	client.Last_login = time.Now()
	h.clientService.UpdateClient(client)

	return c.Redirect(http.StatusFound, "/shop/main")
}

func (h *AuthHandler) AdminLogin(c echo.Context) error {
	var loginData struct {
		Login    string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
	}
	fmt.Println("Auth: Admin Login")
	if err := h.validateForm(c, false); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("Validation error: %s", err),
		})
	}
	if err := c.Bind(&loginData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request",
		})
	}

	hashedPassword, err := h.passwordManager.HashPassword(loginData.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to process password",
		})
	}
	loginData.Password = hashedPassword

	admin, err := h.administratorService.Authenticate(loginData.Login, loginData.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Invalid credentials",
		})
	}
	fmt.Println("Authenticated admin: ", admin.Administrator_ID)
	err = h.sessionService.CreateAdminSession(c, int32(admin.Administrator_ID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create session",
		})
	}

	return c.Redirect(http.StatusFound, "/admin-static/test_admin_panel")
}

func (h *AuthHandler) Logout(c echo.Context) error {
	err := h.sessionService.DestroySession(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to logout",
		})
	}

	return c.Redirect(http.StatusFound, "/shop/main")
}

func (h *AuthHandler) ChangePassword(c echo.Context) error {
	userID, ok := h.sessionService.GetUserID(c)
	if ok == false {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to authenticate user",
		})
	}
	password := c.FormValue("password")
	clientID := userID.(int64)
	client, err := h.clientService.GetClient(clientID)

	hashedPassword, err := h.passwordManager.HashPassword(password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to process password",
		})
	}

	client.Password = hashedPassword
	fmt.Println("NEW PASSWORD: ", client.Password)
	err = h.clientService.UpdateClient(client)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to update the password",
		})
	}
	return c.JSON(http.StatusOK, "Password updated successfully")
}

func (h *AuthHandler) validateForm(c echo.Context, registerForm bool) error {
	email := strings.TrimSpace(c.FormValue("email"))
	if email == "" {
		return errors.New("email is required")
	}
	if _, err := mail.ParseAddress(email); err != nil {
		return errors.New("invalid email address")
	}

	pwd := c.FormValue("password")
	if len(pwd) < 6 {
		return errors.New("password is required and must be at least 6 characters")
	}
	if registerForm {
		first_name := strings.TrimSpace(c.FormValue("First_name"))
		fmt.Println(first_name)
		if first_name == "" {
			return errors.New("First name is required")
		}

		last_name := strings.TrimSpace(c.FormValue("Last_name"))
		if last_name == "" {
			return errors.New("Last name is required")
		}

		login := strings.TrimSpace(c.FormValue("login"))
		if login == "" {
			return errors.New("Login is required")
		}

		if pwd != c.FormValue("password_confirm") {
			return errors.New("passwords do not match")
		}
	}

	return nil
}
