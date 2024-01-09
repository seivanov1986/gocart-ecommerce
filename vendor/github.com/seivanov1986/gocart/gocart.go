package gocart

import (
	"github.com/seivanov1986/gocart/external/cache"
	"github.com/seivanov1986/sql_client"

	"github.com/seivanov1986/gocart/internal/service/image"
	"github.com/seivanov1986/gocart/internal/service/yandex_feed"

	"net/http"

	"github.com/gorilla/mux"

	"github.com/seivanov1986/gocart/external/cache_builder"
	"github.com/seivanov1986/gocart/external/cache_service"
	"github.com/seivanov1986/gocart/external/widget_manager"
	"github.com/seivanov1986/gocart/internal/http/attribute"
	"github.com/seivanov1986/gocart/internal/http/auth"
	"github.com/seivanov1986/gocart/internal/http/category"
	commonHandle "github.com/seivanov1986/gocart/internal/http/common"
	"github.com/seivanov1986/gocart/internal/http/file"
	imageHandle "github.com/seivanov1986/gocart/internal/http/image"
	"github.com/seivanov1986/gocart/internal/http/image_to_category"
	"github.com/seivanov1986/gocart/internal/http/image_to_product"
	"github.com/seivanov1986/gocart/internal/http/page"
	"github.com/seivanov1986/gocart/internal/http/product"
	"github.com/seivanov1986/gocart/internal/http/product_to_category"
	"github.com/seivanov1986/gocart/internal/http/sefurl"
	"github.com/seivanov1986/gocart/internal/http/user"
	yandexHandle "github.com/seivanov1986/gocart/internal/http/yandex"
	auth2 "github.com/seivanov1986/gocart/internal/middleware/auth"
	"github.com/seivanov1986/gocart/internal/middleware/common"
	"github.com/seivanov1986/gocart/internal/middleware/cors"
	"github.com/seivanov1986/gocart/internal/repository"
	attributeService "github.com/seivanov1986/gocart/internal/service/attribute"
	attributeToProductService "github.com/seivanov1986/gocart/internal/service/attribute_to_product"
	authService "github.com/seivanov1986/gocart/internal/service/auth"
	categoryService "github.com/seivanov1986/gocart/internal/service/category"
	commonService "github.com/seivanov1986/gocart/internal/service/common"
	imageToCategoryService "github.com/seivanov1986/gocart/internal/service/image_to_category"
	imageToProductService "github.com/seivanov1986/gocart/internal/service/image_to_product"
	pageService "github.com/seivanov1986/gocart/internal/service/page"
	productService "github.com/seivanov1986/gocart/internal/service/product"
	productToCategoryService "github.com/seivanov1986/gocart/internal/service/product_to_category"
	sefUrlService "github.com/seivanov1986/gocart/internal/service/sefurl"
	userService "github.com/seivanov1986/gocart/internal/service/user"

	exampleAjax "github.com/seivanov1986/gocart/internal/ajax/example"

	"github.com/seivanov1986/gocart/client"
	"github.com/seivanov1986/gocart/internal/http/attribute_to_product"
	"github.com/seivanov1986/gocart/internal/widget/example"
)

type Options struct {
	database           sql_client.DataBase
	transactionManager sql_client.TransactionManager
	sessionManager     client.SessionManager
	cacheBuilder       client.CacheBuilder
	widgetManager      client.WidgetManager
	buildInWidgets     []string
}

type OptionFunc func(*Options)

func WithDatabase(database sql_client.DataBase) OptionFunc {
	return func(o *Options) {
		o.database = database
	}
}

func WithTransactionManager(trx sql_client.TransactionManager) OptionFunc {
	return func(o *Options) {
		o.transactionManager = trx
	}
}

func WithSessionManager(sessionManager client.SessionManager) OptionFunc {
	return func(o *Options) {
		o.sessionManager = sessionManager
	}
}

func WithCacheBuilder(cacheBuilder client.CacheBuilder) OptionFunc {
	return func(o *Options) {
		o.cacheBuilder = cacheBuilder
	}
}

func WithBuildInWidgets(buildInWidgets []string) OptionFunc {
	return func(o *Options) {
		o.buildInWidgets = buildInWidgets
	}
}

func WithWidgetManager(widgetManager client.WidgetManager) OptionFunc {
	return func(o *Options) {
		o.widgetManager = widgetManager
	}
}

type goCart struct {
	database           sql_client.DataBase
	transactionManager sql_client.TransactionManager
	sessionManager     client.SessionManager
	cacheBuilder       client.CacheBuilder
	cacheService       cache_service.CacheService
	widgetManager      client.WidgetManager
	buildInWidgets     []string
}

func New(opts ...OptionFunc) *goCart {
	options := Options{}
	for _, opt := range opts {
		opt(&options)
	}

	return &goCart{
		database:           options.database,
		sessionManager:     options.sessionManager,
		cacheBuilder:       options.cacheBuilder,
		widgetManager:      options.widgetManager,
		buildInWidgets:     options.buildInWidgets,
		transactionManager: options.transactionManager,
	}
}

func (g *goCart) UserHttpHandler() user.Handle {
	g.checkDatabase()
	g.checkTransactionManager()

	hub := repository.New(g.database, g.transactionManager)
	service := userService.New(hub)
	return user.New(service)
}

func (g *goCart) AuthHandler() auth.Handle {
	g.checkDatabase()
	g.checkSessionManager()
	g.checkTransactionManager()

	hub := repository.New(g.database, g.transactionManager)
	service := authService.New(hub, g.sessionManager)
	return auth.New(service)
}

func (g *goCart) FileHandler() file.Handle {
	return file.New()
}

func (g *goCart) AttributeHandler() attribute.Handle {
	g.checkDatabase()
	g.checkTransactionManager()

	hub := repository.New(g.database, g.transactionManager)
	service := attributeService.New(hub)
	return attribute.New(service)
}

func (g *goCart) AttributeToProductHandler() attribute_to_product.Handle {
	g.checkDatabase()
	g.checkTransactionManager()

	hub := repository.New(g.database, g.transactionManager)
	service := attributeToProductService.New(hub)
	return attribute_to_product.New(service)
}

func (g *goCart) CategoryHandler() category.Handle {
	g.checkDatabase()
	g.checkTransactionManager()

	hub := repository.New(g.database, g.transactionManager)
	service := categoryService.New(hub, g.transactionManager)
	return category.New(service, cache.Cache)
}

func (g *goCart) ImageToCategoryHandler() image_to_category.Handle {
	g.checkDatabase()
	g.checkTransactionManager()

	hub := repository.New(g.database, g.transactionManager)
	service := imageToCategoryService.New(hub)
	return image_to_category.New(service)
}

func (g *goCart) ImageToProductHandler() image_to_product.Handle {
	g.checkDatabase()
	g.checkTransactionManager()

	hub := repository.New(g.database, g.transactionManager)
	service := imageToProductService.New(hub)
	return image_to_product.New(service)
}

func (g *goCart) PageHandler() page.Handle {
	g.checkDatabase()
	g.checkTransactionManager()

	hub := repository.New(g.database, g.transactionManager)
	service := pageService.New(hub, g.transactionManager)
	return page.New(service, cache.Cache)
}

func (g *goCart) ProductHandler() product.Handle {
	g.checkDatabase()
	g.checkTransactionManager()

	hub := repository.New(g.database, g.transactionManager)
	service := productService.New(hub, g.transactionManager)
	return product.New(service, cache.Cache)
}

func (g *goCart) ProductToCategoryHandler() product_to_category.Handle {
	g.checkDatabase()
	g.checkTransactionManager()

	hub := repository.New(g.database, g.transactionManager)
	service := productToCategoryService.New(hub)
	return product_to_category.New(service, cache.Cache)
}

func (g *goCart) SefUrlHandler() sefurl.Handle {
	g.checkDatabase()
	g.checkTransactionManager()

	hub := repository.New(g.database, g.transactionManager)
	service := sefUrlService.New(hub)
	return sefurl.New(service)
}

func (g *goCart) CommonHandler() commonHandle.Handle {
	service := commonService.New()
	return commonHandle.New(service, g.cacheBuilder)
}

func (g *goCart) AuthMiddleware() auth2.Middleware {
	g.checkSessionManager()
	return auth2.New(g.sessionManager)
}

func (g *goCart) CommonMiddleware(serviceBasePath string) common.Middleware {
	return common.New(serviceBasePath)
}

func (g *goCart) CorsMiddleware() cors.Middleware {
	return cors.New()
}

func (g *goCart) checkDatabase() {
	if g.database == nil {
		panic("database must be an object")
	}
}

func (g *goCart) checkTransactionManager() {
	if g.database == nil {
		panic("transaction manager must be an object")
	}
}

func (g *goCart) checkSessionManager() {
	if g.sessionManager == nil {
		panic("session manager must be an object")
	}
}

func (g *goCart) CacheService() cache_service.CacheService {
	if g.cacheService != nil {
		return g.cacheService
	}

	if g.cacheBuilder != nil {
		g.cacheBuilder.RegisterWidget("example", example.New())
		service := cache_service.New(g.cacheBuilder)
		g.cacheService = service
		return service
	}

	g.checkDatabase()
	g.checkTransactionManager()

	hub := repository.New(g.database, g.transactionManager)
	if g.widgetManager == nil {
		g.widgetManager = widget_manager.New()
	}

	g.cacheBuilder = cache_builder.NewBuilder(hub, g.widgetManager)
	service := cache_service.New(g.cacheBuilder)
	g.cacheService = service
	return service
}

func (g *goCart) InitAuthHandles(router *mux.Router) {
	router.HandleFunc("/login", g.AuthHandler().Login).
		Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/logout", g.AuthHandler().Logout).
		Methods(http.MethodPost)
}

func (g *goCart) InitAdminHandles(router *mux.Router) {
	pageRouter := router.PathPrefix("/page").Subrouter()
	g.InitPageHandles(pageRouter)

	categoryRouter := router.PathPrefix("/category").Subrouter()
	g.InitCategoryHandles(categoryRouter)

	productRouter := router.PathPrefix("/product").Subrouter()
	g.InitProductHandles(productRouter)
	productToCategoryRouter := productRouter.PathPrefix("/to_category").Subrouter()
	g.InitProductToCategoryHandles(productToCategoryRouter)

	userRouter := router.PathPrefix("/user").Subrouter()
	g.InitUserHandles(userRouter)

	attributeRouter := router.PathPrefix("/attribute").Subrouter()
	g.InitAttributeHandles(attributeRouter)
	attributeToProductRouter := attributeRouter.PathPrefix("/to_product").Subrouter()
	g.InitAttributeToProductHandles(attributeToProductRouter)

	sefurlRouter := router.PathPrefix("/sefurl").Subrouter()
	g.InitSefurlHandles(sefurlRouter)

	router.HandleFunc("/ping", g.AuthHandler().Ping).
		Methods(http.MethodPost, http.MethodOptions)

	imageRouter := router.PathPrefix("/image").Subrouter()
	g.InitImageHandles(imageRouter)
	imageToProductRouter := imageRouter.PathPrefix("/to_product").Subrouter()
	g.InitImageToProduct(imageToProductRouter)
	imageToCategory := imageRouter.PathPrefix("/to_category").Subrouter()
	g.InitImageToCategory(imageToCategory)
}

func (g *goCart) InitImageToProduct(router *mux.Router) {
	handle := g.ImageToProductHandler()
	router.HandleFunc("/list", handle.List).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/create", handle.Create).Methods(http.MethodPost, http.MethodOptions)
}

func (g *goCart) InitImageToCategory(router *mux.Router) {
	handle := g.ImageToCategoryHandler()
	router.HandleFunc("/list", handle.List).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/create", handle.Create).Methods(http.MethodPost, http.MethodOptions)
}

func (g *goCart) InitImageHandles(router *mux.Router) {
	hub := repository.New(g.database, g.transactionManager)
	service := image.New(hub, g.transactionManager)
	handle := imageHandle.New(service)

	router.HandleFunc("/upload", handle.Upload).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/list", handle.List).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/create_folder", handle.CreateFolder).Methods(http.MethodPost, http.MethodOptions)
}

func (g *goCart) InitProductToCategoryHandles(router *mux.Router) {
	hub := repository.New(g.database, g.transactionManager)
	service := productToCategoryService.New(hub)
	handle := product_to_category.New(service, cache.Cache)

	router.HandleFunc("/create", handle.Create).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/delete", handle.Delete).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/list", handle.List).Methods(http.MethodPost, http.MethodOptions)
}

func (g *goCart) InitAttributeToProductHandles(router *mux.Router) {
	hub := repository.New(g.database, g.transactionManager)
	service := attributeToProductService.New(hub)
	handle := attribute_to_product.New(service)

	router.HandleFunc("/create", handle.Create).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/delete", handle.Delete).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/list", handle.List).Methods(http.MethodPost, http.MethodOptions)
}

func (g *goCart) InitSefurlHandles(router *mux.Router) {
	hub := repository.New(g.database, g.transactionManager)
	sefurlService := sefUrlService.New(hub)
	sefurlHandle := sefurl.New(sefurlService)

	router.HandleFunc("/list", sefurlHandle.List).
		Methods(http.MethodPost, http.MethodOptions)
}

func (g *goCart) InitPageHandles(router *mux.Router) {
	hub := repository.New(g.database, g.transactionManager)
	pageService := pageService.New(hub, g.transactionManager)
	pageHandle := page.New(pageService, cache.Cache)

	router.HandleFunc("/create", pageHandle.Create).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/read", pageHandle.Read).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/update", pageHandle.Update).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/delete", pageHandle.Delete).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/list", pageHandle.List).Methods(http.MethodPost, http.MethodOptions)
}

func (g *goCart) InitCategoryHandles(router *mux.Router) {
	hub := repository.New(g.database, g.transactionManager)
	categoryService := categoryService.New(hub, g.transactionManager)
	categoryHandle := category.New(categoryService, cache.Cache)

	router.HandleFunc("/create", categoryHandle.Create).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/read", categoryHandle.Read).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/update", categoryHandle.Update).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/delete", categoryHandle.Delete).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/list", categoryHandle.List).Methods(http.MethodPost, http.MethodOptions)

	router.HandleFunc("/select_list", categoryHandle.SelectList).Methods(http.MethodPost, http.MethodOptions)
}

func (g *goCart) InitProductHandles(router *mux.Router) {
	hub := repository.New(g.database, g.transactionManager)
	productService := productService.New(hub, g.transactionManager)
	productHandle := product.New(productService, cache.Cache)

	router.HandleFunc("/create", productHandle.Create).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/read", productHandle.Read).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/update", productHandle.Update).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/delete", productHandle.Delete).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/list", productHandle.List).Methods(http.MethodPost, http.MethodOptions)
}

func (g *goCart) InitUserHandles(router *mux.Router) {
	hub := repository.New(g.database, g.transactionManager)
	userService := userService.New(hub)
	userHandle := user.New(userService)

	router.HandleFunc("/create", userHandle.Create).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/read", userHandle.Read).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/update", userHandle.Update).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/delete", userHandle.Delete).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/list", userHandle.List).Methods(http.MethodPost, http.MethodOptions)
}

func (g *goCart) InitAttributeHandles(router *mux.Router) {
	hub := repository.New(g.database, g.transactionManager)
	attributeService := attributeService.New(hub)
	attributeHandle := attribute.New(attributeService)

	router.HandleFunc("/create", attributeHandle.Create).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/read", attributeHandle.Read).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/update", attributeHandle.Update).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/delete", attributeHandle.Delete).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/list", attributeHandle.List).Methods(http.MethodPost, http.MethodOptions)

	router.HandleFunc("/select_list", attributeHandle.SelectList).Methods(http.MethodPost, http.MethodOptions)
}

func (g *goCart) InitAjaxManager(manager client.AjaxManager) {
	ajaxHandler := exampleAjax.New()
	manager.RegisterPath("inexample", ajaxHandler)
}

func (g *goCart) InitYandexHandles(router *mux.Router) {
	hub := repository.New(g.database, g.transactionManager)
	service := yandex_feed.New(hub)
	handle := yandexHandle.New(service)
	router.HandleFunc("/feed", handle.Feed).Methods(http.MethodGet)
}
