package service

import (
	cfg "capstone-alta1/config"
	"capstone-alta1/features/client"
	"capstone-alta1/utils/helper"
	"capstone-alta1/utils/thirdparty"
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

type clientService struct {
	clientRepository client.RepositoryInterface
	validate         *validator.Validate
}

func New(repo client.RepositoryInterface) client.ServiceInterface {
	return &clientService{
		clientRepository: repo,
		validate:         validator.New(),
	}
}

// Create implements user.ServiceInterface
func (service *clientService) Create(input client.Core, c echo.Context) (err error) {
	input.User.Role = "Client"
	// validasi input
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	// validasi email harus unik
	data, errFindEmail := service.clientRepository.FindUser(input.User.Email)
	// helper.LogDebug("\n\n\n find email input  ", input.Email)
	// helper.LogDebug("\n\n\n find email data  ", data.Email)

	if data.User.Email == input.User.Email {
		return errors.New("Email " + input.User.Email + " already exist. Please pick another email.")
	}

	if errFindEmail != nil && !strings.Contains(errFindEmail.Error(), "found") {
		return helper.ServiceErrorMsg(errFindEmail)
	}

	bytePass, errEncrypt := bcrypt.GenerateFromPassword([]byte(input.User.Password), 10)
	if errEncrypt != nil {
		log.Error(errEncrypt.Error())
		return err
	}

	input.User.Password = string(bytePass)

	// upload foto
	var errUpload error
	input.ClientImageUrl, errUpload = thirdparty.Upload(c, cfg.CLIENT_IMAGE_FILE, cfg.CLIENT_FOLDER)
	if errUpload != nil {
		return errUpload
	}

	errCreate := service.clientRepository.Create(input)
	if errCreate != nil {
		log.Error(errCreate.Error())
		return err
	}

	return nil
}

// GetAll implements user.ServiceInterface
func (service *clientService) GetAll(query string) (data []client.Core, err error) {
	if query == "" {
		data, err = service.clientRepository.GetAll()
	} else {
		data, err = service.clientRepository.GetAllWithSearch(query)
	}

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return data, err
}

func (service *clientService) GetById(id uint) (data client.Core, err error) {
	data, err = service.clientRepository.GetById(id)
	if err != nil {
		log.Error(err.Error())
		return client.Core{}, err
	}

	return data, err

}

func (service *clientService) Update(input client.Core, clientID uint, userID uint, c echo.Context) error {
	if input.User.Password != "" {
		generate, _ := bcrypt.GenerateFromPassword([]byte(input.User.Password), 10)
		input.User.Password = string(generate)
	}

	// upload file
	var errUpload error
	input.ClientImageUrl, errUpload = thirdparty.Upload(c, cfg.CLIENT_IMAGE_FILE, cfg.CLIENT_FOLDER)
	if errUpload != nil {
		return errUpload
	}

	err := service.clientRepository.Update(input, clientID, userID)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (service *clientService) Delete(id uint) error {
	// proses
	err := service.clientRepository.Delete(id)
	if err != nil {
		log.Error(err.Error())
		return helper.ServiceErrorMsg(err)
	}
	return nil
}

func (service *clientService) GetOrderById(id uint) (data []client.Order, err error) {
	data, err = service.clientRepository.GetOrderById(id)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return data, err

}
