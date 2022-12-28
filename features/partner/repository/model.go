package repository

import (
	"capstone-alta1/features/partner"
	"time"

	"gorm.io/gorm"
)

// struct gorm model
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
	Services           []Service
	Additionals        []Additional
	Discussions        []Discussion
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
	ServiceIncluded    string
	ServiceCategory    string
	ServicePrice       uint
	AverageRating      float64
	ServiceImageFile   string
	City               string
	PartnerID          uint
	ServiceAdditionals []ServiceAdditional
	Review             []Review
	Discussion         []Discussion
	Order              []Order
}
type Additional struct {
	gorm.Model
	AdditionalName     string
	AdditionalPrice    uint
	PartnerID          uint
	ServiceAdditionals []ServiceAdditional
}

type Discussion struct {
	gorm.Model
	Comment   string
	PartnerID uint
	ClientID  uint
	ServiceID uint
}

type Review struct {
	gorm.Model
	Review    string
	Rating    float64
	OrderID   uint
	ClientID  uint
	ServiceID uint
}

type ServiceAdditional struct {
	gorm.Model
	AdditionalID uint
	ServiceID    uint
}

type Order struct {
	ID                uint
	EventName         string
	StartDate         time.Time
	EndDate           time.Time
	EventLocation     string
	EventAddress      string
	NoteForPartner    string
	ServiceName       string
	ServicePrice      uint
	GrossAmmount      uint
	PaymentMethod     string
	OrderStatus       string
	PayoutRecieptFile string
	PayoutDate        time.Time
	ServiceID         uint
	ClientID          uint
	Reviews           []Review
}

type Client struct {
	ID              uint
	Gender          string
	Address         string
	City            string
	Phone           string
	ClientImageFile string
	UserID          uint
	User            User
	Orders          []Order
	Reviews         []Review
	Discussions     []Discussion
}

// mapping

// mengubah struct core ke struct model gorm
func fromCore(dataCore partner.Core) Partner {
	PartnerGorm := Partner{
		PICPosition:        dataCore.PICPosition,
		PICPhone:           dataCore.PICPhone,
		PICAddress:         dataCore.PICAddress,
		CompanyName:        dataCore.CompanyName,
		CompanyPhone:       dataCore.CompanyPhone,
		CompanyCity:        dataCore.CompanyCity,
		CompanyImageFile:   dataCore.CompanyImageFile,
		CompanyAddress:     dataCore.CompanyAddress,
		LinkWebsite:        dataCore.LinkWebsite,
		NIBNumber:          dataCore.NIBNumber,
		NIBImageFile:       dataCore.NIBImageFile,
		SIUPNumber:         dataCore.SIUPNumber,
		SIUPImageFile:      dataCore.SIUPImageFile,
		Event1Name:         dataCore.Event1Name,
		Event1ImageFile:    dataCore.Event1ImageFile,
		Event2Name:         dataCore.Event2Name,
		Event2ImageFile:    dataCore.Event2ImageFile,
		Event3Name:         dataCore.Event3Name,
		Event3ImageFile:    dataCore.Event3ImageFile,
		BankName:           dataCore.BankName,
		BankAccountNumber:  dataCore.BankAccountNumber,
		BankAccountName:    dataCore.BankName,
		VerificationStatus: dataCore.VerificationStatus,
		VerificationLog:    dataCore.VerificationLog,
		UserID:             dataCore.UserID,
		User: User{
			Name:     dataCore.User.Name,
			Email:    dataCore.User.Email,
			Password: dataCore.User.Password,
			Role:     dataCore.User.Role,
		},
	}
	return PartnerGorm
}

// mengubah struct model gorm ke struct core
func (dataModel *Partner) toCore() partner.Core {
	return partner.Core{
		ID:                 dataModel.ID,
		PICPosition:        dataModel.PICPosition,
		PICPhone:           dataModel.PICPhone,
		PICAddress:         dataModel.PICAddress,
		CompanyName:        dataModel.CompanyName,
		CompanyPhone:       dataModel.CompanyPhone,
		CompanyCity:        dataModel.CompanyCity,
		CompanyImageFile:   dataModel.CompanyImageFile,
		CompanyAddress:     dataModel.CompanyAddress,
		LinkWebsite:        dataModel.LinkWebsite,
		NIBNumber:          dataModel.NIBNumber,
		NIBImageFile:       dataModel.NIBImageFile,
		SIUPNumber:         dataModel.SIUPNumber,
		SIUPImageFile:      dataModel.SIUPImageFile,
		Event1Name:         dataModel.Event1Name,
		Event1ImageFile:    dataModel.Event1ImageFile,
		Event2Name:         dataModel.Event2Name,
		Event2ImageFile:    dataModel.Event2ImageFile,
		Event3Name:         dataModel.Event3Name,
		Event3ImageFile:    dataModel.Event3ImageFile,
		BankName:           dataModel.BankName,
		BankAccountNumber:  dataModel.BankAccountNumber,
		BankAccountName:    dataModel.BankName,
		VerificationStatus: dataModel.VerificationStatus,
		VerificationLog:    dataModel.VerificationLog,
		UserID:             dataModel.UserID,
		User: partner.UserCore{
			ID:       dataModel.User.ID,
			Name:     dataModel.User.Name,
			Email:    dataModel.User.Email,
			Password: dataModel.User.Password,
			Role:     dataModel.User.Role,
		},
	}
}

// mengubah slice struct model gorm ke slice struct core
func toCoreList(dataModel []Partner) []partner.Core {
	var dataCore []partner.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}

// mengubah struct model gorm ke struct core
func (dataModel *Service) toCoreService() partner.ServiceCore {
	return partner.ServiceCore{
		ID:                 dataModel.ID,
		ServiceName:        dataModel.ServiceName,
		ServiceDescription: dataModel.ServiceDescription,
		ServiceIncluded:    dataModel.ServiceIncluded,
		ServiceCategory:    dataModel.ServiceCategory,
		ServicePrice:       dataModel.ServicePrice,
		AverageRating:      dataModel.AverageRating,
		ServiceImageFile:   dataModel.ServiceImageFile,
		City:               dataModel.City,
		PartnerID:          dataModel.PartnerID,
	}
}

func toCoreServiceList(dataModel []Service) []partner.ServiceCore {
	var dataCore []partner.ServiceCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCoreService())
	}
	return dataCore
}

func fromOrderCore(dataCore partner.OrderCore) Order {
	modelData := Order{
		EventName:         dataCore.EventName,
		StartDate:         dataCore.StartDate,
		EndDate:           dataCore.EndDate,
		EventLocation:     dataCore.EventLocation,
		ServiceName:       dataCore.ServiceName,
		GrossAmmount:      dataCore.GrossAmmount,
		PaymentMethod:     dataCore.PaymentMethod,
		OrderStatus:       dataCore.OrderStatus,
		PayoutRecieptFile: dataCore.PayoutRecieptFile,
		PayoutDate:        dataCore.PayoutDate,
		ServiceID:         dataCore.ServiceID,
		ClientID:          dataCore.ClientID,
	}
	return modelData
}

// mengubah struct model gorm ke struct core
func (dataModel *Order) toOrderCore() partner.OrderCore {
	return partner.OrderCore{
		ID:                dataModel.ID,
		EventName:         dataModel.EventName,
		StartDate:         dataModel.StartDate,
		EndDate:           dataModel.EndDate,
		EventLocation:     dataModel.EventLocation,
		ServiceName:       dataModel.ServiceName,
		GrossAmmount:      dataModel.GrossAmmount,
		PaymentMethod:     dataModel.PaymentMethod,
		OrderStatus:       dataModel.OrderStatus,
		PayoutRecieptFile: dataModel.PayoutRecieptFile,
		PayoutDate:        dataModel.PayoutDate,
		ServiceID:         dataModel.ServiceID,
		ClientID:          dataModel.ClientID,
	}
}

// mengubah slice struct model gorm ke slice struct core
func toOrderCoreList(dataModel []Order) []partner.OrderCore {
	var dataCore []partner.OrderCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toOrderCore())
	}
	return dataCore
}

// mengubah struct core ke struct model gorm
func fromAdditionalCore(dataCore partner.AdditionalCore) Additional {
	modelData := Additional{
		AdditionalName:  dataCore.AdditionalName,
		AdditionalPrice: dataCore.AdditionalPrice,
		PartnerID:       dataCore.PartnerID,
	}
	return modelData
}

// mengubah struct model gorm ke struct core
func (dataModel *Additional) toCore() partner.AdditionalCore {
	return partner.AdditionalCore{
		ID:              dataModel.ID,
		AdditionalName:  dataModel.AdditionalName,
		AdditionalPrice: dataModel.AdditionalPrice,
		PartnerID:       dataModel.PartnerID,
	}
}

// mengubah slice struct model gorm ke slice struct core
func toAdditionalCoreList(dataModel []Additional) []partner.AdditionalCore {
	var dataCore []partner.AdditionalCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
