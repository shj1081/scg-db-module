package debug

import (
	"fmt"
	"log"

	"scg-inouse-db-module/internal/config"

	"github.com/jmoiron/sqlx"
)

// prints all available databases
func PrintDBTest(db *sqlx.DB) {
	log.Println("Testing database list printing...")
	var databases []string
	err := db.Select(&databases, "SHOW DATABASES")
	if err != nil {
		log.Fatalf("Could not query databases: %v", err)
	}

	for _, dbName := range databases {
		fmt.Println(dbName)
	}
}

// PrintConfig prints the current application configuration
func PrintConfig(cfg *config.Config) {
	log.Printf("AppConfig: \n"+
		"  DB: \n"+
		"    DSN: %s\n"+
		"    MaxOpenConns: %d\n"+
		"    MaxIdleConns: %d\n"+
		"    ConnMaxLifetime: %s\n"+
		"  Server: \n"+
		"    Port: %s\n"+
		"    Environment: %s\n"+
		"  Auth: \n"+
		"    ProxyURL: %s\n",
		cfg.DB.DSN,
		cfg.DB.MaxOpenConns,
		cfg.DB.MaxIdleConns,
		cfg.DB.ConnMaxLifetime,
		cfg.Server.Port,
		cfg.Server.Environment,
		cfg.Auth.ProxyURL)
}
