package main

import (
	"log"
	"net/http"
	"scg-inouse-db-module/internal/config"
	"scg-inouse-db-module/internal/db"
	"scg-inouse-db-module/internal/debug"
	"scg-inouse-db-module/internal/handlers"

	_ "scg-inouse-db-module/docs" // swagger auto generated files

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title SCG Inouse DB Module API
// @version 1.0
// @description This is a database management service API
// @host localhost:8080
// @BasePath /api
func main() {

	// load env variables and config
	config.LoadConfig()

	// initialize db connection
	err := db.InitDB(config.AppConfig.DB.DSN)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.CloseDB()

	// TODO: remove this
	debug.PrintConfig(config.AppConfig)
	debug.PrintDBTest(db.DB)

	// chi router setup
	r := chi.NewRouter()

	// setup middlewares
	r.Use(middleware.Logger)    // reuqest logging
	r.Use(middleware.Recoverer) // panic handling

	// API routes group
	r.Route("/api", func(r chi.Router) {

		// database related endpoints
		r.Get("/databases", handlers.ListDatabasesHandler)
		r.Post("/databases", handlers.CreateDatabaseHandler)
		r.Delete("/databases/{databaseName}", handlers.DropDatabaseHandler)

		// table related endpoints
		r.Get("/databases/{databaseName}/tables", handlers.ListTablesHandler)
		r.Get("/databases/{databaseName}/tables/{tableName}", handlers.GetTableHandler)
		r.Get("/databases/{databaseName}/tables/{tableName}/schema", handlers.GetTableSchemaHandler)

		// TODO: should these endpoints be implemented in db module?
		// 	r.Post("/databases/{databaseName}/tables", handlers.CreateTableHandler)
		// 	r.Delete("/databases/{databaseName}/tables/{tableName}", handlers.DropTableHandler)

		// query related endpoints
		r.Post("/databases/{databaseName}/query", handlers.RawQueryHandler)
	})

	// 	// 쿼리 관련 엔드포인트
	// 	r.Post("/databases/{databaseName}/query", handlers.RawQueryHandler)
	// 	r.Post("/databases/{databaseName}/concurrent-query", handlers.ConcurrentQueryHandler)

	// 	// 쿼리 저장 관련 엔드포인트
	// 	r.Post("/queries", handlers.SaveQueryHandler)
	// 	r.Get("/queries", handlers.ListQueriesHandler)
	// 	r.Put("/queries/{queryId}", handlers.UpdateQueryHandler)
	// 	r.Delete("/queries/{queryId}", handlers.DeleteQueryHandler)

	// 	// 히스토리 및 로그 관련 엔드포인트
	// 	r.Get("/query-history", handlers.ListQueryHistoryHandler)
	// 	r.Get("/logs", handlers.ListLogsHandler)

	// 	// XLSX 내보내기
	// 	r.Post("/export/xlsx", handlers.ExportXLSXHandler)
	// })

	// // 정적 파일 서빙
	// fileServer := http.FileServer(http.Dir("./downloads/"))
	// r.Handle("/downloads/*", http.StripPrefix("/downloads/", fileServer))

	// Swagger UI endpoint
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // swagger JSON 파일의 URL
	))

	log.Println("Server running on ", config.AppConfig.Server.Port)
	if err := http.ListenAndServe(":"+config.AppConfig.Server.Port, r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
