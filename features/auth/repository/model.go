package repository

import (
	"capstone-alta1/features/auth"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint
	Name      string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

//DTO

func (dataModel User) toCore() auth.Core {
	return auth.Core{
		ID:        dataModel.ID,
		Name:      dataModel.Name,
		Email:     dataModel.Email,
		Password:  dataModel.Password,
		Role:      dataModel.Role,
		CreatedAt: dataModel.CreatedAt,
		UpdatedAt: dataModel.UpdatedAt,
	}
}
