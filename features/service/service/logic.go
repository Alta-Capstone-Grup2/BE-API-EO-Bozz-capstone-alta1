package service

import (
	_service "capstone-alta1/features/service"
	"capstone-alta1/utils/helper"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/gommon/log"
)

type serviceService struct {
	serviceRepository _service.RepositoryInterface
	validate          *validator.Validate
}

func New(repo _service.RepositoryInterface) _service.ServiceInterface {
	return &serviceService{
		serviceRepository: repo,
		validate:          validator.New(),
	}
}

func (service *serviceService) Create(input _service.Core) (err error) {
	errCreate := service.serviceRepository.Create(input)
	if errCreate != nil {
		log.Error(errCreate.Error())
		return helper.ServiceErrorMsg(err)
	}

	return nil
}

func (service *serviceService) GetAll(queryName, queryCategory, queryCity, queryMinPrice, queryMaxPrice string) (data []_service.Core, err error) {

	data, err = service.serviceRepository.GetAllWithSearch(queryName, queryCategory, queryCity, queryMinPrice, queryMaxPrice)

	if err != nil {
		helper.LogDebug(err)
		return nil, helper.ServiceErrorMsg(err)
	}

	return data, err
}

func (service *serviceService) GetById(id uint) (data _service.Core, err error) {
	data, err = service.serviceRepository.GetById(id)
	if err != nil {
		log.Error(err.Error())
		return _service.Core{}, err
	}
	return data, err
}

func (service *serviceService) Update(input _service.Core, id uint) error {
	err := service.serviceRepository.Update(input, id)
	if err != nil {
		log.Error(err.Error())
		return helper.ServiceErrorMsg(err)
	}

	return nil
}

func (service *serviceService) Delete(id uint) error {
	err := service.serviceRepository.Delete(id)
	if err != nil {
		log.Error(err.Error())
		return helper.ServiceErrorMsg(err)
	}
	return nil
}

func (service *serviceService) GetAdditionalById(id uint) (data []_service.Additional, err error) {
	data, err = service.serviceRepository.GetAdditionalById(id)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return data, err

}

func (service *serviceService) AddAdditionalToService(input _service.ServiceAdditional, id uint) (err error) {
	errCreate := service.serviceRepository.AddAdditionalToService(input, id)
	if errCreate != nil {
		log.Error(errCreate.Error())
		return helper.ServiceErrorMsg(err)
	}

	return nil
}
