package auth

import "time"

type Core struct {
	ID        uint
	Name      string
	Email     string `validate:"required,email"`
	Password  string `validate:"required"`
	Role      string
	ClientID  uint
	PartnerID uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ClientCore struct {
	ID             uint
	Gender         string
	Address        string
	City           string
	Phone          string
	ClientImageUrl string
	UserID         uint
	User           Core
	// Orders         []OrderCore // krn ga dipake di comment
}

type PartnerCore struct {
	ID                 uint
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
	User               Core
}

type ServiceInterface interface {
	Login(input Core) (data Core, token string, err error)
}

type RepositoryInterface interface {
	FindUser(email string) (result Core, err error)
	FindClient(userID uint) (result ClientCore, err error)
	FindPartner(userID uint) (result PartnerCore, err error)
}
