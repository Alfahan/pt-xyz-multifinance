package pkg

import (
	"database/sql"
	"log"

	"pt-xyz-multifinance/config"

	_ "github.com/lib/pq" // Import postgres driver
)

func NewDatabase(cfg *config.Config) *sql.DB {
	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	log.Println("Database connection success")
	return db
}
