package repository

import (
	_order "capstone-alta1/features/order"
	"capstone-alta1/utils/helper"
	"errors"

	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) _order.RepositoryInterface {
	return &orderRepository{
		db: db,
	}
}

func (repo *orderRepository) Create(inputOrder _order.Core, inputDetail []_order.DetailOrder) error {
	orderGorm := fromCore(inputOrder)
	detailorderGorm := fromDetailOrderList(inputDetail)

	tx := repo.db.Create(&orderGorm) // proses insert data
	if tx.Error != nil {
		helper.LogDebug("Order - query - Create | Error execute query order. Error  = ", tx.Error)
		return tx.Error
	}
	helper.LogDebug("Order - query - create | Row Affected query order : ", tx.RowsAffected)
	if tx.RowsAffected == 0 {
		return errors.New("insert order failed")
	}

	for idx := range detailorderGorm {
		detailorderGorm[idx].OrderID = orderGorm.ID
	}

	helper.LogDebug("Order - query - create | Add order id to order detail slice. []Detail Order : ", detailorderGorm)

	yx := repo.db.Create(&detailorderGorm) // proses insert data
	if yx.Error != nil {
		helper.LogDebug("Order - query - Create | Error execute query detail order. Error  = ", yx.Error)
		return yx.Error
	}
	helper.LogDebug("Order - query - Create | Row Affected query detail order : ", yx.RowsAffected)
	if yx.RowsAffected == 0 {

		return errors.New("insert detail order failed")
	}
	return nil
}

func (do *DetailOrder) BeforeCreate(tx *gorm.DB) (err error) {
	var additionalData Additional
	txBeforeCreate := tx.Raw("SELECT `additionals`.`additional_name`, `additionals`.`additional_price`  FROM `additionals` JOIN `service_additionals` ON `additionals`.`id` = `service_additionals`.`additional_id` WHERE `service_additionals`.`id` = ?;", do.ServiceAdditionalID).Find(&additionalData)

	if txBeforeCreate.Error != nil {
		helper.LogDebug("Order - query - BeforeCreate | Error execute query. Error  = ", txBeforeCreate.Error)
		return txBeforeCreate.Error
	}

	helper.LogDebug("Order - query - BeforeCreate | Row Affected query get additional data : ", txBeforeCreate.RowsAffected)
	if txBeforeCreate.RowsAffected == 0 {
		return txBeforeCreate.Error
	}

	helper.LogDebug("Order - query - BeforeCreate | additionalData = ", additionalData)

	do.AdditionalName = additionalData.AdditionalName
	do.AdditionalPrice = additionalData.AdditionalPrice
	do.DetailOrderTotal = do.Qty * do.AdditionalPrice

	helper.LogDebug("Order - query - BeforeCreate | additionalData = ", additionalData)
	return
}

func (repo *orderRepository) GetAll() (data []_order.Core, err error) {
	var results []Order

	tx := repo.db.Find(&results)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(results)
	return dataCore, nil
}

func (repo *orderRepository) GetAllWithSearch(query string) (data []_order.Core, err error) {
	var order []Order
	tx := repo.db.Where("name LIKE ?", "%"+query+"%").Find(&order)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected == 0 {
		return data, tx.Error
	}

	var dataCore = toCoreList(order)
	return dataCore, nil
}

func (repo *orderRepository) GetById(id uint) (data _order.Core, dataDetail _order.DetailOrder, err error) {
	var order Order
	var detail DetailOrder

	tx := repo.db.First(&order, id)
	yx := repo.db.First(&detail)

	if tx.Error != nil && yx.Error != nil {
		return data, dataDetail, tx.Error
	}

	if tx.RowsAffected == 0 && yx.RowsAffected == 0 {
		return data, dataDetail, tx.Error
	}

	var dataCore = order.toCoreOrder()
	var dataCoreDetail = detail.toCoreDetailOrder()
	return dataCore, dataCoreDetail, nil
}

func (repo *orderRepository) UpdateStatusCancel(input _order.Core, id uint) error {
	resultGorm := fromCoreStatus(input)
	var result Order
	tx := repo.db.Model(&result).Where("ID = ?", id).Updates(&resultGorm.OrderStatus) // proses update
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("update failed")
	}
	return nil
}

func (repo *orderRepository) UpdateStatusPayout(input _order.Core, id uint) error {
	resultGorm := fromCorePayout(input)
	var result Order
	tx := repo.db.Model(&result).Where("ID = ?", id).Updates(&resultGorm) // proses update
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("update failed")
	}
	return nil
}
