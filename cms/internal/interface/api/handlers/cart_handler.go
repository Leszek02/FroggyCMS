package handler

import (
	"errors"
	"fmt"
	"net/http"
	"prestaprest/internal/infrastructure/auth"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CartHandler struct {
	session *auth.SessionService
}

func NewCartHandler(e *echo.Echo, session *auth.SessionService) *CartHandler {
	Handler := &CartHandler{
		session: session,
	}
	return Handler
}

func (ch *CartHandler) UpdateCart(c echo.Context) error {
	fmt.Println("Cart: UpdateCart")
	var data map[string]int

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("Failed to parse request data: %s", err),
		})
	}
	err := ch.validateCart(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("Validation error: %s", err),
		})
	}

	if !(ch.session.IsAuthenticated(c) && !ch.session.IsAdmin(c)) {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to add product to the cart: You're not logged in as a Client"),
		})
	}

	fmt.Println("Cart: Validation passed")
	newCartIDParam := data["product_id"]
	quantityParam, ok := data["quantity"]
	fmt.Println("Cart: Parsed params:", newCartIDParam, quantityParam)
	newQuantityParam := data["newQuantity"]
	cart, err := ch.session.GetCartItems(c)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to retrieve cart from session: %s", err),
		})
	}
	fmt.Println("Cart: Current cart items:", cart)
	fmt.Println("Cart: Parsed params:", newCartIDParam, quantityParam)

	if ok {
		err = ch.session.AddToCart(c, uint64(newCartIDParam), quantityParam)
	} else {
		err = ch.session.UpdateCartItemQuantity(c, uint64(newCartIDParam), newQuantityParam)
	}

	// err = ch.session.UpdateCart(c, cart)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	fmt.Println("Cart: Added product", newCartIDParam, "to cart")
	return c.JSON(http.StatusOK, "Cart updated")
}

func (ch *CartHandler) DeleteCart(c echo.Context) error {
	fmt.Println("Cart: DeleteCart")
	err := ch.session.ClearCart(c)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to clear cart: %s", err),
		})
	}
	return c.JSON(http.StatusOK, "Record Removed")
}

func (ch *CartHandler) DeleteCartItem(c echo.Context) error {
	fmt.Println("Cart: DeleteCartItem")
	productIDParam := c.Param("product_ID")

	productIDParamUint64, err := strconv.ParseUint(productIDParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("Failed to parse product ID: %s", err),
		})
	}

	err = ch.session.RemoveFromCart(c, productIDParamUint64)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to remove item from cart: %s", err),
		})
	}
	return c.JSON(http.StatusOK, "Item Removed")
}

func (ch *CartHandler) validateCart(data map[string]int) error {

	productID, ok := data["product_id"]
	if !ok || productID <= 0 {
		return errors.New("product ID is required")
	}

	quantity, ok := data["quantity"]
	if !ok || quantity <= 0 {
		NewQuantity, ok := data["newQuantity"]
		if !ok || NewQuantity <= 0 {
			return errors.New("quantity is required")
		}
	}
	return nil
}
