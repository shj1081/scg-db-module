package debug

import (
	"database/sql"
	"fmt"
	"log"
)

// prints all available databases
func PrintDBTest(db *sql.DB) {
	log.Println("Testing database list printing...")
	rows, err := db.Query("SHOW DATABASES")
	if err != nil {
		log.Fatalf("Could not query databases: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var dbName string
		err = rows.Scan(&dbName)
		if err != nil {
			log.Fatalf("Could not scan database: %v", err)
		}
		fmt.Println(dbName)
	}
}
