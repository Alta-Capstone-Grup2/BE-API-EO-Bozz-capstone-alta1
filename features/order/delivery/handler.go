package delivery

import (
	"capstone-alta1/features/order"
	"capstone-alta1/middlewares"
	"capstone-alta1/utils/helper"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type orderDelivery struct {
	orderService order.ServiceInterface
}

func New(order order.ServiceInterface, e *echo.Echo) {
	handler := &orderDelivery{
		orderService: order,
	}

	e.POST("/orders", handler.Create, middlewares.JWTMiddleware())
	e.GET("/orders", handler.GetAll, middlewares.JWTMiddleware())
	e.GET("/orders/:id", handler.GetById, middlewares.JWTMiddleware())
	e.PUT("/orders/:id", handler.UpdateStatusCancel, middlewares.JWTMiddleware())
	e.PUT("/orders/:id/payout", handler.UpdateStatusPayout, middlewares.JWTMiddleware())
}

func (delivery *orderDelivery) Create(c echo.Context) error {
	orderInput := OrderRequest{}
	errBind := c.Bind(&orderInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. "+errBind.Error()))
	}

	inputClientID := middlewares.ExtractTokenClientID(c)
	dataCore := toCore(orderInput, uint(inputClientID))
	dataDetailOrder := toDetailOrder(orderInput)
	err := delivery.orderService.Create(dataCore, dataDetailOrder)
	if err != nil {
		if strings.Contains(err.Error(), "Error:Field validation") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Some field cannot Empty. Details : "+err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed insert data. "+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("Success create data"))
}

func (delivery *orderDelivery) GetAll(c echo.Context) error {
	query := c.QueryParam("event_name")
	helper.LogDebug("\n isi query = ", query)

	results, err := delivery.orderService.GetAll(query)
	if err != nil {
		if strings.Contains(err.Error(), "Get data success. No data.") {
			return c.JSON(http.StatusOK, helper.SuccessWithDataResponse(err.Error(), results))
		}
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromCoreList(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read all data.", dataResponse))
}

func (delivery *orderDelivery) GetById(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}
	results, results2, err := delivery.orderService.GetById(uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "Get data success. No data.") {
			return c.JSON(http.StatusOK, helper.SuccessWithDataResponse(err.Error(), results))
		}
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromCoreDetail(results, results2)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read user.", dataResponse))
}

func (delivery *orderDelivery) UpdateStatusCancel(c echo.Context) error {
	idParam := c.Param("id")
	orderId, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}

	orderInput := OrderStatusRequest{}
	errBind := c.Bind(&orderInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. "+errBind.Error()))
	}

	dataCore := toCoreStatus(orderInput, uint(orderId))
	err := delivery.orderService.UpdateStatusCancel(dataCore, uint(orderId))
	if err != nil {
		if strings.Contains(err.Error(), "Error:Field validation") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Some field cannot Empty. Details : "+err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed update data. "+err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success update data."))
}

func (delivery *orderDelivery) UpdateStatusPayout(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}

	err := delivery.orderService.UpdateStatusPayout(uint(id), c)
	if err != nil {
		if strings.Contains(err.Error(), "Error:Field validation") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Some field cannot Empty. Details : "+err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed update data. "+err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success update data."))
}
