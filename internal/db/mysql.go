package db

import (
	"scg-inouse-db-module/internal/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

// initializes the global DB connection.
func InitDB(dsn string) error {
	var err error
	DB, err = sqlx.Open("mysql", dsn)
	if err != nil {
		return err
	}

	// set connection pool (adjust as needed)
	DB.SetMaxOpenConns(config.AppConfig.DB.MaxOpenConns)
	DB.SetMaxIdleConns(config.AppConfig.DB.MaxIdleConns)
	DB.SetConnMaxLifetime(config.AppConfig.DB.ConnMaxLifetime)

	return DB.Ping()
}

// closes the DB connection.
func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
