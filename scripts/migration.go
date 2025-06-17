package scripts

import (
	"log"
	"pt-xyz-multifinance/config"
	"pt-xyz-multifinance/internal/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Migrate() {
	cfg := config.New()
	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("cannot connect database: %v", err)
	}

	// Daftarkan semua entity (struct) yang ingin di-migrate
	err = db.AutoMigrate(
		&domain.User{},
		// Tambah entity lain di sini (misal: &domain.Product{}, &domain.Order{})
	)
	if err != nil {
		log.Fatalf("migration failed: %v", err)
	}
	log.Println("Migration success")
}
