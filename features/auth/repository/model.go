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

type Client struct {
	gorm.Model
	Gender          string
	Address         string
	City            string
	Phone           string
	ClientImageFile string
	UserID          uint
	User            User
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
		ID:              dataModel.ID,
		Gender:          dataModel.Gender,
		Address:         dataModel.Address,
		City:            dataModel.City,
		Phone:           dataModel.Phone,
		ClientImageFile: dataModel.ClientImageFile,
		UserID:          dataModel.User.ID,
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
		User: auth.Core{
			ID:       dataModel.User.ID,
			Name:     dataModel.User.Name,
			Email:    dataModel.User.Email,
			Password: dataModel.User.Password,
			Role:     dataModel.User.Role,
		},
	}
}
