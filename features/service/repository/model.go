package repository

import (
	"capstone-alta1/features/service"
	"time"

	"gorm.io/gorm"
)

type Service struct {
	gorm.Model
	ServiceName        string
	ServiceDescription string
	ServiceCategory    string
	ServicePrice       uint
	AverageRating      float64
	ServiceImageFile   string
	City               string
	PartnerID          uint
	Partner            Partner
	Additional         []Additional
	Review             []Review
	Discussion         []Discussion
	Order              []Order
}

type Discussion struct {
	gorm.Model
	Comment   string
	PartnerID uint
	ClientID  uint
	ServiceID uint
	Service   Service
}

type Review struct {
	gorm.Model
	Review    string
	Rating    float64
	OrderID   uint
	ClientID  uint
	ServiceID uint
	Service   Service
}

type ServiceAdditional struct {
	gorm.Model
	AdditionalID uint
	Additional   Additional
	ServiceID    uint
	Service      Service
}

type Additional struct {
	gorm.Model
	AdditionalName  string
	AdditionalPrice uint
	PartnerID       uint
	ServiceID       uint
	Service         Service
}

type Order struct {
	gorm.Model
	EventName          string
	ServiceName        string
	StartDate          time.Time
	EndDate            time.Time
	AvailabilityStatus string
	ServiceID          uint
	Service            Service
}

type Partner struct {
	gorm.Model
	PICPosition        string
	PICPhone           string
	PICAddress         string
	CompanyName        string
	CompanyPhone       string
	CompanyCity        string
	CompanyImageFile   string
	CompanyAddress     string
	LinkWebsite        string
	NIBNumber          string
	NIBImageFile       string
	SIUPNumber         string
	SIUPImageFile      string
	Event1Name         string
	Event1ImageFile    string
	Event2Name         string
	Event2ImageFile    string
	Event3Name         string
	Event3ImageFile    string
	BankName           string
	BankAccountNumber  string
	BankAccountName    string
	VerificationStatus string
	VerificationLog    string
	UserID             uint
	User               User
}

type User struct {
	gorm.Model
	Name     string `validate:"required"`
	Email    string `validate:"required,email,unique"`
	Password string `validate:"required"`
	Role     string
}

// mapping

// mengubah struct core ke struct model gorm
func fromCore(dataCore service.Core) Service {
	modelData := Service{
		ServiceName:        dataCore.ServiceName,
		ServiceDescription: dataCore.ServiceDescription,
		ServiceCategory:    dataCore.ServiceCategory,
		ServicePrice:       dataCore.ServicePrice,
		ServiceImageFile:   dataCore.ServiceImageFile,
		City:               dataCore.City,
		PartnerID:          dataCore.PartnerID,
	}
	return modelData
}

func fromCoreServiceAdditional(dataCore service.ServiceAdditional) ServiceAdditional {
	modelData := ServiceAdditional{
		ServiceID:    dataCore.ServiceID,
		AdditionalID: dataCore.AdditionalID,
	}
	return modelData
}

// mengubah struct model gorm ke struct core
func (dataModel *Service) toCore() service.Core {
	return service.Core{
		ID:                 dataModel.ID,
		ServiceName:        dataModel.ServiceName,
		ServiceDescription: dataModel.ServiceDescription,
		ServiceCategory:    dataModel.ServiceCategory,
		ServicePrice:       dataModel.ServicePrice,
		AverageRating:      dataModel.AverageRating,
		ServiceImageFile:   dataModel.ServiceImageFile,
		City:               dataModel.City,
		PartnerID:          dataModel.PartnerID,
	}
}

func (dataModel *Additional) toCoreAdditional() service.Additional {
	return service.Additional{
		ID:              dataModel.ID,
		AdditionalName:  dataModel.AdditionalName,
		AdditionalPrice: dataModel.AdditionalPrice,
		PartnerID:       dataModel.PartnerID,
		ServiceID:       dataModel.ServiceID,
	}
}

func (dataModel *Review) toCoreReview() service.Review {
	return service.Review{
		ID:        dataModel.ID,
		Review:    dataModel.Review,
		Rating:    dataModel.Rating,
		OrderID:   dataModel.OrderID,
		ClientID:  dataModel.ClientID,
		ServiceID: dataModel.ServiceID,
	}
}

func (dataModel *Discussion) toCoreDiscussion() service.Discussion {
	return service.Discussion{
		ID:        dataModel.ID,
		Comment:   dataModel.Comment,
		CreatedAt: dataModel.CreatedAt,
		PartnerID: dataModel.PartnerID,
		ClientID:  dataModel.ClientID,
		ServiceID: dataModel.ServiceID,
	}
}

func (dataModel *Order) toCoreAvailable(serviceName string, startDate, endDate time.Time, Available string) service.Order {
	return service.Order{
		ServiceName:        serviceName,
		StartDate:          startDate,
		EndDate:            endDate,
		AvailabilityStatus: Available,
	}
}

func (dataModel *Order) toCoreNotAvailable(serviceName string, startDate, endDate time.Time, NotAvailable string) service.Order {
	return service.Order{
		ServiceName:        serviceName,
		StartDate:          startDate,
		EndDate:            endDate,
		AvailabilityStatus: NotAvailable,
	}
}

// mengubah slice struct model gorm ke slice struct core
func toCoreList(dataModel []Service) []service.Core {
	var dataCore []service.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}

func toCoreListAdditional(dataModel []Additional) []service.Additional {
	var dataCore []service.Additional
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCoreAdditional())
	}
	return dataCore
}

func toCoreListReview(dataModel []Review) []service.Review {
	var dataCore []service.Review
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCoreReview())
	}
	return dataCore
}

func toCoreListDiscussion(dataModel []Discussion) []service.Discussion {
	var dataCore []service.Discussion
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCoreDiscussion())
	}
	return dataCore
}
