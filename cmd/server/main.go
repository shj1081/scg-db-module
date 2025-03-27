package main

import (
	"log"
	"net/http"
	"scg-inouse-db-module/internal/config"
	"scg-inouse-db-module/internal/db"
	"scg-inouse-db-module/internal/debug"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

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
	debug.PrintDBTest(db.DB)

	// chi router setup
	r := chi.NewRouter()

	// setup middlewares
	r.Use(middleware.Logger)    // reuqest logging
	r.Use(middleware.Recoverer) // panic handling

	// // API 라우트 그룹
	// r.Route("/api", func(r chi.Router) {
	// 	// 데이터베이스 관련 엔드포인트
	// 	r.Get("/databases", handlers.ListDatabasesHandler)
	// 	r.Post("/databases", handlers.CreateDatabaseHandler)
	// 	r.Delete("/databases/{databaseName}", handlers.DropDatabaseHandler)

	// 	// 테이블 관련 엔드포인트
	// 	r.Get("/databases/{databaseName}/tables", handlers.ListTablesHandler)
	// 	r.Post("/databases/{databaseName}/tables", handlers.CreateTableHandler)
	// 	r.Delete("/databases/{databaseName}/tables/{tableName}", handlers.DropTableHandler)
	// 	r.Get("/databases/{databaseName}/tables/{tableName}/schema", handlers.GetTableSchemaHandler)

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

	log.Println("Server running on ", config.AppConfig.Server.Port)
	if err := http.ListenAndServe(":"+config.AppConfig.Server.Port, r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
