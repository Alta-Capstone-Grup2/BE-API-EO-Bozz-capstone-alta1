package order

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID                uint
	EventName         string
	StartDate         time.Time
	StartTime         time.Duration
	EndDate           time.Time
	EndTime           time.Duration
	EventLocation     string
	EventAddress      string
	NotesForPartner   string
	ServiceName       string
	ServicePrice      uint
	GrossAmmount      uint
	PaymentMethod     string
	OrderStatus       string
	PayoutRecieptFile string
	PayoutDate        time.Time
	ServiceID         uint
	ClientID          uint
	DetailOrder       []DetailOrder
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
	ServicePrice       uint
	AverageRating      float64
	ServiceImageFile   string
	City               string
	PartnerID          uint
	ServiceAdditional  []ServiceAdditional
}

type ServiceInterface interface {
	Create(input Core, inputDetail []DetailOrder) error
	GetAll(query string) (data []Core, err error)
	GetById(id uint) (data Core, dataDetail DetailOrder, err error)
	UpdateStatusCancel(input Core, id uint) error
	UpdateStatusPayout(id uint, c echo.Context) error
}

type RepositoryInterface interface {
	Create(input Core, inputDetail []DetailOrder) error
	GetAll() (data []Core, err error)
	GetAllWithSearch(query string) (data []Core, err error)
	GetById(id uint) (data Core, dataDetail DetailOrder, err error)
	UpdateStatusCancel(input Core, id uint) error
	UpdateStatusPayout(input Core, id uint) error
}
