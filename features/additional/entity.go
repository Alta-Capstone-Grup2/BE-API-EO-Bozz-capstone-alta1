package additional

type Core struct {
	ID              uint
	AdditionalName  string
	AdditionalPrice uint
	PartnerID       uint
	Partner         Partner
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
}

type User struct {
	ID       uint
	Name     string
	Email    string
	Password string
	Role     string
}

type ServiceInterface interface {
	GetAll() (data []Core, err error)
	Create(input Core) error
	Update(input Core, id uint) error
	Delete(id uint) error
}

type RepositoryInterface interface {
	GetAll() (data []Core, err error)
	Create(input Core) error
	Update(input Core, id uint) error
	Delete(id uint) error
}
