package handler

import (
	"fmt"
	"net/http"
	"prestaprest/internal/application/schema"
	"prestaprest/internal/application/service"
	"prestaprest/internal/domain/entity"
	"prestaprest/internal/infrastructure/auth"
	"strconv"

	"github.com/labstack/echo/v4"
)

type NavigationHandler struct {
	service *service.NavigationService
	session *auth.SessionService
}

func NewNavigationHandler(e *echo.Echo,
	service *service.NavigationService,
	session *auth.SessionService) *NavigationHandler {
	Handler := &NavigationHandler{
		service: service,
		session: session,
	}
	return Handler
}

func (nh *NavigationHandler) GetAllNavigation(c echo.Context) error {
	fmt.Println("Navigation: GetAllNavigation")
	navigation, err := nh.service.GetAllNavigation()
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, navigation)
}

func (nh *NavigationHandler) GetNavigation(c echo.Context) error {
	fmt.Println("Navigation: GetNavigation")
	idParam := c.Param("navigation_ID")
	fmt.Printf("DEBUG: Raw ID from URL: [%s] (Length: %d)\n", idParam, len(idParam))

	navigationIDuint64, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}

	navigationID := uint(navigationIDuint64)
	navigation, err := nh.service.GetNavigation(navigationID)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, navigation)
}

func (nh *NavigationHandler) CreateNavigation(c echo.Context) error {
	fmt.Println("Navigation: CreateNavigation")
	navigation := new(entity.Navigation)
	if err := c.Bind(navigation); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("Failed to parse request body: %s", err),
		})
	}
	administratorID, ok := nh.session.GetUserID(c)
	if ok == false {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to authenticate"),
		})
	}
	navigation.Administrators_Administrators_ID = uint(administratorID.(int32))
	err := nh.service.CreateNavigation(navigation)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, "Record created")
}

func (nh *NavigationHandler) UpdateNavigation(c echo.Context) error {
	fmt.Println("Navigation: UpdateNavigation")
	idParam := c.Param("navigation_ID")
	navigationIDuint64, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": "Invalid ID in URL",
		})
	}
	navigationID := uint(navigationIDuint64)

	navigation := new(entity.Navigation)
	if err := c.Bind(navigation); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to parse body",
		})
	}

	navigation.Navigation_ID = navigationID

	administratorID, ok := nh.session.GetUserID(c)
	if ok == false {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to authenticate"),
		})
	}
	navigation.Administrators_Administrators_ID = uint(administratorID.(int32))

	fmt.Printf("DEBUG: About to update Navigation ID %d with Name: %s\n", navigation.Navigation_ID, navigation.Name)

	err = nh.service.UpdateNavigation(navigation)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, "Record updated")
}

func (nh *NavigationHandler) DeleteNavigation(c echo.Context) error {
	fmt.Println("Navigation: DeleteNavigation")
	navigationIDuint64, err := strconv.ParseUint(c.Param("navigation_ID"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	navigationID := uint(navigationIDuint64)
	err = nh.service.DeleteNavigation(navigationID)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, "Record Removed")
}

func (nh *NavigationHandler) GetNavigationSchema(c echo.Context) error {
	schema := schema.GenerateSchema(entity.Navigation{})
	return c.JSON(http.StatusOK, schema)
}
