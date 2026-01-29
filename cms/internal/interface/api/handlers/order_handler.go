package handler

import (
	"errors"
	"fmt"
	"net/http"
	"prestaprest/internal/application/service"
	"prestaprest/internal/domain/entity"
	"prestaprest/internal/infrastructure/auth"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	service    *service.OrderService
	dictionary *service.DictionaryService
	client     *service.ClientService
	session    *auth.SessionService
}

func NewOrderHandler(e *echo.Echo, service *service.OrderService, dictionary *service.DictionaryService,
	client *service.ClientService, session *auth.SessionService) *OrderHandler {
	Handler := &OrderHandler{
		service:    service,
		dictionary: dictionary,
		client:     client,
		session:    session,
	}
	return Handler
}

func (oh *OrderHandler) GetAllOrders(c echo.Context) error {
	fmt.Println("Order: GetAllOrder")
	order, err := oh.service.GetAllOrders()
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, order)
}

func (oh *OrderHandler) GetOrder(c echo.Context) error {
	fmt.Println("Order: GetOrder")
	orderID, err := strconv.ParseUint(c.Param("order_ID"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	order, err := oh.service.GetOrder(orderID)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, order)
}

func (oh *OrderHandler) CreateOrder(c echo.Context) error {
	fmt.Println("Order: CreateOrder")
	err := oh.validateOrder(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("Validation error: %s", err),
		})
	}
	fmt.Println("Passed validation")
	order := new(entity.Order)
	if err := c.Bind(order); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("Failed to parse request body: %s", err),
		})
	}

	deliveryDictID, err := strconv.ParseUint(strings.TrimSpace(c.FormValue("delivery_method")), 10, 32)
	paymentDictID, err := strconv.ParseUint(strings.TrimSpace(c.FormValue("payment_method")), 10, 32)
	deliveryMethod, err := oh.dictionary.GetDeliveryDictionary(uint(deliveryDictID))
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	paymentMethod, err := oh.dictionary.GetPaymentDictionary(uint(paymentDictID))
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}

	order.SetDeliveryMethod(deliveryMethod)
	order.SetPaymentMethod(paymentMethod)
	order.UpdateStatus(true)
	fmt.Println("DeliveryMethod: ", deliveryMethod)
	fmt.Println("PaymentMethod: ", paymentMethod)
	fmt.Println("Order after binding:", order)
	clientID, ok := oh.session.GetUserID(c)
	fmt.Println("Retrieved client ID from session:", clientID)
	if ok == false {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("User not authenticated: %s", err),
		})
	}
	convertedClientID, ok := clientID.(int64)
	fmt.Println("Converted client ID:", convertedClientID)
	if ok == false {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to convert client ID: %s", err),
		})
	}
	fmt.Println("Converted client ID:", convertedClientID)
	order.SetClientID(uint64(convertedClientID))
	fmt.Println("Order after setting client ID:", order)

	err = oh.service.CreateOrder(order)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}

	cart, err := oh.session.GetCartItems(c)
	totalPrice, err := oh.service.CreateOrderProduct(cart, order.Order_ID)

	totalPrice += deliveryMethod.Price
	totalPrice += paymentMethod.Price
	order.Total_price = totalPrice
	err = oh.service.UpdateOrder(order)

	err = oh.session.ClearCart(c)
	newsletterStr := strings.TrimSpace(c.FormValue("Newsletter"))
	setNewsletter := false
	if newsletterStr == "on" {
		fmt.Println("Failed to parse newsletter: ", err)
		setNewsletter = true
	}
	fmt.Println("Newsletter: ", setNewsletter)
	if setNewsletter {
		client, err := oh.client.GetClient(convertedClientID)
		if err != nil {
			return c.JSON(http.StatusTeapot, map[string]string{
				"error": fmt.Sprintf("Failed to complete request: %s", err),
			})
		}
		client.Client_ID = uint64(convertedClientID)
		client.Newsletter = true
		err = oh.client.UpdateClient(client)
		if err != nil {
			return c.JSON(http.StatusTeapot, map[string]string{
				"error": fmt.Sprintf("Failed to complete request: %s", err),
			})
		}
	}
	return c.Redirect(302, "/cart")
}

func (oh *OrderHandler) UpdateOrder(c echo.Context) error {
	fmt.Println("Order: UpdateOrder")
	err := oh.validateOrder(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("Validation error: %s", err),
		})
	}
	orderIDuint64, err := strconv.ParseUint(c.Param("order_ID"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	orderID := uint(orderIDuint64)
	order := new(entity.Order)
	if err := c.Bind(order); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("Failed to parse request body: %s", err),
		})
	}
	order.SetID(orderID)
	err = oh.service.UpdateOrder(order)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, "Record updated")
}

func (oh *OrderHandler) DeleteOrder(c echo.Context) error {
	fmt.Println("Order: DeleteOrder")
	orderIDuint64, err := strconv.ParseUint(c.Param("order_ID"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	orderID := uint(orderIDuint64)
	err = oh.service.DeleteOrder(orderID)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, "Record Removed")
}

func (oh *OrderHandler) validateOrder(c echo.Context) error {
	Delivery_dict_ID := strings.TrimSpace(c.FormValue("delivery_method"))
	if Delivery_dict_ID == "" {
		return errors.New("Delivery dictionary ID is required")
	}
	Payment_dict_ID := strings.TrimSpace(c.FormValue("payment_method"))
	if Payment_dict_ID == "" {
		return errors.New("Payment dictionary ID is required")
	}

	address_city := strings.TrimSpace(c.FormValue("Address_city"))
	if address_city == "" {
		return errors.New("Address city is required")
	}

	address_postcode := strings.TrimSpace(c.FormValue("Address_postcode"))
	if address_postcode == "" {
		return errors.New("Address postcode is required")
	}

	address_street := strings.TrimSpace(c.FormValue("Address_street"))
	if address_street == "" {
		return errors.New("Address street is required")
	}

	address_state := strings.TrimSpace(c.FormValue("Address_state"))
	if address_state == "" {
		return errors.New("Address state is required")
	}

	return nil
}
