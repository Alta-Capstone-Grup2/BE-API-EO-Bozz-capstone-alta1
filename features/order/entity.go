package order

import "time"

type Core struct {
	ID               uint
	EventName        string
	StartDate        time.Time
	StartTime        time.Duration
	EndDate          time.Time
	EndTime          time.Duration
	EventLocation    string
	EventAddress     string
	NotesForPartner  string
	ServiceName      string
	ServicePrice     uint
	GrossAmmount     uint
	PaymentMethod    string
	OrderStatus      string
	PayoutRecieptUrl string
	PayoutDate       time.Time
	ServiceID        uint
	Service          Service
	ClientID         uint
	Client           Client
	AdditionalID     uint
	Additional       Additional
	Qty              uint
	DetailOrderTotal string
}

type Additional struct {
	ID              uint
	AdditionalName  string
	AdditionalPrice uint
	PartnerID       uint
}

type Client struct {
	ID             uint
	Gender         string
	Address        string
	City           string
	Phone          string
	ClientImageUrl string
	UserID         uint
}

type Service struct {
	ID                 uint
	ServiceName        string
	ServiceDescription string
	ServiceCategory    string
	ServicePrice       uint
	AverageRating      float64
	ServiceImageUrl    string
	City               string
	PartnerID          uint
}

type ServiceInterface interface {
	GetAll() (data []Core, err error)
	GetById(id uint) (data Core, err error)
	Create(input Core) error
	Update(input Core, id uint) error
	Delete(id uint) error
}

type RepositoryInterface interface {
	GetAll() (data []Core, err error)
	GetById(id uint) (data Core, err error)
	Create(input Core) error
	Update(input Core, id uint) error
	Delete(id uint) error
}
