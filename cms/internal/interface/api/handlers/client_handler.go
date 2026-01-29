package handler

import (
	"errors"
	"fmt"
	"net/http"
	"net/mail"
	"prestaprest/internal/application/schema"
	"prestaprest/internal/application/service"
	"prestaprest/internal/domain/entity"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type ClientHandler struct {
	service *service.ClientService
}

func NewClientHandler(e *echo.Echo, service *service.ClientService) *ClientHandler {
	Handler := &ClientHandler{
		service: service,
	}
	return Handler
}

func (ch *ClientHandler) GetAllClient(c echo.Context) error {
	fmt.Println("Client: GetAllClient")
	client, err := ch.service.GetAllClient()
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, client)
}

func (ch *ClientHandler) GetClient(c echo.Context) error {
	fmt.Println("Client: GetClient")
	clientID, err := strconv.ParseUint(c.Param("client_ID"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	fmt.Println("ClientID: ", clientID)
	client, err := ch.service.GetClient(int64(clientID))
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, client)
}

func (ch *ClientHandler) UpdateClient(c echo.Context) error {
	fmt.Println("Client: UpdateClient")
	clientID, err := strconv.ParseUint(c.Param("client_ID"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	client := new(entity.Client)
	if err := c.Bind(client); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("Failed to parse request body: %s", err),
		})
	}
	oldClient, err := ch.service.GetClient(int64(clientID))
	if client.Password == "" {
		client.Password = oldClient.Password
	}
	client.SetID(clientID)
	err = ch.service.UpdateClient(client)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, "Record updated")
}

func (ch *ClientHandler) DeleteClient(c echo.Context) error {
	fmt.Println("Client: DeleteClient")
	clientID, err := strconv.ParseUint(c.Param("client_ID"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	err = ch.service.DeleteClient(clientID)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, "Record Removed")
}

func (ch *ClientHandler) validateRegisterForm(c echo.Context) error {
	fmt.Println(c)
	first_name := strings.TrimSpace(c.FormValue("first_name"))
	if first_name == "" {
		return errors.New("First name is required")
	}

	last_name := strings.TrimSpace(c.FormValue("last_name"))
	if last_name == "" {
		return errors.New("Last name is required")
	}

	login := strings.TrimSpace(c.FormValue("login"))
	if login == "" {
		return errors.New("Login is required")
	}

	email := strings.TrimSpace(c.FormValue("email"))
	if email == "" {
		return errors.New("email is required")
	}
	if _, err := mail.ParseAddress(email); err != nil {
		return errors.New("invalid email address")
	}

	return nil
}

func (ch *ClientHandler) GetClientSchema(c echo.Context) error {
	schema := schema.GenerateSchema(entity.Client{})
	return c.JSON(http.StatusOK, schema)
}
