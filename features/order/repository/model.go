package repository

import (
	"capstone-alta1/features/order"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	EventName        string
	StartDate        time.Time
	StartTime        time.Duration
	EndDate          time.Time
	EndTime          time.Duration
	EventLocation    string
	EventAddress     string
	NotesForPartner  string
	ServiceName      string
	ServicePrice     uint
	GrossAmmount     uint
	PaymentMethod    string
	OrderStatus      string
	PayoutRecieptUrl string
	PayoutDate       time.Time
	ServiceID        uint
	Service          Service
	ClientID         uint
	Client           Client
	AdditionalID     uint
	Additional       Additional
	Qty              uint
	DetailOrderTotal string
}

type Additional struct {
	gorm.Model
	AdditionalName  string
	AdditionalPrice int
	PartnerID       uint
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
}

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Role     string
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
func fromCore(dataCore order.Core) Order {
	modelData := Order{
		EventName:        dataCore.EventName,
		StartDate:        dataCore.StartDate,
		StartTime:        dataCore.StartTime,
		EndDate:          dataCore.EndDate,
		EndTime:          dataCore.EndTime,
		EventLocation:    dataCore.EventLocation,
		EventAddress:     dataCore.EventAddress,
		NotesForPartner:  dataCore.NotesForPartner,
		ServiceName:      dataCore.ServiceName,
		ServicePrice:     dataCore.ServicePrice,
		GrossAmmount:     dataCore.GrossAmmount,
		PaymentMethod:    dataCore.PaymentMethod,
		OrderStatus:      dataCore.OrderStatus,
		PayoutRecieptUrl: dataCore.PayoutRecieptUrl,
		PayoutDate:       dataCore.PayoutDate,
		ServiceID:        dataCore.ServiceID,
		ClientID:         dataCore.ClientID,
		AdditionalID:     dataCore.AdditionalID,
		Qty:              dataCore.Qty,
		DetailOrderTotal: dataCore.DetailOrderTotal,
	}
	return modelData
}

// mengubah struct model gorm ke struct core
func (dataModel *Order) toCore() order.Core {
	return order.Core{
		ID:               dataModel.ID,
		EventName:        dataModel.EventName,
		StartDate:        dataModel.StartDate,
		StartTime:        dataModel.StartTime,
		EndDate:          dataModel.EndDate,
		EndTime:          dataModel.EndTime,
		EventLocation:    dataModel.EventLocation,
		EventAddress:     dataModel.EventAddress,
		NotesForPartner:  dataModel.NotesForPartner,
		ServiceName:      dataModel.ServiceName,
		ServicePrice:     dataModel.ServicePrice,
		GrossAmmount:     dataModel.GrossAmmount,
		PaymentMethod:    dataModel.PaymentMethod,
		OrderStatus:      dataModel.OrderStatus,
		PayoutRecieptUrl: dataModel.PayoutRecieptUrl,
		PayoutDate:       dataModel.PayoutDate,
		ServiceID:        dataModel.ServiceID,
		ClientID:         dataModel.ClientID,
		AdditionalID:     dataModel.AdditionalID,
		Qty:              dataModel.Qty,
		DetailOrderTotal: dataModel.DetailOrderTotal,
	}
}

// mengubah slice struct model gorm ke slice struct core
func toCoreList(dataModel []Order) []order.Core {
	var dataCore []order.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
