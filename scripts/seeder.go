package scripts

import (
	"log"
	"pt-xyz-multifinance/config"
	"pt-xyz-multifinance/internal/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func seeder() {
	cfg := config.New()
	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("cannot connect database: %v", err)
	}

	// Seeder contoh untuk tabel User
	users := []domain.User{
		{Name: "Admin", Email: "admin@mail.com", Password: "hashedpassword"},
		{Name: "User", Email: "user@mail.com", Password: "hashedpassword"},
	}
	for _, user := range users {
		if err := db.FirstOrCreate(&user, domain.User{Email: user.Email}).Error; err != nil {
			log.Fatalf("seeding user failed: %v", err)
		}
	}

	// Tambahkan seeder entity lain sesuai kebutuhan

	log.Println("Seeding success")
}
