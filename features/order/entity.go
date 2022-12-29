package order

import "time"

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
	Service           Service
	ClientID          uint
	Client            Client
	DetailOrder       []DetailOrder
}

type DetailOrder struct {
	ID                  uint
	ServiceAdditionalID uint
	ServiceAdditional   ServiceAdditional
	AdditionalName      string
	AdditionalPrice     uint
	Qty                 uint
	DetailOrderTotal    uint
	OrderID             uint
	Order               Core
}

type ServiceAdditional struct {
	ID           uint
	AdditionalID uint
	Additional   Additional
	ServiceID    uint
	Service      Service
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
	Create(input Core, inputDetail DetailOrder) error
	GetAll(query string) (data []Core, err error)
	GetById(id uint) (data Core, dataDetail DetailOrder, err error)
}

type RepositoryInterface interface {
	Create(input Core, inputDetail DetailOrder) error
	GetAll() (data []Core, err error)
	GetAllWithSearch(query string) (data []Core, err error)
	GetById(id uint) (data Core, dataDetail DetailOrder, err error)
}
