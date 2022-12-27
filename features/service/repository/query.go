package repository

import (
	"capstone-alta1/features/service"
	"capstone-alta1/utils/helper"
	"errors"
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

type serviceRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) service.RepositoryInterface {
	return &serviceRepository{
		db: db,
	}
}

func (repo *serviceRepository) Create(input service.Core) error {
	serviceGorm := fromCore(input)
	tx := repo.db.Create(&serviceGorm) // proses insert data
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}
	return nil
}

func (repo *serviceRepository) GetAll() (data []service.Core, err error) {
	var results []Service

	tx := repo.db.Find(&results)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(results)
	return dataCore, nil
}

func (repo *serviceRepository) GetAllWithSearch(queryName, queryCategory, queryCity, queryMinPrice, queryMaxPrice string) (data []service.Core, err error) {
	var services, services2 []Service

	helper.LogDebug("\n isi queryName = ", queryName)
	helper.LogDebug("\n isi queryCategory= ", queryCategory)
	helper.LogDebug("\n isi queryCity = ", queryCity)
	helper.LogDebug("\n isi queryMinPrice = ", queryMinPrice)
	helper.LogDebug("\n isi queryMaxPrice = ", queryMaxPrice)

	intMinPrice, errConv1 := strconv.Atoi(queryMinPrice)
	intMaxPrice, errConv2 := strconv.Atoi(queryMaxPrice)
	if errConv1 != nil || errConv2 != nil {
		return nil, errors.New("error conver service price to filter")
	}

	fmt.Println("\n\nServices 1", services)
	tx := repo.db.Where("service_name LIKE ?", "%"+queryName+"%").Where(&Service{ServiceCategory: queryCategory, City: queryCity, ServicePrice: uint(intMinPrice) + uint(intMaxPrice)}).Find(&services2)
	fmt.Println("\n\nServices 2", services2)

	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(services2)
	return dataCore, nil
}

func (repo *serviceRepository) GetById(id uint) (data service.Core, err error) {
	var service Service

	tx := repo.db.First(&service, id)

	if tx.Error != nil {
		return data, tx.Error
	}

	if tx.RowsAffected == 0 {
		return data, tx.Error
	}

	var dataCore = service.toCore()
	return dataCore, nil
}

func (repo *serviceRepository) Update(input service.Core, id uint) error {
	resultGorm := fromCore(input)
	var result Service
	tx := repo.db.Model(&result).Where("ID = ?", id).Updates(&resultGorm) // proses update
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("update failed")
	}
	return nil
}

func (repo *serviceRepository) Delete(id uint) error {
	var result Service
	tx := repo.db.Delete(&result, id) // proses delete
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("delete failed")
	}
	return nil
}

func (repo *serviceRepository) GetAdditionalById(id uint) (data []service.Additional, err error) {
	var clientadditional []Additional

	tx := repo.db.Find(&clientadditional, id)

	if tx.Error != nil {
		return data, tx.Error
	}

	if tx.RowsAffected == 0 {
		return data, tx.Error
	}

	var dataCore = toCoreListAdditional(clientadditional)
	return dataCore, nil
}

func (repo *serviceRepository) AddAdditionalToService(input service.ServiceAdditional, id uint) error {
	additionalGorm := fromCoreServiceAdditional(input)
	var service Service
	tx := repo.db.Model(&service).Where("ID = ?", id).Create(&additionalGorm) // proses insert data
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}
	return nil
}
