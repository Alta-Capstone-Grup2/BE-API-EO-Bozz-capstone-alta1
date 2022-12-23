package repository

import (
	client "capstone-alta1/features/client"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type clientRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) client.RepositoryInterface {
	return &clientRepository{
		db: db,
	}
}

// Create implements user.Repository
func (repo *clientRepository) Create(input client.Core) error {
	clientGorm := fromCore(input)
	tx := repo.db.Create(&clientGorm) // proses insert data
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}
	return nil
}

// GetAll implements user.Repository
func (repo *clientRepository) GetAll() (data []client.Core, err error) {
	var client []Client

	tx := repo.db.Find(&client)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(client)
	return dataCore, nil
}

func (repo *clientRepository) GetAllWithSearch(query string) (data []client.Core, err error) {
	var client []Client
	tx := repo.db.Where("name LIKE ?", "%"+query+"%").Find(&client)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected == 0 {
		return data, errors.New("Data not found.")
	}

	fmt.Println("\n\n 2 getall client = ", client)

	var dataCore = toCoreList(client)
	return dataCore, nil
}

// GetById implements user.RepositoryInterface
func (repo *clientRepository) GetById(id int) (data client.Core, err error) {
	var client Client

	tx := repo.db.First(&client, id)

	if tx.Error != nil {
		return data, tx.Error
	}

	if tx.RowsAffected == 0 {
		return data, tx.Error
	}

	var dataCore = client.toCore()
	return dataCore, nil
}

// Update implements user.Repository
func (repo *clientRepository) Update(input client.Core, id int) error {
	clientGorm := fromCore(input)
	var client Client
	tx := repo.db.Model(&client).Where("ID = ?", id).Updates(&clientGorm) // proses update
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("update failed")
	}
	return nil
}

// Delete implements user.Repository
func (repo *clientRepository) Delete(id int) error {
	var client Client
	tx := repo.db.Delete(&client, id) // proses delete
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("delete failed")
	}
	return nil
}

func (repo *clientRepository) FindUser(email string) (result client.Core, err error) {
	var clientData Client
	tx := repo.db.Where("email = ?", email).First(&clientData)
	if tx.Error != nil {
		return client.Core{}, tx.Error
	}

	result = clientData.toCore()

	return result, nil
}
