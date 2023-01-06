package repository

import (
	"capstone-alta1/features/auth"
	_user "capstone-alta1/features/user/repository"
	middlewares "capstone-alta1/middlewares"
	"errors"

	"github.com/labstack/gommon/log"
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

func (repo *authData) LoginOauth(auths auth.Oauth) (string, auth.ClientCore, error) {
	var userData _user.User
	// var dataPartner auth.PartnerCore
	// var dataClient auth.ClientCore
	// clientId := dataClient.ID
	// partnerId := dataPartner.User.ID
	// var dataCore auth.Core

	tx := repo.db.Where("email = ?", auths.Email).First(&userData)
	user := _user.User{}
	user.Email = auths.Email
	user.Name = auths.Name

	if tx.Error != nil {
		user.Role = "Client"
		tx1 := repo.db.Create(&user) // proses insert data

		if tx1.Error != nil {
			return "", auth.ClientCore{}, tx1.Error
		}
		if tx1.RowsAffected == 0 {
			return "", auth.ClientCore{}, errors.New("insert failed")
		}

	}
	var client Client
	tx3 := repo.db.Where("ID = ?", userData.ID).Preload("User").First(&client)
	if tx3.Error != nil {
		return "", auth.ClientCore{}, tx3.Error
	}

	token, errToken := middlewares.CreateToken(int(client.UserID), client.User.Name, client.User.Role, int(client.ID), 0)
	if errToken != nil {
		log.Error(errToken.Error())
		return token, auth.ClientCore{}, errors.New("Failed to login, error on generate token, please try again.")
	}

	var dataCore = client.toCore()

	return token, dataCore, nil
}
