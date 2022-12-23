package client

import (
	"capstone-alta1/features/order"
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
	Order          []order.Core
}

type ServiceInterface interface {
	GetAll(query string) (data []Core, err error)
	GetById(id uint) (data Core, err error)
	Create(input Core, c echo.Context) error
	Update(input Core, id uint, c echo.Context) error
	Delete(id uint) error
	GetOrderById(id uint) (data []order.Core, err error)
}

type RepositoryInterface interface {
	GetAll() (data []Core, err error)
	GetAllWithSearch(query string) (data []Core, err error)
	GetById(id uint) (data Core, err error)
	Create(input Core) error
	Update(input Core, id uint) error
	Delete(id uint) error
	FindUser(email string) (data Core, err error)
	GetOrderById(id uint) (data []order.Core, err error)
}
