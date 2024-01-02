package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"

	"github.com/seivanov1986/gocart"
	header_widget "github.com/seivanov1986/gocart-ecommerce/internal/widget/header"
	"github.com/seivanov1986/gocart/external/ajax_manager"
	"github.com/seivanov1986/gocart/external/cache"
	"github.com/seivanov1986/gocart/external/observer"
	"github.com/seivanov1986/gocart/external/widget_manager"
	"github.com/seivanov1986/sql_client"

	ajaxExample "github.com/seivanov1986/gocart-ecommerce/internal/ajax/example"
	"github.com/seivanov1986/gocart-ecommerce/internal/widget/example"
	"github.com/seivanov1986/gocart-ecommerce/migrations"
	"github.com/seivanov1986/gocart-ecommerce/pkg/session_manager"

	"github.com/seivanov1986/sql_client/sqlite"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	serviceBasePath := os.Getenv("SERVICE_BASE_PATH")
	if serviceBasePath == "" {
		panic("service base path is not found")
	}

	sqliteDBClient, err := sqlite.NewClient("./test.db")
	if err != nil {
		log.Fatalf("db config validate error: %v", err)
	}
	err = sqliteDBClient.RunMigrations(log.Default(), migrations.MigrationFiles)
	if err != nil {
		log.Fatalf("db config validate error: %v", err)
	}

	transactionManager := sql_client.NewTransactionManager(sqliteDBClient)

	router := mux.NewRouter()
	ctx := context.Background()
	ctx = observer.SetServiceBasePath(ctx, serviceBasePath)

	widgetManger := widget_manager.New()
	ajaxManager := ajax_manager.New()
	ajaxManager.RegisterPath("outexample", ajaxExample.New())
	widgetManger.Register("exampleout", example.New())
	widgetManger.Register("header", header_widget.New())

	sessionManager := session_manager.New()

	goLib := gocart.New(
		gocart.WithDatabase(sqliteDBClient),
		gocart.WithSessionManager(sessionManager),
		gocart.WithTransactionManager(transactionManager),
		gocart.WithWidgetManager(widgetManger),
	)

	goLib.InitAjaxManager(ajaxManager)

	cacheService := goLib.CacheService()
	cache.Cache.SetProcess(func() {
		cacheService.Make(ctx)
	})

	cache.Cache.AddEvent()

	corsMiddleware := goLib.CorsMiddleware()
	commonMiddleware := goLib.CommonMiddleware(serviceBasePath)
	authMiddleware := goLib.AuthMiddleware()

	commonHandle := goLib.CommonHandler()
	fileHandle := goLib.FileHandler()

	router.Use(commonMiddleware.Handle, corsMiddleware.Handle)
	notFoundHandler := commonMiddleware.Wrapper(commonHandle.Process)
	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	goLib.InitAuthHandles(router)

	adminRouter := router.PathPrefix("/admin").Subrouter()
	adminRouter.Use(authMiddleware.Handle)
	goLib.InitAdminHandles(adminRouter)

	router.HandleFunc("/ajax", ajaxManager.Handler).
		Methods(http.MethodPost, http.MethodOptions)

	router.PathPrefix("/admin").HandlerFunc(fileHandle.AdminStatic).
		Methods(http.MethodGet)

	router.PathPrefix("/static").HandlerFunc(fileHandle.Static).
		Methods(http.MethodGet)
	router.PathPrefix("/dynamic").HandlerFunc(fileHandle.Dynamic).
		Methods(http.MethodGet)

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("ready")
	log.Fatal(srv.ListenAndServe())
}
