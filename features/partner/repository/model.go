package repository

import (
	"capstone-alta1/features/partner"

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
}

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Role     string
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
