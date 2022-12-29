package partner

import (
	"github.com/labstack/echo/v4"
)

type Core struct {
	ID                 uint
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
	User               UserCore
}

type UserCore struct {
	ID       uint
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
	Role     string
}

type ServiceInterface interface {
	GetAll(query string) (data []Core, err error)
	GetById(id uint) (data Core, err error)
	Create(input Core, c echo.Context) error
	Update(input Core, partnerID uint, userID uint, c echo.Context) error
	Delete(partnerID uint, userID uint) error
}

type RepositoryInterface interface {
	GetAll() (data []Core, err error)
	GetAllWithSearch(query string) (data []Core, err error)
	GetById(id uint) (data Core, err error)
	Create(input Core) error
	Update(input Core, partnerID uint, userID uint) error
	Delete(partnerID uint, userID uint) error
	FindUser(email string) (data Core, err error)
}
