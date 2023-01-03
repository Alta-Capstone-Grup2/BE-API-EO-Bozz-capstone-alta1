package service

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID                 uint
	ServiceName        string
	ServiceIncluded    string
	ServiceDescription string
	ServiceCategory    string
	ServicePrice       uint
	AverageRating      float64
	ServiceImageFile   string
	City               string
	PartnerID          uint
	ServiceAdditional  []ServiceAdditional
	Review             []Review
	Discussion         []Discussion
	Order              []Order
}

type Discussion struct {
	ID        uint
	Comment   string
	CreatedAt time.Time
	PartnerID uint
	ClientID  uint
	ServiceID uint
}

type Review struct {
	ID        uint
	Review    string
	Rating    float64
	OrderID   uint
	ClientID  uint
	ServiceID uint
}

type ServiceAdditional struct {
	ID           uint
	AdditionalID uint
	ServiceID    uint
}

type Additional struct {
	ID                uint
	AdditionalName    string
	AdditionalPrice   uint
	PartnerID         uint
	ServiceAdditional []ServiceAdditional
}

type JoinServiceAdditional struct {
	ServiceAdditionalID uint
	AdditionalName      string
	AdditionalPrice     uint
	ServiceName         string
	ServiceID           uint
	AdditionalID        uint
	PartnerID           uint
}

type Order struct {
	ID                 uint
	EventName          string
	ServiceName        string
	StartDate          time.Time
	EndDate            time.Time
	AvailabilityStatus string
	ServiceID          uint
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
	UserID             uint
	User               User
	Additional         []Additional
	Service            []Core
}

type User struct {
	ID       uint
	Name     string
	Email    string
	Password string
	Role     string
}

type ServiceInterface interface {
	GetAll(queryName, queryCategory, queryCity, queryMinPrice, queryMaxPrice string) (data []Core, err error)
	GetById(id uint) (data Core, err error)
	Create(input Core, c echo.Context) error
	Update(input Core, id uint, c echo.Context) error
	Delete(id uint) error
	GetServiceAdditionalById(id uint) (data []JoinServiceAdditional, err error)
	GetReviewById(id uint) (data []Review, err error)
	GetDiscussionById(id uint) (data []Discussion, err error)
	AddAdditionalToService(input []ServiceAdditional) error
	CheckAvailability(serviceId uint, queryStart, queryEnd string) (data Order, err error)
}

type RepositoryInterface interface {
	GetAll(queryName, queryCategory, queryCity, queryMinPrice, queryMaxPrice string) (data []Core, err error)
	GetById(id uint) (data Core, err error)
	Create(input Core) error
	Update(input Core, id uint) error
	Delete(id uint) error
	GetServiceAdditionalById(id uint) (data []JoinServiceAdditional, err error)
	GetReviewById(id uint) (data []Review, err error)
	GetDiscussionById(id uint) (data []Discussion, err error)
	AddAdditionalToService(input []ServiceAdditional) error
	CheckAvailability(serviceId uint, queryStart, queryEnd time.Time) (data Order, err error)
}
