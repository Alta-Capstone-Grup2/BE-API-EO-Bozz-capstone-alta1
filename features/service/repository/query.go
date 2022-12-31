package repository

import (
	_service "capstone-alta1/features/service"
	"capstone-alta1/utils/helper"
	"errors"
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type serviceRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) _service.RepositoryInterface {
	return &serviceRepository{
		db: db,
	}
}

func (repo *serviceRepository) Create(input _service.Core) error {
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

func (repo *serviceRepository) GetAll(queryName, queryCategory, queryCity, queryMinPrice, queryMaxPrice string) (data []_service.Core, err error) {
	var results []Service
	minInt, _ := strconv.Atoi(queryMinPrice)
	maxInt, _ := strconv.Atoi(queryMaxPrice)
	tx := repo.db.Where(&Service{ServiceName: queryName, ServiceCategory: queryCategory, City: queryCity}).Where("service_price BETWEEN ? AND ?", uint(minInt), uint(maxInt)).Find(&results)

	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreListGetAll(results)
	return dataCore, nil
}

func (repo *serviceRepository) GetById(id uint) (data _service.Core, err error) {
	var service Service

	tx := repo.db.First(&service, id)

	if tx.Error != nil {
		return data, tx.Error
	}

	if tx.RowsAffected == 0 {
		return data, tx.Error
	}

	var dataCore = service.toCoreGetById()
	return dataCore, nil
}

func (repo *serviceRepository) Update(input _service.Core, id uint) error {
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

func (repo *serviceRepository) GetAdditionalById(serviceId uint) (data []_service.Additional, err error) {
	var modelData []Additional

	// tx := repo.db.Model(&serviceAdditional).Where("service_id = ?", serviceId).Find(&additional, serviceAdditionalId.AdditionalID)
	tx := repo.db.Raw("SELECT `additionals`.`id`,`additionals`.`additional_name`,`additionals`.`additional_price`,`additionals`.`partner_id` FROM `additionals` JOIN `partners` ON `additionals`.`partner_id` = `partners`.`id` JOIN `services` ON `services`.`partner_id` = `partners`.`id`;").Where("`services.id` = ?", serviceId).Scan(&modelData)

	if tx.Error != nil {
		helper.LogDebug("service-query-GetAdditional | Error execute query. Error :", tx.Error)
		return data, tx.Error
	}

	helper.LogDebug("service-query-GetAdditional  | Row Affected : ", tx.RowsAffected)
	if tx.RowsAffected == 0 {
		return data, tx.Error
	}

	var dataCore = toCoreListAdditional(modelData)
	return dataCore, nil
}

func (repo *serviceRepository) AddAdditionalToService(input _service.ServiceAdditional, inputAdditional []_service.ServiceAdditional) error {
	serviceadditionalGorm := fromCoreServiceAdditional(input)
	additionallistGorm := fromCoreAdditionalList(inputAdditional)
	var service Service
	tx := repo.db.Model(&service).Where("ID = ?", input.ServiceID).Create(&serviceadditionalGorm) // proses insert data
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}
	for index := range additionallistGorm {
		additionallistGorm[index].AdditionalID = serviceadditionalGorm.AdditionalID
	}
	yx := repo.db.Create(&additionallistGorm) // proses insert data
	if yx.Error != nil {
		return tx.Error
	}
	if yx.RowsAffected == 0 {
		return errors.New("insert failed")
	}
	return nil
}

func (repo *serviceRepository) GetReviewById(serviceId uint) (data []_service.Review, err error) {
	var clientreview []Review

	tx := repo.db.Where("service_id = ?", serviceId).Find(&clientreview)

	if tx.Error != nil {
		helper.LogDebug("service-query-Getreview | Error execute query. Error :", tx.Error)
		return data, tx.Error
	}

	helper.LogDebug("service-query-Getreview  | Row Affected : ", tx.RowsAffected)
	if tx.RowsAffected == 0 {
		return data, tx.Error
	}

	var dataCore = toCoreListReview(clientreview)
	return dataCore, nil
}

func (repo *serviceRepository) GetDiscussionById(serviceId uint) (data []_service.Discussion, err error) {
	var clientdiscussion []Discussion

	tx := repo.db.Where("service_id = ?", serviceId).Find(&clientdiscussion)

	if tx.Error != nil {
		helper.LogDebug("service-query-Getdiscussion | Error execute query. Error :", tx.Error)
		return data, tx.Error
	}

	helper.LogDebug("service-query-Getdiscussion  | Row Affected : ", tx.RowsAffected)
	if tx.RowsAffected == 0 {
		return data, tx.Error
	}

	var dataCore = toCoreListDiscussion(clientdiscussion)
	return dataCore, nil
}

func (repo *serviceRepository) CheckAvailability(serviceId uint, queryStart, queryEnd time.Time) (data _service.Order, err error) {
	//check available
	var services []Service
	var service Service
	queryBuilder := fmt.Sprintf("SELECT * FROM orders WHERE service_id = %d AND (('%s' BETWEEN start_date AND end_date) OR ('%s' BETWEEN start_date AND end_date));", serviceId, queryStart, queryEnd)
	tx := repo.db.Raw(queryBuilder).Find(&services)

	//get data service
	yx := repo.db.First(&service, serviceId)
	if yx.Error != nil {
		return _service.Order{}, err
	}

	//create return
	var orders Order
	serviceName := service.ServiceName
	statusAvailable := "Available"
	statusNotvalable := "Not Available"

	if tx.Error != nil {
		return orders.toCoreNotAvailable(serviceName, queryStart, queryEnd, statusNotvalable), tx.Error
	}

	affectedRow := tx.RowsAffected
	fmt.Println("\n\nHasil check availbility, \n Checkin date = ", queryStart, " \n Checkout date = ", queryEnd, " \n Hasil Row = ", affectedRow)

	if affectedRow == 0 {
		return orders.toCoreAvailable(serviceName, queryStart, queryEnd, statusAvailable), nil
	}

	return orders.toCoreNotAvailable(serviceName, queryStart, queryEnd, statusNotvalable), nil
}
