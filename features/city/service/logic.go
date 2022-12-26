package service

import (
	"capstone-alta1/features/city"
	"capstone-alta1/utils/helper"

	"github.com/go-playground/validator/v10"
)

type cityService struct {
	cityRepository city.RepositoryInterface
	validate       *validator.Validate
}

func New(repo city.RepositoryInterface) city.ServiceInterface {
	return &cityService{
		cityRepository: repo,
		validate:       validator.New(),
	}
}

func (service *cityService) GetAll() (data []city.Core, err error) {

	data, err = service.cityRepository.GetAll()

	if err != nil {
		helper.LogDebug(err)
		return nil, helper.ServiceErrorMsg(err)
	}

	return data, err
}
