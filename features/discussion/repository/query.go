package repository

import (
	"capstone-alta1/features/discussion"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type discussionRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) discussion.RepositoryInterface {
	return &discussionRepository{
		db: db,
	}
}

// Create implements user.Repository
func (repo *discussionRepository) Create(input discussion.Core) error {
	dataModel := fromCore(input)

	fmt.Println("Cek discussion d service data = ", dataModel)

	tx := repo.db.Create(&dataModel) // proses insert data
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}

	return nil
}

// GetAll implements user.Repository
func (repo *discussionRepository) GetAll() (data []discussion.Core, err error) {
	var results []Discussion

	tx := repo.db.Find(&results)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(results)
	return dataCore, nil
}

// GetAll with search by name implements user.Repository
func (repo *discussionRepository) GetAllWithSearch(query string) (data []discussion.Core, err error) {
	var results []Discussion

	tx := repo.db.Where("title LIKE ?", "%"+query+"%").Find(&results)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(results)
	return dataCore, nil
}

// GetById implements user.RepositoryInterface
func (repo *discussionRepository) GetById(id int) (data discussion.Core, err error) {
	var result Discussion

	tx := repo.db.First(&result, id)

	if tx.Error != nil {
		return data, tx.Error
	}

	if tx.RowsAffected == 0 {
		return data, tx.Error
	}

	var dataCore = result.toCore()
	return dataCore, nil
}

// Update implements user.Repository
func (repo *discussionRepository) Update(input discussion.Core, id int) error {
	resultGorm := fromCore(input)
	var result Discussion
	tx := repo.db.Model(&result).Where("ID = ?", id).Updates(&resultGorm) // proses update
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("update failed")
	}
	return nil
}

// Delete implements user.Repository
func (repo *discussionRepository) Delete(id int) error {
	var result Discussion
	tx := repo.db.Delete(&result, id) // proses delete
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("delete failed")
	}
	return nil
}
