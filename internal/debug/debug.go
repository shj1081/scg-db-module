package debug

import (
	"fmt"
	"log"

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
