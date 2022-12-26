package repository

import (
	"capstone-alta1/features/additional"

	"gorm.io/gorm"
)

// struct gorm model
type Additional struct {
	gorm.Model
	AdditionalName  string
	AdditionalPrice int
	PartnerID       uint
	Partner         Partner
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
func fromCore(dataCore additional.Core) Additional {
	modelData := Additional{
		AdditionalName:  dataCore.AdditionalName,
		AdditionalPrice: dataCore.AdditionalPrice,
		PartnerID:       dataCore.PartnerID,
	}
	return modelData
}

// mengubah struct model gorm ke struct core
func (dataModel *Additional) toCore() additional.Core {
	return additional.Core{
		ID:              dataModel.ID,
		AdditionalName:  dataModel.AdditionalName,
		AdditionalPrice: dataModel.AdditionalPrice,
		PartnerID:       dataModel.PartnerID,
	}
}

// mengubah slice struct model gorm ke slice struct core
func toCoreList(dataModel []Additional) []additional.Core {
	var dataCore []additional.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
