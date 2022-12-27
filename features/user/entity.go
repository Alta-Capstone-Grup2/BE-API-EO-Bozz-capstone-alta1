package user

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID        uint
	Name      string `valiidate:"required"`
	Email     string `valiidate:"required,email"`
	Password  string `valiidate:"required"`
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ServiceInterface interface {
	GetAll(query string) (data []Core, err error)
	Create(input Core, c echo.Context) error
	GetById(id uint) (data Core, err error)
	Update(input Core, id uint) error
	Delete(id uint) error
	UpdatePassword(input Core, id uint) error
}

type RepositoryInterface interface {
	GetAll() (data []Core, err error)
	Create(input Core) error
	GetById(id uint) (data Core, err error)
	Update(input Core, id uint) error
	Delete(id uint) error
	FindUser(email string) (data Core, err error)
	UpdatePassword(input Core, id uint) error
}
