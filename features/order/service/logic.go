package service

import (
	cfg "capstone-alta1/config"
	_order "capstone-alta1/features/order"
	"capstone-alta1/utils/helper"
	"capstone-alta1/utils/thirdparty"
	"errors"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
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

func (order *orderService) Create(inputOrder _order.Core, inputDetail []_order.DetailOrder) (data _order.Core, err error) {
	strUuid := uuid.New()
	transactionID := "INV-" + helper.GetDateNowShort() + "-" + strUuid.String()

	serviceData, errGetServiceData := order.orderRepository.GetServiceByID(inputOrder.ServiceID)
	if errGetServiceData != nil {
		helper.LogDebug("Order - logic - GetServiceByID | Error execute GetServiceByID. Error  = ", errGetServiceData.Error())
		return _order.Core{}, helper.ServiceErrorMsg(err)
	}

	if reflect.DeepEqual(serviceData, _order.Service{}) {
		return _order.Core{}, errors.New("Service not found. Please check input again.")
	}

	helper.LogDebug("Order - logic - GetServiceByID | Service Data  = ", serviceData)
	helper.LogDebug("Order - logic - GetServiceByID | Transaction ID  = ", transactionID)
	inputOrder.MidtransTransactionID = transactionID
	inputOrder.ServiceName = serviceData.ServiceName
	inputOrder.ServicePrice = serviceData.ServicePrice

	var grossAmmout uint
	for _, val := range inputDetail {
		grossAmmout += val.DetailOrderTotal
	}
	grossAmmout += inputOrder.ServicePrice

	inputOrder.GrossAmmount = grossAmmout
	inputOrder.OrderStatus = cfg.ORDER_STATUS_WAITING_CONFIRMATION
	helper.LogDebug("Order - logic - GetServiceByID | Input Order   = ", helper.ConvToJson(inputOrder))

	midtransObj := thirdparty.OrderMidtrans(transactionID, int64(inputOrder.GrossAmmount))
	if midtransObj.ErrorMessages != nil {
		helper.LogDebug("Order - logic - Failed process to midtrans")
		return _order.Core{}, errors.New("Payment Failed. Please try again later.")
	}
	inputOrder.MidtransLink = midtransObj.RedirectURL
	inputOrder.MidtransToken = midtransObj.Token
	if inputOrder.MidtransLink == "" {
		helper.LogDebug("Order - logic - Failed get Midtrans link.")
		return _order.Core{}, errors.New("Payment Failed. Please try again later.")
	}
	helper.LogDebug("Order - logic - GetServiceByID | Input Order   = ", inputOrder)

	data, errCreate := order.orderRepository.Create(inputOrder, inputDetail)
	if errCreate != nil {
		helper.LogDebug("Order - logic - Create | Error execute create order. Error  = ", errCreate.Error())
		return _order.Core{}, helper.ServiceErrorMsg(errCreate)
	}

	return data, nil
}

func (order *orderService) GetAll(query string) (data []_order.OrderJoinPartner, err error) {
	// if query == "" {
	data, err = order.orderRepository.GetAll()
	if err != nil {
		helper.LogDebug(err)
		return nil, helper.ServiceErrorMsg(err)
	}
	// } else if query == "query" {
	// 	data, err = order.orderRepository.GetAllWithSearch(query)
	// 	if err != nil {
	// 		helper.LogDebug(err)
	// 		return nil, helper.ServiceErrorMsg(err)
	// 	}
	// }
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
	dataOrder, _, errFindOrder := order.orderRepository.GetById(id)
	if errFindOrder != nil {
		helper.LogDebug("Order - UpdateStatusPayout. Find Order Failed.Erro = ", errFindOrder)
		return helper.ServiceErrorMsg(errFindOrder)
	}

	// convert status to pascal case
	dataOrder.OrderStatus = strings.Title(strings.ToLower(dataOrder.OrderStatus))

	helper.LogDebug("Order - UpdateStatusPayout. DataOrder = ", dataOrder)

	if dataOrder.OrderStatus == cfg.ORDER_STATUS_PAID_OFF {
		return errors.New("Order Status already Paid Off. Please check again.")
	}

	if dataOrder.OrderStatus != cfg.ORDER_STATUS_COMPLETE_ORDER {
		return errors.New("Order Status not Complete yet. Please check again.")
	}

	dataOrder.OrderStatus = cfg.ORDER_STATUS_PAID_OFF

	var errUpload error
	dataOrder.PayoutRecieptFile, errUpload = thirdparty.Upload(c, cfg.ORDER_PAYOUT_RECEIPT_FILE, cfg.ORDER_FOLDER)
	if errUpload != nil {
		return errUpload
	}

	errUpdateStatus := order.orderRepository.UpdateStatusPayout(dataOrder, id)
	if errUpdateStatus != nil {
		helper.LogDebug("Order - UpdateStatusPayout. Update Failed. Error = ", errUpdateStatus.Error())
		return helper.ServiceErrorMsg(errUpdateStatus)
	}

	return nil
}

// SERVICE TO UPDATE BOOKING DATA AFTER PAYMENT MIDTRANS
func (order *orderService) UpdateMidtrans(input _order.Core) error {
	inputMidtrans := thirdparty.CheckMidtrans(strconv.Itoa(int(input.ID)))

	helper.LogDebug("Midtrans data, ", inputMidtrans)

	input.OrderStatus = inputMidtrans.TransactionStatus
	err := order.orderRepository.UpdateMidtrans(input)

	if err != nil {
		helper.LogDebug("Order - UpdateMidtrans. Update Failed. Error = ", err.Error())
		return err
	}
	return err
}
