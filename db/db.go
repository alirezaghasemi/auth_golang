package db

import (
	"auth/config"
	"auth/models"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func Init(cfg *config.Config) {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s",
		cfg.Database.Host, cfg.Database.User, cfg.Database.DBName, cfg.Database.SSLMode, cfg.Database.Password)

	var err error
	DB, err = gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB.AutoMigrate(&models.User{})
}
