package order

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID                    uint
	EventName             string
	StartDate             time.Time
	StartTime             time.Duration
	EndDate               time.Time
	EndTime               time.Duration
	EventLocation         string
	EventAddress          string
	NotesForPartner       string
	ServiceName           string
	ServicePrice          uint
	GrossAmmount          uint
	PaymentMethod         string
	OrderStatus           string
	PayoutRecieptFile     string
	PayoutDate            time.Time
	MidtransTransactionID string
	MidtransToken         string
	MidtransLink          string
	MidtransVaNumber      string
	MidtransExpiredTime   string
	ServiceID             uint
	ClientID              uint
	DetailOrder           []DetailOrder
}

type DetailOrder struct {
	ID                  uint
	AdditionalName      string
	AdditionalPrice     uint
	Qty                 uint
	DetailOrderTotal    uint
	ServiceAdditionalID uint
	OrderID             uint
}

type ServiceAdditional struct {
	ID           uint
	AdditionalID uint
	ServiceID    uint
	DetailOrders []DetailOrder
}

type Additional struct {
	ID                uint
	AdditionalName    string
	AdditionalPrice   uint
	PartnerID         uint
	ServiceAdditional []ServiceAdditional
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
	Core            []Core
}

type User struct {
	ID       uint
	Name     string
	Email    string
	Password string
	Role     string
}

type Service struct {
	ID                 uint
	ServiceName        string
	ServiceDescription string
	ServiceCategory    string
	ServiceIncluded    string
	ServicePrice       uint
	AverageRating      float64
	ServiceImageFile   string
	City               string
	PartnerID          uint
	ServiceAdditional  []ServiceAdditional
}

type Partner struct {
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
	User               User
	Services           []Service
	Additionals        []Additional
}

type OrderJoinPartner struct {
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
	PayoutDate        time.Time `gorm:"default:null"`
	ServiceID         uint
	ClientID          uint
	PartnerID         uint
	CompanyName       string
	BankName          string
	BankAccountNumber string
	BankAccountName   string
}

type ServiceInterface interface {
	Create(input Core, inputDetail []DetailOrder) (data Core, err error)
	GetAll(query string) (data []OrderJoinPartner, err error)
	GetById(id uint) (data Core, dataDetail []DetailOrder, err error)
	UpdateStatusCancel(input Core, id uint) error
	UpdateStatusPayout(id uint, c echo.Context) error
	UpdateMidtrans(input Core) error
}

type RepositoryInterface interface {
	Create(input Core, inputDetail []DetailOrder) (data Core, err error)
	GetAll(query string) (data []OrderJoinPartner, err error)
	GetAllWithSearch(query string) (data []Core, err error)
	GetById(id uint) (data Core, dataDetail []DetailOrder, err error)
	UpdateStatusCancel(input Core, id uint) error
	UpdateStatusPayout(input Core, id uint) error
	UpdateMidtrans(input Core) error
	GetServiceByID(serviceID uint) (data Service, err error)
}
