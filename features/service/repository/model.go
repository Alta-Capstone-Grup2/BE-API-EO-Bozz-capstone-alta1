package repository

import (
	"capstone-alta1/features/service"

	"gorm.io/gorm"
)

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
	Partner            Partner
}

type Partner struct {
	gorm.Model
	PICPosition        string
	PICPhone           string
	PICAddress         string
	CompanyName        string
	CompanyPhone       string
	CompanyCity        string
	CompanyImageUrl    string
	CompanyAddress     string
	LinkWebsite        string
	NIBNumber          string
	NIBImageUrl        string
	SIUPNumber         string
	SIUPImageUrl       string
	Event1Name         string
	Event1ImageUrl     string
	Event2Name         string
	Event2ImageUrl     string
	Event3Name         string
	Event3ImageUrl     string
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
		AverageRating:      dataCore.AverageRating,
		ServiceImageUrl:    dataCore.ServiceImageUrl,
		City:               dataCore.City,
		PartnerID:          dataCore.PartnerID,
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
		ServiceImageUrl:    dataModel.ServiceImageUrl,
		City:               dataModel.City,
		PartnerID:          dataModel.PartnerID,
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
