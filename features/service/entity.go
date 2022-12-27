package service

import "time"

type Core struct {
	ID                 uint
	ServiceName        string
	ServiceDescription string
	ServiceCategory    string
	ServicePrice       uint
	AverageRating      float64
	ServiceImageUrl    string
	City               string
	PartnerID          uint
	Partner            Partner
	Additional         []Additional
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
	Service   Core
}

type Review struct {
	ID        uint
	Review    string
	Rating    float64
	OrderID   uint
	ClientID  uint
	ServiceID uint
	Service   Core
}

type ServiceAdditional struct {
	ID           uint
	AdditionalID uint
	Additional   Additional
	ServiceID    uint
	Service      Core
}

type Additional struct {
	ID              uint
	AdditionalName  string
	AdditionalPrice uint
	PartnerID       uint
	ServiceID       uint
	Service         Core
}

type Order struct {
	ID        uint
	EventName string
	ServiceID uint
}

type Partner struct {
	ID                 uint
	PICPosition        string
	PICPhone           string
	PICAddress         string
	CompanyName        string
	CompanyPhone       string
	CompanyCity        string
	CompanyImageUrl    string
	CompanyAddress     string
	LinkWebsite        string
	NIBNumber          string
	NIBImageUrl        string
	SIUPNumber         string
	SIUPImageUrl       string
	Event1Name         string
	Event1ImageUrl     string
	Event2Name         string
	Event2ImageUrl     string
	Event3Name         string
	Event3ImageUrl     string
	BankName           string
	BankAccountNumber  string
	BankAccountName    string
	VerificationStatus string
	VerificationLog    string
	UserID             uint
	User               User
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
	Create(input Core) error
	Update(input Core, id uint) error
	Delete(id uint) error
	GetAdditionalById(id uint) (data []Additional, err error)
	GetReviewById(id uint) (data []Review, err error)
	GetDiscussionById(id uint) (data []Discussion, err error)
	AddAdditionalToService(input ServiceAdditional, id uint) error
}

type RepositoryInterface interface {
	GetAll() (data []Core, err error)
	GetAllWithSearch(queryName, queryCategory, queryCity, queryMinPrice, queryMaxPrice string) (data []Core, err error)
	GetById(id uint) (data Core, err error)
	Create(input Core) error
	Update(input Core, id uint) error
	Delete(id uint) error
	GetAdditionalById(id uint) (data []Additional, err error)
	GetReviewById(id uint) (data []Review, err error)
	GetDiscussionById(id uint) (data []Discussion, err error)
	AddAdditionalToService(input ServiceAdditional, id uint) error
}
