package repository

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	EventName        string
	StartDate        time.Time
	EndDate          time.Time
	EventLocation    string
	NotesForPartner  string
	ServiceName      string
	ServicePrice     uint
	GrossAmmount     uint
	PaymentMethod    string
	OrderStatus      string
	PayoutReceiptUrl string
	PayoutDate       time.Time
	ServiceID        uint
	ClientID         uint
}
