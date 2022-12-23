package repository

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	EventName     string
	StartDate     time.Time
	EndDate       time.Time
	EventLocation string
	ServiceName   string
	GrossAmmount  int
	OrderStatus   string
	ServiceID     uint
	UserID        uint
}
