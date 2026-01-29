package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"prestaprest/internal/application/schema"
	"prestaprest/internal/application/service"
	"prestaprest/internal/domain/entity"
	"prestaprest/internal/infrastructure/auth"
	templates "prestaprest/internal/infrastructure/templates"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PageHandler struct {
	service           *service.PageService
	navigationService *service.NavigationService
	clientService     *service.ClientService
	sessionService    *auth.SessionService
	productService    *service.ProductService
	dictionaryService *service.DictionaryService
	orderService      *service.OrderService
	template          *templates.Template
}

func NewPageHandler(e *echo.Echo,
	service *service.PageService,
	navigationService *service.NavigationService,
	clientService *service.ClientService,
	sessionService *auth.SessionService,
	productService *service.ProductService,
	dictionaryService *service.DictionaryService,
	orderService *service.OrderService,
	template *templates.Template) *PageHandler {
	Handler := &PageHandler{
		service:           service,
		navigationService: navigationService,
		clientService:     clientService,
		sessionService:    sessionService,
		productService:    productService,
		dictionaryService: dictionaryService,
		orderService:      orderService,
		template:          template,
	}

	return Handler
}

func (pc *PageHandler) RenderPage(template string, c echo.Context) error {
	fmt.Println("Shop: Validate incoming data here and extract needed parameters (if needed)")
	pageName := c.Param("page_name")
	page, err := pc.service.GetPage(pageName)
	fmt.Println("Am here? (api)")
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": "Failed to retrieve Page: failed to parse page content",
		})
	}

	var content map[string]any
	if err := json.Unmarshal([]byte(page.Content), &content); err != nil {
		content = make(map[string]any)
	}
	navigations, err := pc.navigationService.GetAllNavigation()
	if err != nil {
		fmt.Println("Failed to get navigation:", err)
		navigations = []*entity.Navigation{}
	}

	content["Navigations"] = navigations
	if pc.sessionService.IsAuthenticated(c) {
		content["IsLogged"] = true
	} else {
		content["IsLogged"] = false
	}
	content["CartQuantity"] = pc.sessionService.GetCartQuantity(c)

	if template == "search" {
		products, err := pc.productService.GetAllProduct()
		if err != nil {
			fmt.Println("Failed to get products:", err)
			products = []entity.Product{}
		}
		content["Products"] = products
	}

	fmt.Println(content)
	fmt.Println("Am before")
	err = c.Render(http.StatusOK, template, content)
	if err != nil {
		fmt.Println("RENDER ERROR:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": fmt.Sprintf("Failed to render template: %s", err),
		})
	}
	return nil
}

func (pc *PageHandler) RenderProductPage(template string, c echo.Context) error {
	fmt.Println("RenderProductPage: Validate incoming data here and extract needed parameters (if needed)")
	productID, err := strconv.ParseUint(c.Param("product_id"), 10, 64)
	product, err := pc.productService.GetProduct(productID)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": "Failed to retrieve Product",
		})
	}

	navigations, err := pc.navigationService.GetAllNavigation()
	if err != nil {
		fmt.Println("Failed to get navigation:", err)
		navigations = []*entity.Navigation{}
	}

	var content map[string]any
	if err := json.Unmarshal([]byte(""), &content); err != nil {
		content = make(map[string]any)
	}

	page, err := pc.service.GetPage("products")

	var options map[string]interface{}
	if page != nil && page.Content != "" {
		if err := json.Unmarshal([]byte(page.Content), &options); err != nil {
			fmt.Println("Failed to parse options:", err)
			options = make(map[string]interface{})
		}
	} else {
		options = make(map[string]interface{})
	}

	content["Options"] = options

	fmt.Println(product)

	content["Product"] = product
	content["Navigations"] = navigations
	if pc.sessionService.IsAuthenticated(c) {
		content["IsLogged"] = true
	} else {
		content["IsLogged"] = false
	}
	content["CartQuantity"] = pc.sessionService.GetCartQuantity(c)
	content["Quantity"] = 1
	err = c.Render(http.StatusOK, template, content)
	if err != nil {
		log.Printf("Template render error: %v", err)
		log.Printf("Template name: %s", template)
		log.Printf("Content: %+v", content)
		return c.String(500, fmt.Sprintf("Render error: %v", err))
	}
	return nil
}

func (pc *PageHandler) RenderCartPage(template string, c echo.Context) error {
	type cartItem struct {
		ProductID uint64
		Name      string
		Quantity  int
		Price     float32
		Total     float32
		Image     string
	}

	fmt.Println("RenderCartPage: Validate incoming data here and extract needed parameters (if needed)")
	page, err := pc.service.GetPage(template)
	fmt.Println("Am here? (page)")
	var content map[string]any
	if err := json.Unmarshal([]byte(page.Content), &content); err != nil {
		content = make(map[string]any)
	}
	navigations, err := pc.navigationService.GetAllNavigation()
	if err != nil {
		fmt.Println("Failed to get navigation:", err)
		navigations = []*entity.Navigation{}
	}

	content["Navigations"] = navigations
	fmt.Println("Is authenticated:", pc.sessionService.IsAuthenticated(c))
	if pc.sessionService.IsAuthenticated(c) {
		cartItems, err := pc.sessionService.GetCartItems(c)
		if err != nil {
			return c.JSON(http.StatusTeapot, map[string]string{
				"error": "Failed to retrieve Cart Items",
			})
		}
		var productList []*cartItem
		var productTotal float32
		for key, value := range cartItems {
			product, err := pc.productService.GetProduct(key)
			if err != nil {
				return c.JSON(http.StatusTeapot, map[string]string{
					"error": "Failed to retrieve Product",
				})
			}
			newItem := &cartItem{
				ProductID: product.ProductID,
				Name:      product.Name,
				Quantity:  value,
				Price:     product.Price,
				Total:     product.Price * float32(value),
				Image:     product.ProductIconImage,
			}
			productList = append(productList, newItem)
			productTotal += newItem.Total
		}
		content["Cart"] = productList
		content["CartTotal"] = productTotal
	} else {
		content["Cart"] = ""
	}
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": "Failed to retrieve Product",
		})
	}
	if pc.sessionService.IsAuthenticated(c) {
		content["IsLogged"] = true
	} else {
		content["IsLogged"] = false
	}
	content["CartQuantity"] = pc.sessionService.GetCartQuantity(c)
	fmt.Println("ADASDASDFSA")
	fmt.Println(content)
	err = c.Render(http.StatusOK, template, content)
	if err != nil {
		log.Printf("Template render error: %v", err)
		log.Printf("Template name: %s", template)
		log.Printf("Content: %+v", content)
		return c.String(500, fmt.Sprintf("Render error: %v", err))
	}
	return nil
}

func (pc *PageHandler) RenderCheckoutPage(template string, c echo.Context) error {
	type cartItem struct {
		ProductID uint64
		Name      string
		Quantity  int
		Price     float32
		Total     float32
	}
	fmt.Println("RenderCheckoutPage: Validate incoming data here and extract needed parameters (if needed)")
	page, err := pc.service.GetPage(template)
	fmt.Println("Am here? (page)")
	var content map[string]any
	if err := json.Unmarshal([]byte(page.Content), &content); err != nil {
		content = make(map[string]any)
	}
	if pc.sessionService.IsAuthenticated(c) {
		cartItems, err := pc.sessionService.GetCartItems(c)
		if err != nil {
			return c.JSON(http.StatusTeapot, map[string]string{
				"error": "Failed to retrieve Cart Items",
			})
		}
		var productList []*cartItem
		var productSubtotal, productTotal float32 = 0, 0
		for key, value := range cartItems {
			product, err := pc.productService.GetProduct(key)
			if err != nil {
				return c.JSON(http.StatusTeapot, map[string]string{
					"error": "Failed to retrieve Product",
				})
			}
			newItem := &cartItem{
				ProductID: product.ProductID,
				Name:      product.Name,
				Quantity:  value,
				Price:     product.Price,
				Total:     product.Price * float32(value),
			}
			productList = append(productList, newItem)
			productSubtotal += newItem.Total
			productTotal += newItem.Total
		}
		content["Cart"] = productList
		content["CartProductSubtotal"] = productSubtotal
		content["CartProductTotal"] = productTotal

		clientID, ok := pc.sessionService.GetUserID(c)
		if ok == false {
			return c.JSON(http.StatusTeapot, map[string]string{
				"error": "Failed to retrieve Client ID",
			})
		}
		convertedClientID, ok := clientID.(int64)
		client, err := pc.clientService.GetClient(convertedClientID)
		if client.Newsletter {
			content["Newsletter"] = true
		} else {
			content["Newsletter"] = false
		}
	} else {
		content["Cart"] = ""
	}
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": "Failed to retrieve Product",
		})
	}
	navigations, err := pc.navigationService.GetAllNavigation()
	if err != nil {
		fmt.Println("Failed to get navigation:", err)
		navigations = []*entity.Navigation{}
	}

	content["Navigations"] = navigations
	content["IsLogged"] = true
	content["CartQuantity"] = pc.sessionService.GetCartQuantity(c)

	paymentMethods, err := pc.dictionaryService.GetAllPaymentDictionary()
	if err != nil {
		fmt.Println("Failed to get payment methods:", err)
		paymentMethods = []entity.PaymentDictionary{}
	}
	content["PaymentMethods"] = paymentMethods

	deliveryMethods, err := pc.dictionaryService.GetAllDeliveryDictionary()
	if err != nil {
		fmt.Println("Failed to get delivery methods:", err)
		deliveryMethods = []entity.DeliveryDictionary{}
	}
	content["DeliveryMethods"] = deliveryMethods

	fmt.Println("ADASDASDFSA")
	fmt.Println(content)
	err = c.Render(http.StatusOK, template, content)
	if err != nil {
		log.Printf("Template render error: %v", err)
		log.Printf("Template name: %s", template)
		log.Printf("Content: %+v", content)
		return c.String(500, fmt.Sprintf("Render error: %v", err))
	}
	return nil
}

func (pc *PageHandler) RenderProfilePage(template string, c echo.Context) error {
	fmt.Println("RenderProfilePage: Validate incoming data here and extract needed parameters (if needed)")
	page, err := pc.service.GetPage(template)
	fmt.Println("Am here? (page)")
	var content map[string]any
	if err := json.Unmarshal([]byte(page.Content), &content); err != nil {
		content = make(map[string]any)
	}
	if pc.sessionService.IsAuthenticated(c) {
		clientID, ok := pc.sessionService.GetUserID(c)
		if ok == false {
			return c.JSON(http.StatusTeapot, map[string]string{
				"error": "Failed to retrieve Client ID",
			})
		}
		convertedClientID, ok := clientID.(int64)
		client, err := pc.clientService.GetClient(convertedClientID)
		if err != nil {
			return c.JSON(http.StatusTeapot, map[string]string{
				"error": "Failed to retrieve client",
			})
		}
		content["Client"] = client
		orderedProducts, err := pc.orderService.GetAllOrderProducts(client.Client_ID)

		content["Orders"] = orderedProducts
	}

	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": "Failed to retrieve Product",
		})
	}
	navigations, err := pc.navigationService.GetAllNavigation()
	if err != nil {
		fmt.Println("Failed to get navigation:", err)
		navigations = []*entity.Navigation{}
	}

	content["Navigations"] = navigations
	content["IsLogged"] = true
	content["CartQuantity"] = pc.sessionService.GetCartQuantity(c)

	fmt.Println("ADASDASDFSA")
	fmt.Println(content)
	err = c.Render(http.StatusOK, template, content)
	if err != nil {
		log.Printf("Template render error: %v", err)
		log.Printf("Template name: %s", template)
		log.Printf("Content: %+v", content)
		return c.String(500, fmt.Sprintf("Render error: %v", err))
	}
	return nil
}

// For login, register etc.
func (pc *PageHandler) ReturnPage(pageName string, c echo.Context) error {
	return c.File(fmt.Sprintf("../template/shop/%s", pageName))
}

func (pc *PageHandler) GetDashboard(c echo.Context) error {
	fmt.Println("Shop: Validate incoming data here and extract needed parameters (if needed)")
	template := "admin/dashboard"
	page, err := pc.service.GetPage("shop")
	fmt.Println("Am here? (api)")
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": "Failed to retrieve Page: failed to parse page content",
		})
	}
	var content map[string]any
	if err := json.Unmarshal([]byte(page.Content), &content); err != nil {
		content = nil
	}
	fmt.Println(content)
	return c.Render(http.StatusOK, template, content)
}

func (pc *PageHandler) GetAllPage(c echo.Context) error {
	fmt.Println("Navigation: GetAllNavigation")
	page, err := pc.service.GetAllPages()
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, page)
}

func (pc *PageHandler) GetPageByID(c echo.Context) error {
	fmt.Println("Page: GetPage")
	idParam := c.Param("page_ID")
	fmt.Printf("DEBUG: Raw ID from URL: [%s] (Length: %d)\n", idParam, len(idParam))

	pageIDuint64, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}

	pageID := uint(pageIDuint64)
	page, err := pc.service.GetPageByID(pageID)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, page)
}

func (pc *PageHandler) CreatePage(c echo.Context) error {
	fmt.Println("Page: CreatePage")
	page := new(entity.Page)
	if err := c.Bind(page); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("Failed to parse request body: %s", err),
		})
	}
	administratorID, ok := pc.sessionService.GetUserID(c)
	if ok == false {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to authenticate"),
		})
	}
	page.Administrators_Administrator_ID = int(administratorID.(int32))
	err := pc.service.CreatePage(page)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, "Record created")
}

func (pc *PageHandler) UpdatePage(c echo.Context) error {
	fmt.Println("Page: UpdatePage")
	idParam := c.Param("page_ID")
	pageIDuint64, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": "Invalid ID in URL",
		})
	}
	pageID := uint(pageIDuint64)

	page := new(entity.Page)
	if err := c.Bind(page); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to parse body",
		})
	}

	page.Page_ID = pageID

	fmt.Printf("DEBUG: About to update Page ID %d with Name: %s\n", page.Page_ID, page.Name)

	err = pc.service.UpdatePage(page)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, "Record updated")
}

func (pc *PageHandler) DeletePage(c echo.Context) error {
	fmt.Println("Page: DeletePage")
	pageIDuint64, err := strconv.ParseUint(c.Param("page_ID"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	pageID := uint(pageIDuint64)
	err = pc.service.DeletePage(pageID)
	if err != nil {
		return c.JSON(http.StatusTeapot, map[string]string{
			"error": fmt.Sprintf("Failed to complete request: %s", err),
		})
	}
	return c.JSON(http.StatusOK, "Record Removed")
}

func (pc *PageHandler) GetNotFoundPage(c echo.Context) error {
	c.Response().Status = http.StatusNotFound
	return c.File("../template/shop/not_found.html")
}

func (pc *PageHandler) GetPageSchema(c echo.Context) error {
	schema := schema.GenerateSchema(entity.Page{})
	return c.JSON(http.StatusOK, schema)
}
