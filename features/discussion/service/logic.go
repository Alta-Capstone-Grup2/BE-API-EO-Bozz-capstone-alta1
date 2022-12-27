package service

import (
	"capstone-alta1/features/discussion"
	"capstone-alta1/utils/helper"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type discussionService struct {
	discussionRepository discussion.RepositoryInterface
	validate             *validator.Validate
}

func New(repo discussion.RepositoryInterface) discussion.ServiceInterface {
	return &discussionService{
		discussionRepository: repo,
		validate:             validator.New(),
	}
}

func (service *discussionService) Create(input discussion.Core, c echo.Context) (err error) {
	// validasi input
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	fmt.Println("Cek discussion d service data = ", input)

	errCreate := service.discussionRepository.Create(input)
	if errCreate != nil {
		log.Error(errCreate.Error())
		return helper.ServiceErrorMsg(err)
	}

	return nil
}

// GetAll implements user.ServiceInterface
func (service *discussionService) GetAll() (data []discussion.Core, err error) {

	data, err = service.discussionRepository.GetAll()

	if err != nil {
		helper.LogDebug(err)
		return nil, helper.ServiceErrorMsg(err)
	}

	return data, err
}

func (service *discussionService) GetById(id uint) (data discussion.Core, err error) {
	data, err = service.discussionRepository.GetById(id)
	if err != nil {
		log.Error(err.Error())
		return discussion.Core{}, helper.ServiceErrorMsg(err)
	}

	return data, err

}

func (service *discussionService) Update(input discussion.Core, id uint, c echo.Context) error {
	// validasi input
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	// validasi user dgn id path param, apakah ada datanya di database
	_, errFindId := service.discussionRepository.GetById(id)
	if errFindId != nil {
		log.Error(errFindId.Error())
		return helper.ServiceErrorMsg(errFindId)
	}

	// proses
	err := service.discussionRepository.Update(input, id)
	if err != nil {
		log.Error(err.Error())
		return helper.ServiceErrorMsg(err)
	}

	return nil
}

func (service *discussionService) Delete(id uint) error {
	// validasi user dgn id path param, apakah ada datanya di database
	_, errFindId := service.discussionRepository.GetById(id)
	if errFindId != nil {
		log.Error(errFindId.Error())
		return helper.ServiceErrorMsg(errFindId)
	}

	// proses
	err := service.discussionRepository.Delete(id)
	if err != nil {
		log.Error(err.Error())
		return helper.ServiceErrorMsg(err)
	}
	return nil
}
