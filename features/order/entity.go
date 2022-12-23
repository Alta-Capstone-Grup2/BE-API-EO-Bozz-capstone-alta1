package order

import "time"

type Core struct {
	ID            uint
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
