package repository

import (
	"capstone-alta1/features/additional"
	"errors"

	"gorm.io/gorm"
)

type additionalRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) additional.RepositoryInterface {
	return &additionalRepository{
		db: db,
	}
}

func (repo *additionalRepository) Create(input additional.Core) error {
	additionalGorm := fromCore(input)
	tx := repo.db.Create(&additionalGorm) // proses insert data
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}
	return nil
}

func (repo *additionalRepository) GetAll() (data []additional.Core, err error) {
	var results []Additional

	tx := repo.db.Find(&results)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(results)
	return dataCore, nil
}

func (repo *additionalRepository) Update(input additional.Core, id uint) error {
	resultGorm := fromCore(input)
	var result Additional
	tx := repo.db.Model(&result).Where("ID = ?", id).Updates(&resultGorm) // proses update
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("update failed")
	}
	return nil
}

func (repo *additionalRepository) Delete(id uint) error {
	var result Additional
	tx := repo.db.Delete(&result, id) // proses delete
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("delete failed")
	}
	return nil
}
