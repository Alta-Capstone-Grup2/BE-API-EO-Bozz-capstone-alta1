package service

import (
	"capstone-alta1/features/additional"
	"capstone-alta1/utils/helper"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/gommon/log"
)

type additionalService struct {
	additionalRepository additional.RepositoryInterface
	validate             *validator.Validate
}

func New(repo additional.RepositoryInterface) additional.ServiceInterface {
	return &additionalService{
		additionalRepository: repo,
		validate:             validator.New(),
	}
}

func (service *additionalService) Create(input additional.Core) (err error) {
	// validasi input
	errCreate := service.additionalRepository.Create(input)
	if errCreate != nil {
		log.Error(errCreate.Error())
		return helper.ServiceErrorMsg(err)
	}

	return nil
}

// GetAll implements user.ServiceInterface
func (service *additionalService) GetAll() (data []additional.Core, err error) {

	data, err = service.additionalRepository.GetAll()

	if err != nil {
		helper.LogDebug(err)
		return nil, helper.ServiceErrorMsg(err)
	}

	return data, err
}

func (service *additionalService) Update(input additional.Core, id uint) error {
	err := service.additionalRepository.Update(input, id)
	if err != nil {
		log.Error(err.Error())
		return helper.ServiceErrorMsg(err)
	}

	return nil
}

func (service *additionalService) Delete(id uint) error {
	err := service.additionalRepository.Delete(id)
	if err != nil {
		log.Error(err.Error())
		return helper.ServiceErrorMsg(err)
	}
	return nil
}
