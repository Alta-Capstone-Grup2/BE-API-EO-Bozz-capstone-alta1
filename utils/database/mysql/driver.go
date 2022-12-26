package mysql

import (
	"capstone-alta1/config"
	additional "capstone-alta1/features/additional/repository"
	city "capstone-alta1/features/city/repository"
	client "capstone-alta1/features/client/repository"
	discussion "capstone-alta1/features/discussion/repository"
	order "capstone-alta1/features/order/repository"
	partner "capstone-alta1/features/partner/repository"
	review "capstone-alta1/features/review/repository"
	service "capstone-alta1/features/service/repository"
	user "capstone-alta1/features/user/repository"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to DB")
	}

	migrateDB(db)

	return db
}

func migrateDB(db *gorm.DB) {
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&client.Client{})
	db.AutoMigrate(&partner.Partner{})
	db.AutoMigrate(&review.Review{})
	db.AutoMigrate(&service.Service{})
	db.AutoMigrate(&order.Order{})
	db.AutoMigrate(&additional.Additional{})
	db.AutoMigrate(&discussion.Discussion{})
	db.AutoMigrate(&city.City{})
}
