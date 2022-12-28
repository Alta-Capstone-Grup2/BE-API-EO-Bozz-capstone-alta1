package mysql

import (
	"capstone-alta1/config"
	city "capstone-alta1/features/city/repository"
	partner "capstone-alta1/features/partner/repository"
	user "capstone-alta1/features/user/repository"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Cannot connect to DB")
	}

	migrateDB(db)

	return db
}

func migrateDB(db *gorm.DB) {
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&partner.Client{})
	db.AutoMigrate(&partner.Partner{})
	db.AutoMigrate(&partner.Review{})
	db.AutoMigrate(&partner.Service{})
	db.AutoMigrate(&partner.Order{})
	db.AutoMigrate(&partner.Additional{})
	db.AutoMigrate(&partner.Discussion{})
	db.AutoMigrate(&city.City{})
}
