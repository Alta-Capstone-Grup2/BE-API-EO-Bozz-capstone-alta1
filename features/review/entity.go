package review

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID        uint
	Review    string  `valiidate:"required"`
	Rating    float64 `valiidate:"required"`
	OrderID   uint    `valiidate:"required"`
	ClientID  uint    `valiidate:"required"`
	ServiceID uint    `valiidate:"required"`
	Order     Order
	Client    Client
	Service   Service
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

type Service struct {
	ID                 uint
	ServiceName        string
	ServiceDescription string
	ServiceIncluded    string
	ServiceCategory    string
	ServicePrice       uint
	AverageRating      float64
	ServiceImage_Url   string
	City               string
	PartnerID          uint
}

type ServiceInterface interface {
	GetAll(query string) (data []Core, err error)
	Create(input Core, c echo.Context) error
	GetById(id int) (data Core, err error)
	Update(input Core, id int, c echo.Context) error
	Delete(id int) error
}

type RepositoryInterface interface {
	GetAll() (data []Core, err error)
	Create(input Core) error
	GetById(id int) (data Core, err error)
	Update(input Core, id int) error
	Delete(id int) error
}
