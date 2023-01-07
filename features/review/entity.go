package review

import (
	"time"
)

type Core struct {
	ID        uint
	Review    string  `valiidate:"required"`
	Rating    float64 `valiidate:"required"`
	OrderID   uint    `valiidate:"required"`
	ClientID  uint    `valiidate:"required"`
	ServiceID uint    `valiidate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
type Client struct {
	ID             uint
	Gender         string
	Address        string
	City           string
	Phone          string
	ClientImageUrl string
	UserID         uint
	User           User
	Orders         []Order
}

type User struct {
	ID       uint
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
	Role     string
}

type Order struct {
	ID                uint
	EventName         string
	StartDate         time.Time
	EndDate           time.Time
	EventLocation     string
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
}

type ServiceInterface interface {
	GetAll(query string) (data []Core, err error)
	Create(input Core) error
	GetById(id uint) (data Core, err error)
	Update(input Core, id uint) error
	Delete(id uint) error
}

type RepositoryInterface interface {
	GetAll() (data []Core, err error)
	Create(input Core) error
	GetById(id uint) (data Core, err error)
	Update(input Core, id uint) error
	Delete(id uint) error
}
