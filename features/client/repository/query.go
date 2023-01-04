package repository

import (
	client "capstone-alta1/features/client"
	"capstone-alta1/utils/helper"
	"errors"

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
func (repo *clientRepository) GetAll(query string) (data []client.Core, err error) {
	var clients []Client
	// var users []User
	if query != "" {
		tx := repo.db.Raw("SELECT `clients`.`id`,`users`.`name`,`users`.`email`,`users`.`role`,`clients`.`gender`,`clients`.`address`,`client`.`city`,`clients`.`phone`,`clients`.`client_image_file`,`users`.`id` FROM `clients` LEFT JOIN `users` ON `clients`.`users_id` = `users`.`id` WHERE `users`.`name` LIKE ?", query)
		if tx.Error != nil {
			return nil, tx.Error
		}
		var dataCore = toCoreList(clients)
		return dataCore, nil
	} else {
		tx := repo.db.Preload("User").Find(&clients)
		if tx.Error != nil {
			return nil, tx.Error
		}
		var dataCore = toCoreList(clients)
		return dataCore, nil
	}
}

// func (repo *clientRepository) GetAllWithSearch(query string) (data []client.Core, err error) {
// 	var client []Client
// 	tx := repo.db.Preload("User").Where("name LIKE ?", "%"+query+"%").Find(&client)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}

// 	if tx.RowsAffected == 0 {
// 		return data, tx.Error
// 	}

// 	fmt.Println("\n\n 2 getall client = ", client)

// 	var dataCore = toCoreList(client)
// 	return dataCore, nil
// }

// GetById implements user.RepositoryInterface
func (repo *clientRepository) GetById(id uint) (data client.Core, err error) {
	var client Client

	tx := repo.db.Preload("User").First(&client, id)

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
func (repo *clientRepository) Update(input client.Core, clientID uint, userID uint) error {
	clientGorm := fromCore(input)
	var client Client
	var user User

	tx := repo.db.Model(&user).Where("ID = ?", userID).Updates(&clientGorm.User)
	yx := repo.db.Model(&client).Where("ID = ?", clientID).Updates(&clientGorm)
	if tx.Error != nil && yx.Error != nil {
		return errors.New("failed update client")
	}
	if tx.RowsAffected == 0 {
		return errors.New("update failed")
	}
	return nil
}

// Delete implements user.Repository
func (repo *clientRepository) Delete(clientID uint, userID uint) error {
	var client Client
	var user User
	tx := repo.db.Delete(&client, clientID)
	yx := repo.db.Delete(&user, userID)
	if tx.Error != nil && yx.Error != nil {
		return errors.New("failed update client")
	}
	if tx.RowsAffected == 0 {
		return errors.New("update failed")
	}
	return nil
}

func (repo *clientRepository) FindUser(email string) (result client.Core, err error) {
	var clientData Client
	tx := repo.db.Where("email = ?", email).First(&clientData.User)
	if tx.Error != nil {
		return client.Core{}, tx.Error
	}

	result = clientData.toCore()

	return result, nil
}

func (repo *clientRepository) GetOrderById(clientId uint) (data []client.Order, err error) {
	var clientorder []Order

	tx := repo.db.Where("client_id = ?", clientId).Find(&clientorder)

	if tx.Error != nil {
		helper.LogDebug("client-query-Getorder | Error execute query. Error :", tx.Error)
		return data, tx.Error
	}

	helper.LogDebug("client-query-Getorder  | Row Affected : ", tx.RowsAffected)
	if tx.RowsAffected == 0 {
		return data, tx.Error
	}

	var dataCore = toCoreListOrder(clientorder)
	return dataCore, nil
}
