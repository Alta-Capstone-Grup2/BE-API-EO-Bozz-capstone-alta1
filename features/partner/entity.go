package partner

import (
	"time"

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
	CreatedAt          time.Time
	UpdatedAt          time.Time
	UserID             uint
	User               UserCore
	Services           []ServiceCore
	Additionals        []AdditionalCore
	Discussions        []DiscussionCore
}

type UserCore struct {
	ID       uint
	Name     string
	Email    string
	Password string
	Role     string
}

type ServiceCore struct {
	ID                 uint
	ServiceName        string
	ServiceDescription string
	ServiceIncluded    string
	ServiceCategory    string
	ServicePrice       uint
	AverageRating      float64
	ServiceImageFile   string
	City               string
	PartnerID          uint
	ServiceAdditionals []ServiceAdditionalCore
	Review             []ReviewCore
	Discussion         []DiscussionCore
	Order              []OrderCore
}

type AdditionalCore struct {
	ID                 uint
	AdditionalName     string
	AdditionalPrice    uint
	PartnerID          uint
	ServiceAdditionals []ServiceAdditionalCore
}

type DiscussionCore struct {
	ID        uint
	Comment   string
	PartnerID uint
	ClientID  uint
	ServiceID uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ReviewCore struct {
	ID        uint
	Review    string
	Rating    float64
	OrderID   uint
	ClientID  uint
	ServiceID uint
}

type ServiceAdditionalCore struct {
	ID           uint
	AdditionalID uint
	ServiceID    uint
}

type OrderCore struct {
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
	Reviews           []ReviewCore
}

type ClientCore struct {
	ID              uint
	Gender          string
	Address         string
	City            string
	Phone           string
	ClientImageFile string
	UserID          uint
	User            UserCore
	Orders          []OrderCore
	Reviews         []ReviewCore
	Discussions     []DiscussionCore
}

type ServiceInterface interface {
	GetAll(query string) (data []Core, err error)
	GetById(id uint) (data Core, err error)
	Create(input Core, c echo.Context) error
	Update(input Core, partnerID uint, userID uint, c echo.Context) error
	Delete(partnerId, userId uint) error
	GetServices(partnerID uint) (data []ServiceCore, err error)
	GetOrders(partnerID uint) (data []OrderCore, err error)
	GetAdditionals(partnerID uint) (data []AdditionalCore, err error)
	GetPartnerRegisterData(queryCompanyName, queryPICName, queryPartnerStatus string) (data []Core, err error)
	GetPartnerRegisterDataByID(partnerID uint) (data Core, err error)
	UpdatePartnerVerifyStatus(verificationLog, verificationStatus string, partnerID uint) (err error)
	UpdateOrderConfirmStatus(orderID uint, partnerID uint) (err error)
}

type RepositoryInterface interface {
	GetAll(query string) (data []Core, err error)
	GetById(id uint) (data Core, err error)
	Create(input Core) error
	Update(input Core, partnerID uint, userID uint) error
	Delete(partnerID uint, userID uint) error
	FindUser(email string) (data Core, err error)
	GetServices(partnerID uint) (data []ServiceCore, err error)
	GetOrders(partnerID uint) (data []OrderCore, err error)
	GetAdditionals(partnerID uint) (data []AdditionalCore, err error)
	GetPartnerRegisterData(queryCompanyName, queryPICName, queryPartnerStatus string) (data []Core, err error)
	GetPartnerRegisterDataByID(partnerID uint) (data Core, err error)
	UpdatePartnerVerifyStatus(verificationLog, verificationStatus string, partnerID uint) (err error)
	UpdateOrderConfirmStatus(orderID uint, partnerID uint) (err error)
}
