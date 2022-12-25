package service

import (
	"capstone-alta1/features/review"
	"capstone-alta1/utils/helper"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type reviewService struct {
	reviewRepository review.RepositoryInterface
	validate         *validator.Validate
}

func New(repo review.RepositoryInterface) review.ServiceInterface {
	return &reviewService{
		reviewRepository: repo,
		validate:         validator.New(),
	}
}

func (service *reviewService) Create(input review.Core, c echo.Context) (err error) {
	// validasi input
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	errCreate := service.reviewRepository.Create(input)
	if errCreate != nil {
		log.Error(errCreate.Error())
		return helper.ServiceErrorMsg(err)
	}

	return nil
}

// GetAll implements user.ServiceInterface
func (service *reviewService) GetAll(query string) (data []review.Core, err error) {

	data, err = service.reviewRepository.GetAll()

	if err != nil {
		helper.LogDebug(err)
		return nil, helper.ServiceErrorMsg(err)
	}

	if len(data) == 0 {
		helper.LogDebug("Get data success. No data.")
		return nil, errors.New("Get data success. No data.")
	}

	return data, err
}

func (service *reviewService) GetById(id int) (data review.Core, err error) {
	data, err = service.reviewRepository.GetById(id)
	if err != nil {
		log.Error(err.Error())
		return review.Core{}, helper.ServiceErrorMsg(err)
	}

	if data == (review.Core{}) {
		helper.LogDebug("Get data success. No data.")
		return review.Core{}, errors.New("Get data success. No data.")
	}

	return data, err

}

func (service *reviewService) Update(input review.Core, id int, c echo.Context) error {
	// validasi input
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	// validasi user dgn id path param, apakah ada datanya di database
	_, errFindId := service.reviewRepository.GetById(id)
	if errFindId != nil {
		log.Error(errFindId.Error())
		return helper.ServiceErrorMsg(errFindId)
	}

	// proses
	err := service.reviewRepository.Update(input, id)
	if err != nil {
		log.Error(err.Error())
		return helper.ServiceErrorMsg(err)
	}

	return nil
}

func (service *reviewService) Delete(id int) error {
	// validasi user dgn id path param, apakah ada datanya di database
	_, errFindId := service.reviewRepository.GetById(id)
	if errFindId != nil {
		log.Error(errFindId.Error())
		return helper.ServiceErrorMsg(errFindId)
	}

	// proses
	err := service.reviewRepository.Delete(id)
	if err != nil {
		log.Error(err.Error())
		return helper.ServiceErrorMsg(err)
	}
	return nil
}
