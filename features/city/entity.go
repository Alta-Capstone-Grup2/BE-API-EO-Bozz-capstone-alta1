package city

type Core struct {
	ID       uint
	CityName string
}

type ServiceInterface interface {
	GetAll() (data []Core, err error)
}

type RepositoryInterface interface {
	GetAll() (data []Core, err error)
}
