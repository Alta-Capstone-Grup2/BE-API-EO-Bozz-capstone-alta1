package client

import (
	"capstone-alta1/features/user"

	"github.com/labstack/echo/v4"
)

type Core struct {
	User           user.Core
	Gender         string
	Address        string
	City           string
	Phone          string
	ClientImageUrl string
}

type ServiceInterface interface {
	GetAll(query string) (data []Core, err error)
	GetById(id int) (data Core, err error)
	Create(input Core, c echo.Context) error
	Update(input Core, id int, c echo.Context) error
	Delete(id int) error
}

type RepositoryInterface interface {
	GetAll() (data []Core, err error)
	GetAllWithSearch(query string) (data []Core, err error)
	GetById(id int) (data Core, err error)
	Create(input Core) error
	Update(input Core, id int) error
	Delete(id int) error
	FindUser(email string) (data Core, err error)
}
