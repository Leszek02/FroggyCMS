package handler

import (
	"fmt"
	"net/http"
	"prestaprest/internal/application/schema"
	"prestaprest/internal/application/service"
	"prestaprest/internal/domain/entity"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AdministratorHandler struct {
	service *service.AdministratorService
}

func NewAdministratorHandler(e *echo.Echo, service *service.AdministratorService) *AdministratorHandler {
	Handler := &AdministratorHandler{
		service: service,
	}
	return Handler
}

func (ah *AdministratorHandler) GetAllAdministrators(c echo.Context) error {
	fmt.Println("Administrator: GetAllAdministrators")
	administrator, err := ah.service.GetAllAdministrators()
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, administrator)
}

func (ah *AdministratorHandler) GetAdministrator(c echo.Context) error {
	fmt.Println("Administrator: GetAdministrator")
	administratorIDuint64, err := strconv.ParseUint(c.Param("administrator_ID"), 10, 32)
	administratorID := uint(administratorIDuint64)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	administrator, err := ah.service.GetAdministrator(administratorID)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, administrator)
}

func (ah *AdministratorHandler) CreateAdministrator(c echo.Context) error {
	fmt.Println("Administrator: CreateAdministrator")
	administrator := new(entity.Administrator)
	if err := c.Bind(administrator); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("Failed to parse request body: %s", err),
		})
	}
	err := ah.service.CreateAdministrator(administrator)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, "Record created")
}

func (ah *AdministratorHandler) UpdateAdministrator(c echo.Context) error {
	fmt.Println("Administrator: UpdateAdministrator")
	administratorIDuint64, err := strconv.ParseUint(c.Param("administrator_ID"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	administratorID := uint(administratorIDuint64)
	administrator := new(entity.Administrator)
	if err := c.Bind(administrator); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("Failed to parse request body: %s", err),
		})
	}
	administrator.SetID(administratorID)
	err = ah.service.UpdateAdministrator(administrator)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, "Record updated")
}

func (ah *AdministratorHandler) DeleteAdministrator(c echo.Context) error {
	fmt.Println("Administrator: DeleteAdministrator")
	administratorIDuint64, err := strconv.ParseUint(c.Param("administrator_ID"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	administratorID := uint(administratorIDuint64)
	err = ah.service.DeleteAdministrator(administratorID)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, "Record Removed")
}

func (ah *AdministratorHandler) GetAdministratorSchema(c echo.Context) error {
	schema := schema.GenerateSchema(entity.Administrator{})
	return c.JSON(http.StatusOK, schema)
}
