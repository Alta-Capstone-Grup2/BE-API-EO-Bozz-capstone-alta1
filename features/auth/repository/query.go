package repository

import (
	"capstone-alta1/features/auth"
	"errors"

	"gorm.io/gorm"
)

type authData struct {
	db *gorm.DB
}

func New(db *gorm.DB) auth.RepositoryInterface {
	return &authData{
		db: db,
	}
}

func (repo *authData) FindUser(email string) (result auth.Core, err error) {
	var userData User
	tx := repo.db.Where("email = ?", email).First(&userData)

	if tx.Error != nil {
		return auth.Core{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return auth.Core{}, errors.New("login failed")
	}

	result = userData.toCore()

	return result, nil
}

func (repo *authData) FindClient(userID uint) (result auth.ClientCore, err error) {
	var client Client

	tx := repo.db.Where("user_id", userID).Preload("User").First(&client)

	if tx.Error != nil {
		return result, tx.Error
	}

	if tx.RowsAffected == 0 {
		return result, tx.Error
	}

	var dataCore = client.toCore()
	return dataCore, nil
}

func (repo *authData) FindPartner(userID uint) (result auth.PartnerCore, err error) {
	var partner Partner

	tx := repo.db.Where("user_id", userID).Preload("User").First(&partner)

	if tx.Error != nil {
		return result, tx.Error
	}

	if tx.RowsAffected == 0 {
		return result, tx.Error
	}

	var dataCore = partner.toCore()
	return dataCore, nil
}
