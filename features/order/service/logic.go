package service

import (
	cfg "capstone-alta1/config"
	_order "capstone-alta1/features/order"
	"capstone-alta1/utils/helper"
	"capstone-alta1/utils/thirdparty"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type orderService struct {
	orderRepository _order.RepositoryInterface
	validate        *validator.Validate
}

func New(repo _order.RepositoryInterface) _order.ServiceInterface {
	return &orderService{
		orderRepository: repo,
		validate:        validator.New(),
	}
}

func (order *orderService) Create(inputOrder _order.Core, inputDetail []_order.DetailOrder) (err error) {
	errCreate := order.orderRepository.Create(inputOrder, inputDetail)
	if errCreate != nil {
		log.Error(errCreate.Error())
		return helper.ServiceErrorMsg(err)
	}

	return nil
}

func (order *orderService) GetAll(query string) (data []_order.Core, err error) {
	if query == "" {
		data, err = order.orderRepository.GetAll()
		if err != nil {
			helper.LogDebug(err)
			return nil, helper.ServiceErrorMsg(err)
		}
	} else if query == "query" {
		data, err = order.orderRepository.GetAllWithSearch(query)
		if err != nil {
			helper.LogDebug(err)
			return nil, helper.ServiceErrorMsg(err)
		}
	}
	return data, err
}

func (order *orderService) GetById(id uint) (data _order.Core, dataDetail []_order.DetailOrder, err error) {
	data, dataDetail, err = order.orderRepository.GetById(id)
	if err != nil {
		log.Error(err.Error())
		return _order.Core{}, nil, err
	}
	return data, dataDetail, err
}

func (order *orderService) UpdateStatusCancel(input _order.Core, id uint) error {
	err := order.orderRepository.UpdateStatusCancel(input, id)
	if err != nil {
		log.Error(err.Error())
		return helper.ServiceErrorMsg(err)
	}

	return nil
}

func (order *orderService) UpdateStatusPayout(id uint, c echo.Context) error {
	input := _order.Core{}
	var errUpload error
	input.PayoutRecieptFile, errUpload = thirdparty.Upload(c, cfg.SERVICE_IMAGE_FILE, cfg.SERVICE_FOLDER)
	if errUpload != nil {
		return errUpload
	}

	err := order.orderRepository.UpdateStatusPayout(input, id)
	if err != nil {
		log.Error(err.Error())
		return helper.ServiceErrorMsg(err)
	}

	return nil
}
