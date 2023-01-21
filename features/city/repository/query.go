package repository

import (
	"capstone-alta1/features/city"

	"gorm.io/gorm"
)

type cityRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) city.RepositoryInterface {
	return &cityRepository{
		db: db,
	}
}

func (repo *cityRepository) GetAll() (data []city.Core, err error) {
	var results []City
	tx := repo.db.Order("city_name asc").Find(&results)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(results)
	return dataCore, nil
}
