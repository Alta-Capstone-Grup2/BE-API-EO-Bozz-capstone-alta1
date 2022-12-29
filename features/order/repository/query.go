package repository

import (
	_order "capstone-alta1/features/order"
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

func (repo *orderRepository) Create(inputOrder _order.Core, inputDetail _order.DetailOrder) error {
	orderGorm := fromCore(inputOrder)
	detailorderGorm := fromDetailOrder(inputDetail)
	tx := repo.db.Create(&orderGorm) // proses insert data
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}
	yx := repo.db.Create(&detailorderGorm) // proses insert data
	if yx.Error != nil {
		return yx.Error
	}
	if yx.RowsAffected == 0 {
		return errors.New("insert failed")
	}
	return nil
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
