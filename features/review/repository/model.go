package repository

import (
	_review "capstone-alta1/features/review"
	"time"

	"gorm.io/gorm"
)

// struct gorm model
type Review struct {
	gorm.Model
	Review    string
	Rating    float64
	OrderID   uint
	ClientID  uint
	ServiceID uint
	Order     Order
	Client    Client
	Service   Service
}

type Client struct {
	gorm.Model
	Gender         string
	Address        string
	City           string
	Phone          string
	ClientImageUrl string
	UserID         uint
	User           User
	Orders         []Order
}

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Role     string
}

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

type Service struct {
	gorm.Model
	ServiceName        string
	ServiceDescription string
	ServiceCategory    string
	ServicePrice       uint
	AverageRating      float64
	ServiceImageUrl    string
	City               string
	PartnerID          uint
}

// mapping

// mengubah struct core ke struct model gorm
func fromCore(dataCore _review.Core) Review {
	modelData := Review{
		Review:    dataCore.Review,
		Rating:    dataCore.Rating,
		OrderID:   dataCore.OrderID,
		ClientID:  dataCore.ClientID,
		ServiceID: dataCore.ServiceID,
	}
	return modelData
}

// mengubah struct model gorm ke struct core
func (dataModel *Review) toCore() _review.Core {
	return _review.Core{
		ID:        dataModel.ID,
		Review:    dataModel.Review,
		Rating:    dataModel.Rating,
		OrderID:   dataModel.OrderID,
		ClientID:  dataModel.ClientID,
		ServiceID: dataModel.ServiceID,
		CreatedAt: dataModel.CreatedAt,
		UpdatedAt: dataModel.UpdatedAt,
	}
}

// mengubah slice struct model gorm ke slice struct core
func toCoreList(dataModel []Review) []_review.Core {
	var dataCore []_review.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
