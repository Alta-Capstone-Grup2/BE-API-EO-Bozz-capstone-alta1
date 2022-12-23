package service

import (
	"capstone-alta1/features/partner"
	"capstone-alta1/utils/helper"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type partnerService struct {
	partnerRepository partner.RepositoryInterface
	validate          *validator.Validate
}

func New(repo partner.RepositoryInterface) partner.ServiceInterface {
	return &partnerService{
		partnerRepository: repo,
		validate:          validator.New(),
	}
}

// Create implements user.ServiceInterface
func (service *partnerService) Create(input partner.Core, c echo.Context) (err error) {
	// validasi input
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	input.User.Role = "Partner"

	errCreate := service.partnerRepository.Create(input)
	if errCreate != nil {
		log.Error(errCreate.Error())
		return err
	}

	return nil
}

// GetAll implements user.ServiceInterface
func (service *partnerService) GetAll(query string) (data []partner.Core, err error) {
	if query == "" {
		data, err = service.partnerRepository.GetAll()
	} else {
		data, err = service.partnerRepository.GetAllWithSearch(query)
	}

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return data, err
}

func (service *partnerService) GetById(id int) (data partner.Core, err error) {
	data, err = service.partnerRepository.GetById(id)
	if err != nil {
		log.Error(err.Error())
		return partner.Core{}, err
	}

	return data, err

}

func (service *partnerService) Update(input partner.Core, id int, c echo.Context) error {
	err := service.partnerRepository.Update(input, id)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (service *partnerService) Delete(id int) error {
	// proses
	err := service.partnerRepository.Delete(id)
	if err != nil {
		log.Error(err.Error())
		return helper.ServiceErrorMsg(err)
	}
	return nil
}
