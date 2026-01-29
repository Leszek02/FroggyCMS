package main

import (
	"context"
	"encoding/gob"
	"fmt"

	"prestaprest/internal/application/service"
	"prestaprest/internal/infrastructure/auth"
	"prestaprest/internal/infrastructure/postgres"
	template "prestaprest/internal/infrastructure/templates"
	"prestaprest/internal/interface/api"
	handler "prestaprest/internal/interface/api/handlers"
	session "prestaprest/internal/interface/api/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type DBConfig struct {
	Host     string `env:"DB_HOST"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	DBname   string `env:"DB_NAME"`
	Port     string `env:"DB_PORT"`
	SSLmode  string `env:"DB_SSLMODE"`
}

func main() {
	//TODO: ADD TEMPLATE EDITING: footer, page, product_page, search, shop, ??? cart, checkout,
	//TODO: BLOCK EDITING: search, cart,
	//TODO: Delete all prints

	//TODO: Add proper error handling in shop pages
	//TODO: Add order.html to adminPanel

	ctx := context.Background()
	db := ("host=localhost user=postgres password=postgres dbname=postgres port=5433 sslmode=disable")
	conn, err := postgres.NewConnection(ctx, db)
	if err != nil {
		fmt.Println(err)
	} else {
		sqlDB, err := conn.DB()
		if err == nil {
			defer sqlDB.Close()
		}
		fmt.Println("DEBUG: database connected")
	}

	authKey := "Llf+Oj0x88iDaomsS+ARmlPTcHnM8SEL"
	encryptionKey := "JrGb6V33DIEOL3Prj3kgJYaNIZqPTpiN"

	gob.Register(map[uint64]int{})
	auth.InitSessionStore(authKey, encryptionKey)
	if auth.Store == nil {
		fmt.Println("ERROR: auth.Store is still nil after InitSessionStore!")
	} else {
		fmt.Println("SUCCESS: auth.Store initialized")
	}

	sessionService := auth.NewSessionService(auth.Store)
	SessionMiddleware := session.NewSessionMiddleware(sessionService)

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Logger.SetLevel(log.DEBUG)
	e.Use(middleware.Recover())
	fmt.Println("DEBUG: Echo server initiatied")

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.DELETE},
	}))

	// Render templates
	t := template.RenderTemplates()
	e.Renderer = t
	fmt.Println("DEBUG: Templates rendered")

	// Create services
	pageService := service.NewPageService(conn)
	navigationService := service.NewNavigationService(conn)
	clientService := service.NewClientService(conn)
	productService := service.NewProductService(conn)
	orderService := service.NewOrderService(conn)
	dictionaryService := service.NewDictionaryService(conn)
	administratorService := service.NewAdministratorService(conn)
	mediaService := service.NewMediaService(conn)
	fmt.Println("DEBUG: Services created")

	// Create handlers
	pageHandler := handler.NewPageHandler(e,
		&pageService,
		&navigationService,
		&clientService,
		sessionService,
		&productService,
		&dictionaryService,
		&orderService,
		t)
	navigationHandler := handler.NewNavigationHandler(e, &navigationService, sessionService)
	clientHandler := handler.NewClientHandler(e, &clientService)
	productHandler := handler.NewProductHandler(e, &productService, &dictionaryService, &mediaService)
	cartHandler := handler.NewCartHandler(e, sessionService)
	orderHandler := handler.NewOrderHandler(e, &orderService, &dictionaryService, &clientService, sessionService)
	dictionaryHandler := handler.NewDictionaryHandler(e, &dictionaryService)
	administratorHandler := handler.NewAdministratorHandler(e, &administratorService)
	authHandler := handler.NewAuthHandler(&clientService, &administratorService, sessionService)
	fmt.Println("DEBUG: Handlers created")

	// Attach controllers
	router := api.NewRouter(e,
		*SessionMiddleware,
		*pageHandler,
		*navigationHandler,
		*clientHandler,
		*productHandler,
		*cartHandler,
		*orderHandler,
		*dictionaryHandler,
		*administratorHandler,
		*authHandler)
	router.Setup(e)
	fmt.Println("DEBUG: Controllers created")

	port := ":8080"
	e.Logger.Fatal(e.Start("0.0.0.0" + port))

}
