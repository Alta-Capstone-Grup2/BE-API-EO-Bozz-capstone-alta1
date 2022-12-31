package repository

import (
	"capstone-alta1/features/review"
	"errors"

	"gorm.io/gorm"
)

type reviewRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) review.RepositoryInterface {
	return &reviewRepository{
		db: db,
	}
}

func (repo *reviewRepository) Create(input review.Core) error {

	var service Service
	userGorm := fromCore(input)
	tx := repo.db.Create(&userGorm) // proses insert data
	if tx.Error != nil {
		return tx.Rollback().Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}

	tx1 := repo.db.First(&service, input.OrderID)
	if tx1.Error != nil {
		return tx1.Error
	}
	var avg float64
	if service.AverageRating == 0 {
		avg = input.Rating
	} else {
		avg = (service.AverageRating + input.Rating) / 2
	}
	tx2 := repo.db.Model(&service).Where("id = ?", input.ServiceID).Update("average_rating", avg)
	if tx2.Error != nil {
		return tx.Error
	}
	return nil
}

// GetAll implements user.Repository
func (repo *reviewRepository) GetAll() (data []review.Core, err error) {
	var results []Review

	tx := repo.db.Find(&results)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(results)
	return dataCore, nil
}

// GetAll with search by name implements user.Repository
func (repo *reviewRepository) GetAllWithSearch(query string) (data []review.Core, err error) {
	var results []Review

	tx := repo.db.Where("title LIKE ?", "%"+query+"%").Find(&results)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(results)
	return dataCore, nil
}

// GetById implements user.RepositoryInterface
func (repo *reviewRepository) GetById(id uint) (data review.Core, err error) {
	var result Review

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
func (repo *reviewRepository) Update(input review.Core, id uint) error {
	resultGorm := fromCore(input)
	var result Review
	tx := repo.db.Model(&result).Where("ID = ?", id).Updates(&resultGorm) // proses update
	if tx.Error != nil {
		return tx.Rollback().Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("update failed")
	}
	return nil
}

// Delete implements user.Repository
func (repo *reviewRepository) Delete(id uint) error {
	var result Review
	tx := repo.db.Delete(&result, id) // proses delete
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("delete failed")
	}
	return nil
}

func (repo *reviewRepository) FindUser(email string) (result review.Core, err error) {
	var data Review
	tx := repo.db.Where("email = ?", email).First(&data)
	if tx.Error != nil {
		return review.Core{}, tx.Error
	}

	result = data.toCore()

	return result, nil
}
