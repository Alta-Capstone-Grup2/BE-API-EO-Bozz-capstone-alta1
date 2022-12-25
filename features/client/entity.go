package client

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID             uint
	Gender         string
	Address        string
	City           string
	Phone          string
	ClientImageUrl string
	UserID         uint
	User           UserCore
	Orders         []OrderCore
}

type UserCore struct {
	ID       uint
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
	Role     string
}

type OrderCore struct {
	ID               uint
	EventName        string
	StartDate        time.Time
	EndDate          time.Time
	EventLocation    string
	NotesForPartner  string
	ServiceName      string
	ServicePrice     uint
	GrossAmmount     uint
	PaymentMethod    string
	OrderStatus      string
	PayoutReceiptUrl string
	PayoutDate       time.Time
	ServiceID        uint
	ClientID         uint
}

type ServiceInterface interface {
	GetAll(query string) (data []Core, err error)
	GetById(id uint) (data Core, err error)
	Create(input Core, c echo.Context) error
	Update(input Core, id uint, c echo.Context) error
	Delete(id uint) error
	GetOrderById(id uint) (data []OrderCore, err error)
}

type RepositoryInterface interface {
	GetAll() (data []Core, err error)
	GetAllWithSearch(query string) (data []Core, err error)
	GetById(id uint) (data Core, err error)
	Create(input Core) error
	Update(input Core, id uint) error
	Delete(id uint) error
	FindUser(email string) (data Core, err error)
	GetOrderById(id uint) (data []OrderCore, err error)
}
