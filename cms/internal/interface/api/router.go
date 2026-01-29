package api

import (
	"fmt"
	"net/http"
	handler "prestaprest/internal/interface/api/handlers"
	session "prestaprest/internal/interface/api/middleware"

	"github.com/labstack/echo/v4"
	// "path/filepath"
)

type Router struct {
	sessionMiddleware    session.SessionMiddleware
	pageHandler          handler.PageHandler
	navigationHandler    handler.NavigationHandler
	clientHandler        handler.ClientHandler
	productHandler       handler.ProductHandler
	cartHandler          handler.CartHandler
	orderHandler         handler.OrderHandler
	dictionaryHandler    handler.DictionaryHandler
	administratorHandler handler.AdministratorHandler
	authHandler          handler.AuthHandler
}

func NewRouter(e *echo.Echo,
	sessionMiddleware session.SessionMiddleware,
	pageHandler handler.PageHandler,
	navigationHandler handler.NavigationHandler,
	clientHandler handler.ClientHandler,
	productHandler handler.ProductHandler,
	cartHandler handler.CartHandler,
	orderHandler handler.OrderHandler,
	dictionaryHandler handler.DictionaryHandler,
	administratorHandler handler.AdministratorHandler,
	authHandler handler.AuthHandler) *Router {
	router := &Router{
		sessionMiddleware:    sessionMiddleware,
		pageHandler:          pageHandler,
		navigationHandler:    navigationHandler,
		clientHandler:        clientHandler,
		productHandler:       productHandler,
		cartHandler:          cartHandler,
		orderHandler:         orderHandler,
		dictionaryHandler:    dictionaryHandler,
		administratorHandler: administratorHandler,
		authHandler:          authHandler,
	}

	return router
}

func (r *Router) Setup(e *echo.Echo) {

	// Prepare admin endpoints
	admin := e.Group("")
	admin.Use(r.sessionMiddleware.RequireAdmin)

	// Expose public folder
	e.Static("/static", "../public")
	// e.Static("/admin-static", "../public/admin")
	admin.GET("/admin-static/test_admin_panel", func(c echo.Context) error {
		return c.File(fmt.Sprintf("../template/admin/test_admin_panel.html"))
	})

	admin.GET("/admin-static/navigation", func(c echo.Context) error {
		return c.File(fmt.Sprintf("../template/admin/navigation.html"))
	})

	admin.GET("/admin-static/delivery_dictionary", func(c echo.Context) error {
		return c.File(fmt.Sprintf("../template/admin/delivery_dictionary.html"))
	})

	admin.GET("/admin-static/payment_dictionary", func(c echo.Context) error {
		return c.File(fmt.Sprintf("../template/admin/payment_dictionary.html"))
	})

	admin.GET("/admin-static/vat_percentage_dictionary", func(c echo.Context) error {
		return c.File(fmt.Sprintf("../template/admin/vat_percentage_dictionary.html"))
	})

	admin.GET("/admin-static/product_category_dictionary", func(c echo.Context) error {
		return c.File(fmt.Sprintf("../template/admin/product_category_dictionary.html"))
	})

	admin.GET("/admin-static/pages", func(c echo.Context) error {
		return c.File(fmt.Sprintf("../template/admin/pages.html"))
	})

	admin.GET("/admin-static/products", func(c echo.Context) error {
		return c.File(fmt.Sprintf("../template/admin/products.html"))
	})

	admin.GET("/admin-static/create", func(c echo.Context) error {
		return c.File(fmt.Sprintf("../template/admin/create.html"))
	})

	admin.GET("/admin-static/edit", func(c echo.Context) error {
		return c.File("../template/admin/edit.html")
	})

	admin.GET("/admin-static/edit_template_products", func(c echo.Context) error {
		return c.File("../template/admin/edit_template_products.html")
	})

	admin.GET("/admin-static/administrators", func(c echo.Context) error {
		return c.File("../template/admin/administrators.html")
	})

	admin.GET("/admin-static/client", func(c echo.Context) error {
		return c.File("../template/admin/client.html")
	})

	admin.GET("/admin-static/order", func(c echo.Context) error {
		return c.File("../template/admin/order.html")
	})

	// Expose public folder
	e.Static("/static/*", "../public")
	e.Static("/media/*", "../media")
	//login/register hanglers
	e.GET("/login", r.sessionMiddleware.IsLogged(func(c echo.Context) error {
		return c.Render(http.StatusOK, fmt.Sprintf("login"), map[string]any{
			"LoginType": "client",
		})
	}))
	e.GET("/register", r.sessionMiddleware.IsLogged(func(c echo.Context) error {
		return r.pageHandler.ReturnPage("register.html", c)
	}))
	e.GET("/admin/login", r.sessionMiddleware.IsAdminLogged(func(c echo.Context) error {
		return c.Render(http.StatusOK, fmt.Sprintf("login"), map[string]any{
			"LoginType": "admin",
		})
	}))
	e.GET("/logout", r.authHandler.Logout)
	e.POST("/login", r.authHandler.Login)
	e.POST("/admin/login", r.authHandler.AdminLogin)
	e.POST("/register", r.authHandler.Register)

	//shop redirections handler  (Follow order <template_name>/<page_name or id or something page specific>)
	e.GET("/shop/:page_name", func(c echo.Context) error {
		return r.pageHandler.RenderPage("shop", c)
	})
	e.GET("/cart", func(c echo.Context) error {
		return r.pageHandler.RenderCartPage("cart", c)
	})
	e.GET("/search/:page_name", func(c echo.Context) error {
		return r.pageHandler.RenderPage("search", c)
	})
	e.GET("/product/:product_id", func(c echo.Context) error {
		return r.pageHandler.RenderProductPage("products", c)
	})
	e.GET("/checkout", r.sessionMiddleware.RequireAuth(func(c echo.Context) error {
		return r.pageHandler.RenderCheckoutPage("checkout", c)
	}))

	e.GET("/page/:page_name", func(c echo.Context) error {
		return r.pageHandler.RenderPage("page", c)
	})

	e.GET("/profile", r.sessionMiddleware.RequireAuth(func(c echo.Context) error {
		return r.pageHandler.RenderProfilePage("profile", c)
	}))

	//Cart handler
	e.PATCH("/carts", r.cartHandler.UpdateCart)
	e.DELETE("/carts", r.cartHandler.DeleteCart)
	e.DELETE("/carts/:product_ID", r.cartHandler.DeleteCartItem)

	//Order handler
	admin.GET("/orders", r.orderHandler.GetAllOrders)
	e.GET("/orders/:order_ID", r.sessionMiddleware.IsLogged(r.orderHandler.GetOrder))
	e.POST("/orders", r.sessionMiddleware.RequireAuth(r.orderHandler.CreateOrder))
	e.PATCH("/orders/:order_ID", r.sessionMiddleware.IsLogged(r.orderHandler.UpdateOrder))
	admin.DELETE("/orders/:order_ID", r.sessionMiddleware.IsLogged(r.orderHandler.DeleteOrder))

	//Navigation handler
	e.GET("/navigation", r.navigationHandler.GetAllNavigation)
	e.GET("/navigation/:navigation_ID", r.navigationHandler.GetNavigation)
	admin.POST("/navigation", r.navigationHandler.CreateNavigation)
	admin.PATCH("/navigation/:navigation_ID", r.navigationHandler.UpdateNavigation)
	admin.DELETE("/navigation/:navigation_ID", r.navigationHandler.DeleteNavigation)
	admin.GET("/navigation/schema", r.navigationHandler.GetNavigationSchema)

	//Page handler
	e.GET("/pages", r.pageHandler.GetAllPage)
	e.GET("/pages/:page_ID", r.pageHandler.GetPageByID)
	admin.POST("/pages", r.pageHandler.CreatePage)
	admin.PATCH("/pages/:page_ID", r.pageHandler.UpdatePage)
	admin.DELETE("/pages/:page_ID", r.pageHandler.DeletePage)
	admin.GET("/pages/schema", r.pageHandler.GetPageSchema)

	//Client handler
	admin.GET("/client", r.clientHandler.GetAllClient)
	admin.GET("/client/:client_ID", r.clientHandler.GetClient)
	e.PATCH("/client/:client_ID", r.clientHandler.UpdateClient)
	admin.DELETE("/client/:client_ID", r.clientHandler.DeleteClient)
	admin.GET("/client/schema", r.clientHandler.GetClientSchema)
	e.PATCH("/client/password/:client_ID", r.authHandler.ChangePassword)

	//Product handler
	admin.GET("/products", r.productHandler.GetAllProduct)
	admin.GET("/products/:product_ID", r.productHandler.GetProduct)
	admin.POST("/products", r.productHandler.CreateProduct)
	admin.PATCH("/products/:product_ID", r.productHandler.UpdateProduct)
	admin.DELETE("/products/:product_ID", r.productHandler.DeleteProduct)
	admin.GET("/products/schema", r.productHandler.GetProductSchema)
	admin.GET("/media", r.productHandler.GetProductPhotos)

	//Administrator handler
	admin.GET("/administrators", r.administratorHandler.GetAllAdministrators)
	admin.GET("/administrators/:administrator_ID", r.administratorHandler.GetAdministrator)
	admin.POST("/administrators", r.administratorHandler.CreateAdministrator)
	admin.PATCH("/administrators/:administrator_ID", r.administratorHandler.UpdateAdministrator)
	admin.DELETE("/administrators/:administrator_ID", r.administratorHandler.DeleteAdministrator)
	admin.GET("/administrators/schema", r.administratorHandler.GetAdministratorSchema)

	//Dictionary handlers
	admin.GET("/delivery_dictionary", r.dictionaryHandler.GetAllDeliveryDictionaries)
	admin.GET("/delivery_dictionary/:delivery_dictionary_ID", r.dictionaryHandler.GetDeliveryDictionary)
	admin.POST("/delivery_dictionary", r.dictionaryHandler.CreateDeliveryDictionary)
	admin.PATCH("/delivery_dictionary/:delivery_dictionary_ID", r.dictionaryHandler.UpdateDeliveryDictionary)
	admin.DELETE("/delivery_dictionary/:delivery_dictionary_ID", r.dictionaryHandler.DeleteDeliveryDictionary)
	admin.GET("/delivery_dictionary/schema", r.dictionaryHandler.GetDeliveryDictionarySchema)

	admin.GET("/payment_dictionary", r.dictionaryHandler.GetAllPaymentDictionaries)
	admin.GET("/payment_dictionary/:payment_dictionary_ID", r.dictionaryHandler.GetPaymentDictionary)
	admin.POST("/payment_dictionary", r.dictionaryHandler.CreatePaymentDictionary)
	admin.PATCH("/payment_dictionary/:payment_dictionary_ID", r.dictionaryHandler.UpdatePaymentDictionary)
	admin.DELETE("/payment_dictionary/:payment_dictionary_ID", r.dictionaryHandler.DeletePaymentDictionary)
	admin.GET("/payment_dictionary/schema", r.dictionaryHandler.GetPaymentDictionarySchema)

	admin.GET("/vat_percentage_dictionary", r.dictionaryHandler.GetAllVATPercentageDictionary)
	admin.GET("/vat_percentage_dictionary/:vat_percentage_dictionary_ID", r.dictionaryHandler.GetVATPercentageDictionary)
	admin.POST("/vat_percentage_dictionary", r.dictionaryHandler.CreateVATPercentageDictionary)
	admin.PATCH("/vat_percentage_dictionary/:vat_percentage_dictionary_ID", r.dictionaryHandler.UpdateVATPercentageDictionary)
	admin.DELETE("/vat_percentage_dictionary/:vat_percentage_dictionary_ID", r.dictionaryHandler.DeleteVATPercentageDictionary)
	admin.GET("/vat_percentage_dictionary/schema", r.dictionaryHandler.GetVATPercentageDictionarySchema)

	admin.GET("/product_category_dictionary", r.dictionaryHandler.GetAllProductCategoryDictionaries)
	admin.GET("/product_category_dictionary/:product_category_dictionary_ID", r.dictionaryHandler.GetProductCategoryDictionary)
	admin.POST("/product_category_dictionary", r.dictionaryHandler.CreateProductCategoryDictionary)
	admin.PATCH("/product_category_dictionary/:product_category_dictionary_ID", r.dictionaryHandler.UpdateProductCategoryDictionary)
	admin.DELETE("/product_category_dictionary/:product_category_dictionary_ID", r.dictionaryHandler.DeleteProductCategoryDictionary)
	admin.GET("/product_category_dictionary/schema", r.dictionaryHandler.GetProductCategorySchema)

	e.GET("/*", r.pageHandler.GetNotFoundPage)

}

func LoggedClientEndpoint(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("CLIENT ROUTE 🐈🐈🐈")
		err := h(c)
		if err != nil {
			fmt.Printf("Error: %v", err)
		}
		return err
	}
}
