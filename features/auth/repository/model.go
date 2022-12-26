package repository

import (
	"capstone-alta1/features/auth"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint
	Name      string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
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

type Client struct {
	gorm.Model
	Gender         string
	Address        string
	City           string
	Phone          string
	ClientImageUrl string
	UserID         uint
	User           User
	// Orders         []Order // krn ga di pake di comment
}

//DTO

func (dataModel User) toCore() auth.Core {
	return auth.Core{
		ID:        dataModel.ID,
		Name:      dataModel.Name,
		Email:     dataModel.Email,
		Password:  dataModel.Password,
		Role:      dataModel.Role,
		CreatedAt: dataModel.CreatedAt,
		UpdatedAt: dataModel.UpdatedAt,
	}
}

// mengubah struct model gorm ke struct core
func (dataModel *Client) toCore() auth.ClientCore {
	return auth.ClientCore{
		ID:             dataModel.ID,
		Gender:         dataModel.Gender,
		Address:        dataModel.Address,
		City:           dataModel.City,
		Phone:          dataModel.Phone,
		ClientImageUrl: dataModel.ClientImageUrl,
		UserID:         dataModel.User.ID,
		User: auth.Core{
			ID:       dataModel.User.ID,
			Name:     dataModel.User.Name,
			Email:    dataModel.User.Email,
			Password: dataModel.User.Password,
			Role:     dataModel.User.Role,
		},
	}
}

// mengubah struct model gorm ke struct core
func (dataModel *Partner) toCore() auth.PartnerCore {
	return auth.PartnerCore{
		ID:                 dataModel.ID,
		PICPosition:        dataModel.PICPosition,
		PICPhone:           dataModel.PICPhone,
		PICAddress:         dataModel.PICAddress,
		CompanyName:        dataModel.CompanyName,
		CompanyPhone:       dataModel.CompanyPhone,
		CompanyCity:        dataModel.CompanyCity,
		CompanyImageUrl:    dataModel.CompanyImageUrl,
		CompanyAddress:     dataModel.CompanyAddress,
		LinkWebsite:        dataModel.LinkWebsite,
		NIBNumber:          dataModel.NIBNumber,
		NIBImageUrl:        dataModel.NIBImageUrl,
		SIUPNumber:         dataModel.SIUPNumber,
		SIUPImageUrl:       dataModel.SIUPImageUrl,
		Event1Name:         dataModel.Event1Name,
		Event1ImageUrl:     dataModel.Event1ImageUrl,
		Event2Name:         dataModel.Event2Name,
		Event2ImageUrl:     dataModel.Event2ImageUrl,
		Event3Name:         dataModel.Event3Name,
		Event3ImageUrl:     dataModel.Event3ImageUrl,
		BankName:           dataModel.BankName,
		BankAccountNumber:  dataModel.BankAccountNumber,
		BankAccountName:    dataModel.BankName,
		VerificationStatus: dataModel.VerificationStatus,
		VerificationLog:    dataModel.VerificationLog,
		UserID:             dataModel.UserID,
		User: auth.Core{
			ID:       dataModel.User.ID,
			Name:     dataModel.User.Name,
			Email:    dataModel.User.Email,
			Password: dataModel.User.Password,
			Role:     dataModel.User.Role,
		},
	}
}
